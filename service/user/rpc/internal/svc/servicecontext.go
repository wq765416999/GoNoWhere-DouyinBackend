package svc

import (
	"mini-douyin/service/user/model"

	"mini-douyin/service/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel  model.UsersModel
	LoginModel model.LoginsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUsersModel(conn, c.CacheRedis),
		LoginModel: model.NewLoginsModel(conn, c.CacheRedis),
	}
}
