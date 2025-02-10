package dal

import (
	"github.com/wuyuesong/douyinmall/app/user/biz/dal/mysql"
	"github.com/wuyuesong/douyinmall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
