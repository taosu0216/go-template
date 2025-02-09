package utils

import (
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// GenerateRandomString 生成指定长度的随机字符串，包含大小写字母和数字
func GenerateRandomString() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	b := make([]byte, 6)
	for i := 0; i < 6; i++ {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}
