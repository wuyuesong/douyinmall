package service

import (
	"context"
	"testing"
	product "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product"
)

func TestListProductsAll_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductsAllService(ctx)
	// init req and assert value

	req := &product.ListProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
