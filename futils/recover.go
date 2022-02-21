package futils

import (
	"github.com/futurez/freedom/flog"
)

//通用recover函数，在单独协程的最开始使用defer调用
func RecoverFromPanic(fName string, cb func()) {
	if r := recover(); r != nil {
		flog.Errorf("%s recover from panic!!!, error:%v", fName, r)
		if cb != nil {
			cb()
		}
	}
}
