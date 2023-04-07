package main

import (
	"errors"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

const (
	SERVICE_NAME = "demo_log"
	MODE         = "pro"
)

var _onceInit sync.Once
var log *zap.Logger
var writer zapcore.WriteSyncer

func main() {
	/*core := zapcore.NewCore(getEncoder(), getWriteSyncer(), zapcore.DebugLevel)
	zap.New(core, zap.AddCaller())*/
	_onceInit.Do(func() {
		fmt.Println("init log")
		writer = getStdLogWriter()
		//if MODE == "pro" {
		//	writer = getFileLogWriter()
		//}
		// 开启开发模式，堆栈跟踪
		caller := zap.AddCaller()
		// 开启文件及行号
		development := zap.Development()
		// 设置初始化字段,如：添加一个服务器名称
		filed := zap.Fields(zap.String("serviceName", SERVICE_NAME))
		log = zap.New(
			zapcore.NewCore(getEncoder(), writer, zapcore.DebugLevel),
			caller,
			development,
			filed,
		)
	})
	for i := 0; i < 100; i++ {
		log.Error("this is error message", zap.Error(errors.New("error test")))
		log.Info("this is info message", zap.Any("a", "any value"))
	}
	time.Sleep(100 * time.Second)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 修改时间编码器
	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.CallerKey = "file"
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	if MODE == "pro" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getFileLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/" + SERVICE_NAME + ".log", // 日志文件的位置；
		MaxSize:    1,                                // 在进行切割之前，日志文件的最大大小（以MB为单位）；
		MaxBackups: 10,                               // 保留旧文件的最大个数；
		MaxAge:     30,                               // 保留旧文件的最大天数；
		Compress:   false,                            // 是否压缩/归档旧文件；
	}
	return zapcore.AddSync(lumberJackLogger)

}

func getStdLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}
