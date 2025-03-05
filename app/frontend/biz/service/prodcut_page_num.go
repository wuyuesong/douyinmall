package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/wuyuesong/douyinmall/app/frontend/hertz_gen/frontend/common"
	product "github.com/wuyuesong/douyinmall/app/frontend/hertz_gen/frontend/product"
)

type ProdcutPageNumService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewProdcutPageNumService(Context context.Context, RequestContext *app.RequestContext) *ProdcutPageNumService {
	return &ProdcutPageNumService{RequestContext: RequestContext, Context: Context}
}

func (h *ProdcutPageNumService) Run(req *product.ProdcutPageNumReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
