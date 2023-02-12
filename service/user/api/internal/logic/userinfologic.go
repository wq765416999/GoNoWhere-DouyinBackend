package logic

import (
	"context"
	"encoding/json"
	"mini-douyin/service/user/rpc/userclient"

	"mini-douyin/service/user/api/internal/svc"
	"mini-douyin/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.Douyin_user_response, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.DouyinUserRequest{
		UserId: int64(uid),
	})
	if err != nil {
		return nil, err
	}

	return &types.Douyin_user_response{
		User: types.Douyin_user_info{ID: int(res.User.Id), Name: res.User.Name},
	}, nil
}
