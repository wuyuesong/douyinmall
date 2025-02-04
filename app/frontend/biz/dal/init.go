package dal

import (
	"github.com/wuyuesong/gomall/app/frontend/biz/dal/mysql"
	"github.com/wuyuesong/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
