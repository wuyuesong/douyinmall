package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	checkout "github.com/wuyuesong/gomall/app/frontend/hertz_gen/frontend/checkout"
	common "github.com/wuyuesong/gomall/app/frontend/hertz_gen/frontend/common"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
