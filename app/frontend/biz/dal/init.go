package dal

import (
	"github.com/wuyuesong/gomall/biz/dal/mysql"
	"github.com/wuyuesong/gomall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
