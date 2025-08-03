package zap_log

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// Logger结构体包装zap logger
type Logger struct {
	*zap.Logger
	tid string
}

// 从context获取logger（带TID）
func GetLogger(ctx context.Context) *Logger {
	tid, _ := ctx.Value(TIDKey).(string)
	return &Logger{
		Logger: Log,
		tid:    tid,
	}
}

// 全局logger实例
var Log *zap.Logger

// TID上下文键
type ContextKey string

const TIDKey ContextKey = "tid"

// 初始化zap logger
func InitLogger() {
	// 配置日志输出格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 日志文件配置
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    100, // MB
		MaxBackups: 3,
		MaxAge:     7, // days
		Compress:   true,
	})

	// 控制台输出
	consoleDebugging := zapcore.Lock(zapcore.AddSync(zapcore.AddSync(zapcore.Lock(os.Stdout))))

	// 创建core
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writeSyncer, zapcore.DebugLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), consoleDebugging, zapcore.DebugLevel),
	)

	// 创建logger
	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// 生成trace id
func generateTraceID() string {
	tid, _ := uuid.NewUUID()
	//println("=========================================" + tid)
	//node, err := snowflake.NewNode(0)
	//if err != nil {
	//	panic(err)
	//}
	//tid := node.Generate().String()
	return fmt.Sprintf("%s", tid)
}

var LOGContext *context.Context

// 中间件：添加TID到上下文和日志
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成trace id
		tid := generateTraceID()

		// 将tid添加到context
		ctx := context.WithValue(c.Request.Context(), TIDKey, tid)
		c.Request = c.Request.WithContext(ctx)
		ctx2 := c.Request.Context()
		LOGContext = &ctx2
		// 添加到响应头
		c.Header("X-Trace-ID", tid)
		c.Next()
	}
}

// Info方法，自动添加TID
func (l *Logger) GetTid() string {
	return l.tid
}
func (l *Logger) Info(msg string, fields ...zap.Field) {
	if l.tid != "" {
		fields = append(fields, zap.String("tid", l.tid))
	}
	l.Logger.Info(msg, fields...)
}

// Error方法，自动添加TID
func (l *Logger) Error(msg string, fields ...zap.Field) {
	if l.tid != "" {
		fields = append(fields, zap.String("tid", l.tid))
	}
	l.Logger.Error(msg, fields...)
}

// Debug方法，自动添加TID
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	if l.tid != "" {
		fields = append(fields, zap.String("tid", l.tid))
	}
	l.Logger.Debug(msg, fields...)
}

// Warn方法，自动添加TID
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	if l.tid != "" {
		fields = append(fields, zap.String("tid", l.tid))
	}
	l.Logger.Warn(msg, fields...)
}

// 从gin context获取logger（带TID）
func GetLoggerFromGin(c *gin.Context) *Logger {
	tid, _ := c.Request.Context().Value(TIDKey).(string)
	return &Logger{
		Logger: Log,
		tid:    tid,
	}
}

// 全局logger函数，用于没有context的场景
func GetGlobalLogger() *Logger {
	return &Logger{
		Logger: Log,
		tid:    "",
	}
}

// 业务函数示例
func someBusinessLogic(ctx context.Context) {
	logger := GetLogger(ctx)
	logger.Info("开始处理业务逻辑")

	// 模拟一些操作
	time.Sleep(100 * time.Millisecond)

	logger.Info("业务逻辑处理成功")
}
