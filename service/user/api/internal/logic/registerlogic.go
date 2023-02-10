package logic

import (
	"context"

	"mini-douyin/service/user/api/internal/svc"
	"mini-douyin/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Douyin_user_register_request) (resp *types.Douyin_user_register_response, err error) {
	// todo: add your logic here and delete this line

	return
}
