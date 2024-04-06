package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Loggerインスタンス
// Newで構築された内容をここで保持し、logger()で取得できるようにする。
var _option *LoggerOption

type LogLevel zapcore.Level

type LoggerOption struct {
	Level    LogLevel
	Encoding string
}

const (
	DebugLevel = LogLevel(zap.DebugLevel)
	InfoLevel  = LogLevel(zap.InfoLevel)
	WarnLevel  = LogLevel(zap.WarnLevel)
)

// NewLogLevel 文字列からログレベルを取得する
func NewLogLevel(level string) LogLevel {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	default:
		return InfoLevel
	}
}

func option() *LoggerOption {
	if _option == nil {
		_option = &LoggerOption{
			Level:    LogLevel(DebugLevel),
			Encoding: "json",
		}
	}
	return _option
}

func InitLoggerWithLevel(logLevel LogLevel) {
	option().Level = logLevel
}
