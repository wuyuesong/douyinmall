package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/wuyuesong/gomall/hertz_gen/frontend/common"
	product "github.com/wuyuesong/gomall/hertz_gen/frontend/product"
	"github.com/wuyuesong/gomall/infra/rpc"
	rpcproduct "github.com/wuyuesong/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductReq) (resp *common.Empty, err error) {
	products, err := rpc.ProductClient.SearchProduct(h.Context, &rpcproduct.SearchProductReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items": products.Results,
		"q":     req.Q,
	}, nil
}
