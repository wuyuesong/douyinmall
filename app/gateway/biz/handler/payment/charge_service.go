package payment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wuyuesong/douyinmall/app/gateway/biz/service"
	"github.com/wuyuesong/douyinmall/app/gateway/biz/utils"
	payment "github.com/wuyuesong/douyinmall/app/gateway/hertz_gen/gateway/payment"
)

// PaymentCreate .
// @router /payment [POST]
func PaymentCreate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.ChargeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewPaymentCreateService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
