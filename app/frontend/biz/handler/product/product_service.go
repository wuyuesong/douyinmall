package product

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	commonUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wuyuesong/douyinmall/app/frontend/biz/service"
	"github.com/wuyuesong/douyinmall/app/frontend/biz/utils"
	product "github.com/wuyuesong/douyinmall/app/frontend/hertz_gen/frontend/product"
)

const (
	UploadDir      = "/home/pictures" // 文件保存目录
	MaxUploadSize  = 2 << 20          // 2MB
	AllowedFormats = "image/jpeg,image/png,image/gif"
)

// GetProduct .
// @router /product [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProdcutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "product", utils.WarpResponse(ctx, c, resp))
}

// SearchProducts .
// @router /search [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewSearchProductsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "search", resp)
}

func UploadImages(ctx context.Context, c *app.RequestContext) {
	// 1. 验证请求大小
	if c.Request.Header.ContentLength() > MaxUploadSize {
		c.JSON(consts.StatusRequestEntityTooLarge, commonUtils.H{
			"error": "文件大小不能超过2MB",
		})
		return
	}

	// 2. 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(consts.StatusBadRequest, commonUtils.H{
			"error": "获取文件失败: " + err.Error(),
		})
		return
	}

	// 3. 验证文件类型
	if !isAllowedType(file.Header.Get("Content-Type")) {
		c.JSON(consts.StatusUnsupportedMediaType, commonUtils.H{
			"error": "只允许上传图片文件（JPEG/PNG/GIF）",
		})
		return
	}

	// 4. 创建保存目录
	if err := os.MkdirAll(UploadDir, 0o755); err != nil {
		hlog.Errorf("创建上传目录失败: %v", err)
		c.JSON(consts.StatusInternalServerError, commonUtils.H{
			"error": "服务器内部错误",
		})
		return
	}

	// 5. 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(UploadDir, newFilename)

	// 6. 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		hlog.Errorf("文件保存失败: %v", err)
		c.JSON(consts.StatusInternalServerError, commonUtils.H{
			"error": "文件保存失败",
		})
		return
	}

	// 7. 返回成功响应
	c.JSON(consts.StatusOK, commonUtils.H{
		"url": fmt.Sprintf("%s%s", UploadDir, newFilename),
	})
}

// 验证文件类型是否允许
func isAllowedType(contentType string) bool {
	switch contentType {
	case "image/jpeg", "image/png", "image/gif":
		return true
	}
	return false
}

// AddProduct .
// @router /product [POST]
func AddProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.AddProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewAddProductService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ProdcutPageNum .
// @router /product/pageNum [GET]
func ProdcutPageNum(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProdcutPageNumReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewProdcutPageNumService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
