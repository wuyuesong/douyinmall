package middleware

import (
    "context"
    "github.com/cloudwego/hertz/pkg/app"

    casbinHertz "github.com/hertz-contrib/casbin" // 重命名导入避免混淆
    gormadapter "github.com/casbin/gorm-adapter/v3"
    "github.com/wuyuesong/douyinmall/app/gateway/biz/dal/mysql"
    "github.com/cloudwego/kitex/pkg/klog"
    "github.com/casbin/casbin/v2"
)
const (
	ADMIN_ROLE    = "admin"
	USER_ROLE     = "user"
    ANONY_ROLE    = "anonymous"
)
var (
	CasbinHertzMiddleware *casbinHertz.Middleware
)

func getSubjectFromJWT(ctx context.Context, c *app.RequestContext) string {
    return USER_ROLE
}




func InitCasbin() {
    enforcer , err := InitCasbinEnforcer()
    CasbinHertzMiddleware, err = casbinHertz.NewCasbinMiddlewareFromEnforcer(enforcer, getSubjectFromJWT)
    if err != nil {
        panic(err)
    }

	_,err  = enforcer.AddPolicy(ADMIN_ROLE, "Admin","*") 
    _,err  = enforcer.AddPolicy(ANONY_ROLE, "HomePage","GET")
    _,err  = enforcer.AddPolicy(USER_ROLE, "Cart","*")
    _,err  = enforcer.AddPolicy(USER_ROLE, "Order","*")
    _,err  = enforcer.AddPolicy(USER_ROLE, "Payment","*")

    enforcer.SavePolicy()
}




func InitCasbinEnforcer() (*casbin.Enforcer,error) { 
    a, err := gormadapter.NewAdapterByDB(mysql.DB)
	if err != nil {
		klog.Fatal(err)
        return nil,err
	}
    Enforcer, err := casbin.NewEnforcer("./middleware/conf/model.conf", a)
	if err != nil {
		klog.Fatal(err)
        return nil,err
	}
	Enforcer.EnableAutoSave(true)

	err = Enforcer.LoadPolicy()
	if err != nil {
		klog.Fatal(err)
        return nil,err
	}
    return Enforcer,nil
}