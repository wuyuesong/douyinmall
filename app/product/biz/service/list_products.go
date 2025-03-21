package service

import (
	"context"

	"github.com/wuyuesong/douyinmall/app/product/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/product/biz/model"
	product "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.GetProductsByCategoryName(req.Page, req.PageSize, req.CategoryName)
	resp = &product.ListProductsResp{}
	for _, v1 := range c.Categories {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
				Stock:       v.Stock,
			})
		}
	}
	resp.TotalPage = uint32(c.TotalPages)
	return
}
