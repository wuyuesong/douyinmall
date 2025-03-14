package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	utils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	common "github.com/wuyuesong/douyinmall/app/gateway/hertz_gen/gateway/common"
	"github.com/wuyuesong/douyinmall/app/gateway/infra/rpc"
	rpcproduct "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *common.Empty) (resp map[string]any, err error) {
	// 正确获取路由参数
	idStr := h.RequestContext.Param("id")
	if idStr == "" {
		return utils.H{
			"code": consts.StatusBadRequest,
			"msg":  "缺少商品ID参数",
		}, nil
	}

	// 字符串转uint32
	parsedID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.H{
			"code": consts.StatusBadRequest,
			"msg":  "商品ID必须为数字",
		}, nil
	}
	id := uint32(parsedID)

	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: id})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"item": p.Product,
	}, nil
}
