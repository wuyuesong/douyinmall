package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	product "github.com/wuyuesong/douyinmall/app/gateway/hertz_gen/gateway/product"
	"github.com/wuyuesong/douyinmall/app/gateway/infra/rpc"
	rpcproduct "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductReq{
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
