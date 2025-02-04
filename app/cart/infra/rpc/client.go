package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/wuyuesong/gomall/app/cart/conf"
	cartutils "github.com/wuyuesong/gomall/app/cart/utils"
	"github.com/wuyuesong/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
