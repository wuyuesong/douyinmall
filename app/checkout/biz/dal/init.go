package dal

import (
	"github.com/wuyuesong/gomall/app/checkout/biz/dal/mysql"
	"github.com/wuyuesong/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
