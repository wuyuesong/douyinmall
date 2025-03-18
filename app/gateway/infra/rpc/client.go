package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/transmeta"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/wuyuesong/douyinmall/app/gateway/conf"
	frontendUtils "github.com/wuyuesong/douyinmall/app/gateway/utils"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	PaymentClient  paymentservice.Client
	once           sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
		initPaymentClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	var opts []client.Option
	opts = append(opts, client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	frontendUtils.MustHandleError(err)
}
