package main

import (
	"context"
	"photo-share/back/infrastructure"
	"photo-share/back/infrastructure/restexe/apidef/server"
	"photo-share/back/sharelib/domain/logging"
	repolib "photo-share/back/sharelib/domain/repository"
	"photo-share/back/sharelib/domain/secret"
	"photo-share/back/sharelib/mysql"
	timezone "photo-share/back/sharelib/timezone"
	"photo-share/back/usecase/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	timezone.SetTimeZoneAsiaTokyo()

	config := infrastructure.Configs{
		LogLevel: "debug",
		GinMode:  gin.DebugMode,
		DBName:   "photo-share",
	}

	// Logger初期化
	logging.InitLoggerWithLevel(logging.NewLogLevel(config.LogLevel))
	logging.InitLoggerEncoding(logging.CONSOLE)
	// Loggerの後始末処理
	defer logging.Sync()

	// DB作成処理
	secret := secret.NewLocalSecret("root", "my-secret-pw", "localhost:3306")

	db := mysql.NewMysqlDB(config.DBName, secret)
	err := db.Open()
	if err != nil {
		logging.Error("DB Openエラー")
		return
	}
	defer db.Close()

	// リポジトリプロバイダを作成
	ctx := context.Background()
	repository := infrastructure.NewResourceProvider(config, db.GetDB(), ctx)

	// コントローラーの登録
	controller := handler.NewServer(repository)

	gin.SetMode(config.GinMode)
	r := gin.Default()
	// ログの出力設定
	r.Use(logging.LogRequestResponseMiddleware)
	// トランザクション設定
	r.Use(repolib.DefaultTransactionMiddleware(repository))

	server.RegisterHandlers(r, controller)
	r.Run(":9993")
}
