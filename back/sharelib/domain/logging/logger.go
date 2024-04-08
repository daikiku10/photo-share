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

func InitLogger() {
	option := option()

	config := zap.NewDevelopmentConfig()
	config.Encoding = option.Encoding
	config.Level = zap.NewAtomicLevelAt(zapcore.Level(option.Level))

	if option.Encoding == CONSOLE {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	_logger, _ = config.Build(
		zap.AddCallerSkip(1),
	)
}

// Sync 出力されていないログを処理する(蓄積されたログを全て吐き出す)
// ログの消失を防ぐため、main関数のdeferで呼び出すこと
func Sync() {
	if _logger != nil {
		_logger.Sync()
	}
}

func logger() *zap.Logger {
	if _logger == nil {
		InitLogger()
	}
	return _logger
}

type LogValue struct {
	Title string
	Value any
}

// Var 変数のログ出力設定
func Var(t string, v any) LogValue {
	return LogValue{
		Title: t,
		Value: v,
	}
}

// Debug DEBUGレベルでログ出力する
func Debug(message string, values ...LogValue) {
	var params []zapcore.Field
	for _, kv := range values {
		params = append(params, zap.Any(kv.Title, kv.Value))
	}
	logger().Debug(message, params...)
}
