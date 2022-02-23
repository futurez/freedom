package futils

import (
	"testing"
)

type RedisConfig struct {
	Url string
	Pwd string
}

type configs struct {
	IsTest   bool
	SvrIP    string
	SvrPort  int
	UrlDebug RedisConfig
}

func TestHandle(t *testing.T) {
	cfg := configs{}
	ReadConfig("test.cfg", &cfg)
	t.Logf("%+v", cfg)
}
