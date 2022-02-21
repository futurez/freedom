package flog

import "go.uber.org/zap"

type Level int8

const (
	DebugLevel = 0 + iota
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)

var (
	logger *zap.Logger
	supar  *zap.SugaredLogger
)

func Debug(args ...interface{}) {
	supar.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	supar.Debugf(template, args...)
}

func Info(args ...interface{}) {
	supar.Info(args...)
}

func Infof(template string, args ...interface{}) {
	supar.Infof(template, args...)
}

func Warn(args ...interface{}) {
	supar.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	supar.Warnf(template, args...)
}

func Error(args ...interface{}) {
	supar.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	supar.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	supar.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	supar.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	supar.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	supar.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	supar.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	supar.Fatalf(template, args...)
}
