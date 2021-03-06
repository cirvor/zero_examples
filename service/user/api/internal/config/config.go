package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct { // jwt鉴权配置
		AccessSecret string // jwt密钥
		AccessExpire int64 // 有效期，单位：秒
	}
	MysqlRead struct { // 数据库配置，除mysql外，可能还有mongo等其他数据库
		DataSource string // mysql链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
	}
	CacheRedis cache.CacheConf // redis缓存
	LogConf logx.LogConf
}
