package dal

import (
	"github.com/wuyuesong/douyinmall/app/product/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
