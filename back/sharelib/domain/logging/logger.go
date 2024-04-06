package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Loggerインスタンス
// Newで構築された内容をここで保持し、logger()で取得できるようにする。
var _logger *zap.Logger
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

	JSON    = "json"
	CONSOLE = "console"
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
			Encoding: JSON,
		}
	}
	return _option
}

func InitLoggerWithLevel(logLevel LogLevel) {
	option().Level = logLevel
}

func InitLoggerEncoding(encode string) {
	option().Encoding = encode
}

// Sync 出力されていないログを処理する
// ログの消失を防ぐため、main関数のdeferで呼び出すこと
func Sync() {
	if _logger == nil {
		_logger.Sync()
	}
}
