package dal

import (
	"github.com/wuyuesong/gomall/app/order/biz/dal/mysql"
	"github.com/wuyuesong/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
