package redis

import (
	"github.com/HCH1212/blog/backend/conf"
	uredis "github.com/HCH1212/utils/redis"
	"github.com/go-redis/redis/v8"
)

var ReClient *redis.Client

func Init() {
	uredis.InitRedis(conf.GetConf().Redis.Address, conf.GetConf().Redis.Password, conf.GetConf().Redis.DB)
	ReClient = uredis.GetRedisClient()
}
