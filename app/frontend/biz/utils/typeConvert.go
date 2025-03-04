package utils

import (
	"fmt"
	"strconv"
)

// 封装工具函数（放到 utils 包中）

func ParseInt32(s string) (int32, error) {
	val, err := strconv.ParseInt(s, 10, 32) // 关键点：第三个参数设为 32 会自动检查 int32 范围
	if err != nil {
		return 0, fmt.Errorf("invalid int32: %s", s)
	}
	return int32(val), nil
}
