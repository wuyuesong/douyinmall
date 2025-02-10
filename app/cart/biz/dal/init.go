package dal

import (
	"github.com/wuyuesong/douyinmall/app/cart/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
