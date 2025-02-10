package dal

import (
	"github.com/wuyuesong/douyinmall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
