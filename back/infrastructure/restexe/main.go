package main

import (
	"fmt"
	"photo-share/back/infrastructure"
	"photo-share/back/sharelib/domain/logging"
	timezone "photo-share/back/sharelib/timezone"
)

func main() {
	timezone.SetTimeZoneAsiaTokyo()

	config := infrastructure.Configs{
		LogLevel: "debug",
	}
	fmt.Println(config)

	// Logger初期化
	logging.InitLoggerWithLevel(logging.NewLogLevel(config.LogLevel))
}
