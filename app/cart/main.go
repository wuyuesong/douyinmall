package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/wuyuesong/gomall/app/cart/biz/dal"
	"github.com/wuyuesong/gomall/app/cart/conf"
	"github.com/wuyuesong/gomall/app/cart/infra/rpc"
	"github.com/wuyuesong/gomall/common/mtl"
	"github.com/wuyuesong/gomall/common/serversuite"
	"github.com/wuyuesong/gomall/rpc_gen/kitex_gen/cart/cartservice"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	_ = godotenv.Load()
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
	mtl.InitTracing(serviceName)
	mtl.InitLog()
	rpc.Init()
	dal.Init()
	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	opts = append(opts,
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName}),
	)
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	return
}
