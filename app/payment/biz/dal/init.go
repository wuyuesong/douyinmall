package dal

import (
	"github.com/wuyuesong/douyinmall/app/payment/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
