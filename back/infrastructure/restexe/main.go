package main

import (
	"photo-share/back/infrastructure"
	"photo-share/back/sharelib/domain/logging"
	timezone "photo-share/back/sharelib/timezone"
)

func main() {
	timezone.SetTimeZoneAsiaTokyo()

	config := infrastructure.Configs{
		LogLevel: "debug",
	}

	// Logger初期化
	logging.InitLoggerWithLevel(logging.NewLogLevel(config.LogLevel))
	logging.InitLoggerEncoding(logging.CONSOLE)
	// Loggerの後始末処理
	defer logging.Sync()
}
