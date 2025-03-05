package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	utils "github.com/cloudwego/hertz/pkg/common/utils"
	product "github.com/wuyuesong/douyinmall/app/frontend/hertz_gen/frontend/product"
	"github.com/wuyuesong/douyinmall/app/frontend/infra/rpc"
	rpcproduct "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

type AddProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddProductService(Context context.Context, RequestContext *app.RequestContext) *AddProductService {
	return &AddProductService{RequestContext: RequestContext, Context: Context}
}

func (h *AddProductService) Run(req *product.AddProductReq) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	productId, err := rpc.ProductClient.AddProduct(h.Context, &rpcproduct.AddProductReq{
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		Description: req.Description,
		Picture:     req.Picture,
		Categorie:   req.Categorie,
	})

	return utils.H{
		"productId": productId,
	}, err
}
