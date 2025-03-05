package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/wuyuesong/douyinmall/app/product/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/product/biz/model"
	product "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	// Finish your business logic.
	var categorieId uint
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	if req.Categorie == "normal" {
		categorieId = 1
	} else if req.Categorie == "discount" {
		categorieId = 2
	}
	productItem := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	productId, err := productQuery.AddProduct(productItem, categorieId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &product.AddProductResp{ProductId: uint32(productId)}, nil
}
