package dal

import (
	"github.com/wuyuesong/douyinmall/app/email/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
