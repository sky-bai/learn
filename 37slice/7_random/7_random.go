package __random

import (
	"math/rand"
	"time"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func RandomStr(len int) string {
	if len == 0 {
		len = 32
	}

	rand.Seed(time.Now().UnixNano())

	pwd := make([]byte, len)
	for i := 0; i < len; i++ {
		pwd[i] = chars[rand.Intn(len)]
	}

	return string(pwd)
}
