package flog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// 默认初始化log配置
func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999999999"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	atomLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	config := zap.Config{
		Level:             atomLevel,
		Development:       true,
		DisableStacktrace: false, //false warn, error, panic, fatal 时打印栈

		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		//InitialFields:     map[string]interface{}{"flog": 1},
	}
	skip := zap.AddCallerSkip(1)
	stack := zap.AddStacktrace(zap.NewAtomicLevelAt(zap.ErrorLevel)) //设置哪一个等级level打印栈空间
	//stack := zap.AddStacktrace(zap.NewAtomicLevelAt(zap.ErrorLevel))//设置哪一个等级level打印栈空间
	logger, _ = config.Build(skip, stack)
	supar = logger.Sugar()
}

func InitLogger(serverName string, serverId int32, logPath string, level Level) {
	if logger != nil {
		logger.Sync()
		logger = nil
		supar = nil
	}
	if logPath == "" {
		logPath = "./log"
	}
	fileName := fmt.Sprintf("%s/%s_%d.log", logPath, serverName, serverId)
	fmt.Println(fileName)
	hook := lumberjack.Logger{
		Filename:   fileName, // 日志文件路径
		MaxSize:    2048,     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,       // 日志文件最多保存多少个备份
		MaxAge:     7,        // 文件最多保存多少天
		Compress:   true,     // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999999999"),
		EncodeDuration: zapcore.MillisDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}

	logLevel := zapcore.DebugLevel
	switch level {
	case DebugLevel:
		logLevel = zapcore.DebugLevel
	case InfoLevel:
		logLevel = zapcore.InfoLevel
	case WarnLevel:
		logLevel = zapcore.WarnLevel
	case ErrorLevel:
		logLevel = zapcore.ErrorLevel
	case DPanicLevel:
		logLevel = zapcore.DPanicLevel
	case PanicLevel:
		logLevel = zapcore.PanicLevel
	case FatalLevel:
		logLevel = zapcore.FatalLevel
	}

	atomLevel := zap.NewAtomicLevelAt(logLevel)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomLevel)

	caller := zap.AddCaller()
	skip := zap.AddCallerSkip(1)
	develop := zap.Development()
	stack := zap.AddStacktrace(zap.NewAtomicLevelAt(zap.ErrorLevel)) //设置哪一个等级level打印栈空间
	//filed := zap.Fields(zap.Int32(serverName, serverId))
	logger = zap.New(core, caller, skip, develop, stack) //, filed)
	supar = logger.Sugar()
}

func SyncLogger() {
	if logger != nil {
		logger.Sync()
	}
}
