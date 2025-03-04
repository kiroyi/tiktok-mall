package dal

import (
	"github.com/cloudewego/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
