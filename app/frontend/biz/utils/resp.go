package utils

import (
	"context"
	"errors"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/wuyuesong/douyinmall/app/frontend/infra/rpc"
	frontendUtils "github.com/wuyuesong/douyinmall/app/frontend/utils"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/cart"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	var bizErr *kerrors.BizStatusError
	if errors.As(err, &bizErr) {
		c.JSON(http.StatusOK, ApiResponse{
			Code:    int(bizErr.BizStatusCode()),
			Message: bizErr.BizMessage(),
			Data:    nil,
		})
	} else {
		c.JSON(http.StatusInternalServerError, ApiResponse{
			Code:    50000,
			Message: "系统错误",
			Data:    nil,
		})
	}
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, ApiResponse{
		Code:    200,
		Message: "成功",
		Data:    data,
	})
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := frontendUtils.GetUserIdFromCtx(ctx)
	content["user_id"] = userId

	if userId > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: uint32(userId),
		})
		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Items)
		}
	}

	return content
}
