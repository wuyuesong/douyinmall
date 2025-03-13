package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	common "github.com/wuyuesong/douyinmall/app/frontend/hertz_gen/frontend/common"
)

type ImageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

const (
	ImageBasePath     = "/home/pictures/"      // 集中管理图片存储路径
	AllowedExtensions = ".jpg,.jpeg,.png,.gif" // 允许的文件类型白名单
)

func NewImageService(Context context.Context, RequestContext *app.RequestContext) *ImageService {
	return &ImageService{RequestContext: RequestContext, Context: Context}
}

func (h *ImageService) Run(req *common.Empty) (common.Empty, error) {
	// 获取路由参数
	imagePath := h.RequestContext.Param("image")
	if imagePath == "" {
		return common.Empty{}, errors.New("image参数不能为空")
	}

	fullPath := ImageBasePath + imagePath

	// 检查文件是否存在
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		h.RequestContext.AbortWithStatus(consts.StatusNotFound)
		return common.Empty{}, errors.New("图片不存在")
	}

	// 读取文件并返回
	fileBytes, err := os.ReadFile(fullPath)
	if err != nil {
		return common.Empty{}, errors.New("文件读取失败")
	}

	h.RequestContext.Data(
		consts.StatusOK,
		getMimeType(fullPath),
		fileBytes,
	)
	fmt.Print(h.RequestContext.Data)
	return common.Empty{}, nil
}

// 安全校验函数
func validateImagePath(filename string) error {
	// 防止路径遍历
	if strings.Contains(filename, "..") || strings.HasPrefix(filename, "/") {
		return errors.New("非法路径")
	}

	// 检查文件扩展名
	ext := strings.ToLower(filepath.Ext(filename))
	allowed := strings.Split(AllowedExtensions, ",")
	for _, a := range allowed {
		if ext == a {
			return nil
		}
	}
	return errors.New("不支持的文件类型")
}

// 获取MIME类型
func getMimeType(filename string) string {
	switch filepath.Ext(filename) {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}
