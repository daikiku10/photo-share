package main

import (
	"fmt"
	"photo-share/back/infrastructure"
	"photo-share/back/sharelib/domain/logging"
	timezone "photo-share/back/sharelib/timezone"
	"photo-share/back/usecase/handler"
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

	// リポジトリプロバイダを作成
	repository := infrastructure.NewResourceProvider(config)

	// コントローラーの登録
	controller := handler.NewServer(repository)
	fmt.Println(controller)

}
