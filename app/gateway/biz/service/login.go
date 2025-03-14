package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/wuyuesong/douyinmall/app/gateway/hertz_gen/gateway/auth"
	"github.com/wuyuesong/douyinmall/app/gateway/infra/rpc"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *user.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	loginResp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", loginResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return loginResp, nil
}
