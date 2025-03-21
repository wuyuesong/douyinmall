package service

import (
	"context"

	"github.com/wuyuesong/douyinmall/app/product/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/product/biz/model"
	product "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	var results []*product.Product
	for _, v := range products {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}

	return &product.SearchProductResp{
		Results: results,
	}, err
}
