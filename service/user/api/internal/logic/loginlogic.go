package logic

import (
	"context"

	"mini-douyin/service/user/api/internal/svc"
	"mini-douyin/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Douyin_user_login_request) (resp *types.Douyin_user_login_response, err error) {
	// todo: add your logic here and delete this line

	return
}
