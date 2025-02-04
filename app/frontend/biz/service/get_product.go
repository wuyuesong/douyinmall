package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	product "github.com/wuyuesong/gomall/hertz_gen/frontend/product"
	"github.com/wuyuesong/gomall/infra/rpc"
	rpcproduct "github.com/wuyuesong/gomall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProdcutReq) (resp map[string]any, err error) {
	rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}
