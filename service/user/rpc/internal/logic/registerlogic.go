package logic

import (
	"context"

	"mini-douyin/service/user/rpc/internal/svc"
	"mini-douyin/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	// todo: add your logic here and delete this line

	return &user.DouyinUserRegisterResponse{}, nil
}
