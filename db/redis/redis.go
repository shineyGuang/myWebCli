package redis

import (
	"fmt"
	"myWebCli/settings"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init(cfg *settings.Redis) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("redis connect failed", zap.Error(err))
		return err
	}
	return
}

func Close() {
	_ = rdb.Close()
}
