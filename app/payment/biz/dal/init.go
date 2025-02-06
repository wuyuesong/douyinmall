package dal

import (
	"github.com/wuyuesong/gomall/app/payment/biz/dal/mysql"
	"github.com/wuyuesong/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
