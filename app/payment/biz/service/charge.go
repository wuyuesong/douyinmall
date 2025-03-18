package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/wuyuesong/douyinmall/app/payment/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/payment/biz/model"
	"github.com/wuyuesong/douyinmall/app/payment/infra/mq"
	"github.com/wuyuesong/douyinmall/app/payment/infra/rpc"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/email"
	payment "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/payment"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/user"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}

	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}

	userResp, err := rpc.UserClient.GetUser(s.ctx, &user.GetUserReq{UserId: int32(req.UserId)})
	fmt.Println("userResp: ", userResp)
	fmt.Println("err: ", err)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005003, err.Error())
	}

	emailMsg, _ := json.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          userResp.Email,
		ContentType: "text/plain", // 或改为 "text/plain; charset=UTF-8"
		Subject:     "订单支付成功",
		Content:     fmt.Sprintf("您成功购买了商品，订单号为%s", req.OrderId),
		OrderId:     req.OrderId,
	})

	// 创建 RocketMQ 消息（Topic: EmailTopic, Tag: payment_success）
	rmqMsg := primitive.NewMessage("EmailTopic", emailMsg)
	rmqMsg.WithTag("payment_success")
	carrier := propagation.HeaderCarrier{}
	otel.GetTextMapPropagator().Inject(s.ctx, carrier)

	for key, values := range carrier {
		for _, value := range values {
			rmqMsg.WithProperty(key, value)
		}
	}
	// 异步发送消息
	fmt.Print("start to send email")
	fmt.Print("mq.RmqProducer: ", mq.RmqProducer)
	err = mq.RmqProducer.SendAsync(s.ctx, func(ctx context.Context, result *primitive.SendResult, e error) {
		if e != nil {
			klog.Error("邮件发送失败: %v, 订单ID: %s\n", e, req.OrderId)
		} else {
			klog.Info("邮件发送成功: MsgID=%s, 订单ID=%s\n", result.MsgID, req.OrderId)
		}
	}, rmqMsg)
	if err != nil {
		fmt.Print(err)
		return nil, kerrors.NewGRPCBizStatusError(4005005, "邮件服务暂不可用")
	}

	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
