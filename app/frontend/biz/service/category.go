package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wuyuesong/gomall/biz/utils"
	category "github.com/wuyuesong/gomall/hertz_gen/frontend/category"
	common "github.com/wuyuesong/gomall/hertz_gen/frontend/common"
	"github.com/wuyuesong/gomall/infra/rpc"
	"github.com/wuyuesong/gomall/rpc_gen/kitex_gen/product"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp *common.Empty, err error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
