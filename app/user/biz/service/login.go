package service

import (
	"context"
	"errors"

	"github.com/wuyuesong/douyinmall/app/user/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/user/biz/model"
	"github.com/wuyuesong/douyinmall/app/user/infra/rpc"
	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/cart"
	user "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	row, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: uint32(row.ID),
	})
	if err != nil {
		return nil, err
	}

	resp = &user.LoginResp{
		UserId:  int32(row.ID),
		Role:    row.Role,
		CartNum: int32(len(cartResp.Items)),
	}

	return resp, nil
}
