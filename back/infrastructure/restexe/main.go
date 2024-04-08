package main

import (
	"photo-share/back/infrastructure"
	"photo-share/back/infrastructure/restexe/apidef/server"
	"photo-share/back/sharelib/domain/logging"
	timezone "photo-share/back/sharelib/timezone"
	"photo-share/back/usecase/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	timezone.SetTimeZoneAsiaTokyo()

	config := infrastructure.Configs{
		LogLevel: "debug",
		GinMode:  gin.DebugMode,
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

	gin.SetMode(config.GinMode)
	r := gin.Default()
	// ログの出力設定
	r.Use(logging.LogRequestResponseMiddleware)
	// TODO:トランザクション設定

	server.RegisterHandlers(r, controller)
	r.Run(":9993")
}
