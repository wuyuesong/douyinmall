package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	utils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/wuyuesong/douyinmall/app/gateway/biz/constants"
	checkout "github.com/wuyuesong/douyinmall/app/gateway/hertz_gen/gateway/checkout"
	"github.com/wuyuesong/douyinmall/app/gateway/infra/rpc"
	rpccheckout "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/checkout"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/user"
)

type CheckoutCreateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutCreateService(Context context.Context, RequestContext *app.RequestContext) *CheckoutCreateService {
	return &CheckoutCreateService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutCreateService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
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

	rpcResp, err := rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId: uint32(userId),
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			State:         req.Province,
			StreetAddress: req.Street,
		},
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"orderId": rpcResp.OrderId,
	}, nil
}
