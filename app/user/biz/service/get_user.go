package service

import (
	"context"

	"github.com/wuyuesong/douyinmall/app/user/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/user/biz/model"
	user "github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	// Finish your business logic.
	row, err := model.GetByUserID(mysql.DB, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	resp = &user.GetUserResp{
		UserId: int32(row.ID),
		Role:   row.Role,
		Email:  row.Email,
	}
	return
}
