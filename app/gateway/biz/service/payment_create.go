package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	utils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/wuyuesong/douyinmall/app/gateway/biz/constants"
	paymentHertz "github.com/wuyuesong/douyinmall/app/gateway/hertz_gen/gateway/payment"
	"github.com/wuyuesong/douyinmall/app/gateway/infra/rpc"
	paymentKitex "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/payment"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/user"
)

type PaymentCreateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPaymentCreateService(Context context.Context, RequestContext *app.RequestContext) *PaymentCreateService {
	return &PaymentCreateService{RequestContext: RequestContext, Context: Context}
}

func (h *PaymentCreateService) Run(req *paymentHertz.ChargeReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var userId int32
	if userInfo, exists := h.RequestContext.Get(constants.IdentityKey); exists {
		userId = userInfo.(user.LoginResp).UserId
	} else {
		h.RequestContext.JSON(401, utils.H{"error": "未认证用户"})
		return
	}

	payReq := &paymentKitex.ChargeReq{
		UserId:  uint32(userId),
		OrderId: req.OrderId,
		Amount:  req.Amount,
		CreditCard: &paymentKitex.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	paymentResult, err := rpc.PaymentClient.Charge(h.Context, payReq)
	if err != nil {
		return nil, err
	}
	return utils.H{
		"TransactionId": paymentResult.TransactionId,
	}, nil
}
