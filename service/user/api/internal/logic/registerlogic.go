package logic

import (
	"context"
	"mini-douyin/service/user/rpc/userclient"

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
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	print(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &types.Douyin_user_register_response{
		UserID: int(res.UserId),
	}, nil
}
