package dal

import (
	"github.com/kiroyi/gomall/biz/dal/mysql"
	"github.com/kiroyi/gomall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
