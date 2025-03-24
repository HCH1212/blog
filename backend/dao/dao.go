package dao

import (
	"github.com/HCH1212/blog/backend/dao/redis"
	"github.com/HCH1212/blog/backend/dao/sqlite"
)

func Init() {
	sqlite.Init()
	redis.Init()
}
