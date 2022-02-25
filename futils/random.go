package futils

import (
	"bytes"
	"crypto/rand" //加密安全的随机数生成器
	"math/big"
	mRand "math/rand" //伪随机数生成器
	"time"
)

func init() {
	//如果每次运行需要不同的行为，使用种子函数初始化默认的源。
	mRand.Seed(time.Now().UTC().UnixNano())
}

func RandInt32(min, max int32) int32 {
	if min >= max || max == 0 {
		return max
	}
	return mRand.Int31n(max-min) + min
}

func RandInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	return mRand.Intn(max-min) + min
}

func RandInt64(min, max int64) int64 {
	if min >= max || max == 0 {
		return max
	}
	return mRand.Int63n(max-min) + min
}

func RandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	buffer := bytes.NewBufferString(str)
	length := buffer.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
