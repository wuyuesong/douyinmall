package service

import (
	"context"

	"github.com/wuyuesong/douyinmall/app/product/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/product/biz/model"
	product "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

type ListProductsAllService struct {
	ctx context.Context
} // NewListProductsAllService new ListProductsAllService
func NewListProductsAllService(ctx context.Context) *ListProductsAllService {
	return &ListProductsAllService{ctx: ctx}
}

// Run create note info
func (s *ListProductsAllService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewProductQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.ListProductsAll(req.Page, req.PageSize, req.CategoryName)
	resp = &product.ListProductsResp{}
	for _, v := range c.Products {
		resp.Products = append(resp.Products, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
			Stock:       v.Stock,
		})
	}
	resp.TotalPage = uint32(c.TotalPages)
	return
}
