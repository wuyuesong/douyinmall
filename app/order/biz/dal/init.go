package dal

import (
	"github.com/wuyuesong/douyinmall/app/order/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
