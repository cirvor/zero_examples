package svc

import (
	"book/service/user/api/internal/config"
	"book/service/user/model"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
	UserModelNoCache model.UserModelNoCache
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewMysql(c.MysqlRead.DataSource)

	for _, redisItem := range c.CacheRedis {
		r := redisItem.RedisConf.NewRedis()
		r.Pipelined(func (pipe redis.Pipeliner) error {
			pipe.Do("select" ,1)
			return nil
		})
	}



	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		UserModelNoCache: model.NewUserModelNoCache(conn),
	}
}
