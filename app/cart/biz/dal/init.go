package dal

import (
	"github.com/wuyuesong/gomall/app/cart/biz/dal/mysql"
	"github.com/wuyuesong/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
