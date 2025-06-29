package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"gvadmin_core/config"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func Instance() *zap.Logger {
	if logger == nil {
		InitLog()
	}
	return logger
}

func InitLog() {
	hook := lumberjack.Logger{
		Filename:   path.Join(config.Instance().ZapLog.Director, time.Now().Format("2006-01-02")+".log"), // 日志文件路径
		MaxSize:    100,                                                                                  // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 50,                                                                                   // 日志文件最多保存多少个备份
		MaxAge:     30,                                                                                   // 文件最多保存多少天
		Compress:   true,                                                                                 // 是否压缩
		LocalTime:  true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serverName", E.ProjectName))
	// 构造日志
	logger = zap.New(core, caller, development) //, filed)
	//logger.Info("log 初始化成功")
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
