package dal

import (
	"github.com/wuyuesong/douyinmall/app/checkout/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
