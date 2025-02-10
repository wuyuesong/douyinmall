package dal

import (
	"github.com/wuyuesong/douyinmall/app/frontend/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
