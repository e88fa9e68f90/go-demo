package loggerop

import (
	zap_log "github.com/kiwi633/go-demo/zap-log"
	"go.uber.org/zap"
)

func Info(msg string, fields ...zap.Field) {
	zap_log.GetLogger(*zap_log.LOGContext).Info(msg, fields...)
}
func GetTid() string {
	return zap_log.GetLogger(*zap_log.LOGContext).GetTid()
}
func Error(msg string, fields ...zap.Field) {
	zap_log.GetLogger(*zap_log.LOGContext).Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	zap_log.GetLogger(*zap_log.LOGContext).Debug(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
	zap_log.GetLogger(*zap_log.LOGContext).Warn(msg, fields...)
}
