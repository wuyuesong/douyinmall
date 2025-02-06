package dal

import (
	"github.com/wuyuesong/gomall/app/email/biz/dal/mysql"
	"github.com/wuyuesong/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
