package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	utils "github.com/cloudwego/hertz/pkg/common/utils"
	myUtils "github.com/wuyuesong/douyinmall/app/frontend/biz/utils"
	common "github.com/wuyuesong/douyinmall/app/frontend/hertz_gen/frontend/common"
	"github.com/wuyuesong/douyinmall/app/frontend/infra/rpc"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

type AdminHomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAdminHomeService(Context context.Context, RequestContext *app.RequestContext) *AdminHomeService {
	return &AdminHomeService{RequestContext: RequestContext, Context: Context}
}

func (h *AdminHomeService) Run(req *common.Empty) (map[string]any, error) {
	pageStr := h.RequestContext.DefaultQuery("page", "1")  // 默认第1页
	sizeStr := h.RequestContext.DefaultQuery("size", "10") // 默认每页10条
	page, err := myUtils.ParseInt32(pageStr)
	if err != nil {
		return nil, err
	}
	size, err := myUtils.ParseInt32(sizeStr)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.ProductClient.ListProductsAll(h.Context, &product.ListProductsReq{Page: page, PageSize: size})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"items":      resp.Products,
		"total_page": resp.TotalPage,
	}, nil
}
