package dal

import (
	"github.com/wuyuesong/douyinmall/app/gateway/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/gateway/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
