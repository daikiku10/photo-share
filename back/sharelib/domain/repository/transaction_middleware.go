package repository

import (
	"database/sql"
	"net/http"
	"photo-share/back/sharelib/domain/logging"

	"github.com/gin-gonic/gin"
)

const ContextKeyTransactor = "_transactor"

type TransactionFilter func(string, string) (bool, *sql.TxOptions)

// DefaultTransactionMiddleware デフォルトトランザクションを制御する
func DefaultTransactionMiddleware(
	provider TransactorProvider,
) gin.HandlerFunc {
	return DefaultTransactionMiddlewareWithFilter(
		provider,
		func(method string, path string) (bool, *sql.TxOptions) {
			return true, nil
		},
	)
}

func DefaultTransactionMiddlewareWithFilter(
	provider TransactorProvider,
	filter TransactionFilter,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// デフォルトトランザクションを開始するか判断
		toOpen, opts := filter(ctx.Request.Method, ctx.Request.URL.Path)

		// デフォルトTransactorの構築とトランザクション開始
		trns, err := provider.NewTransactorWithFlag(opts, toOpen)
		if err != nil {
			logging.Error("start transaction failed", logging.Err(err))
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer func() {
			if trns.IsOpen() {
				logging.Debug("force rollback...")
				err = trns.Rollback()
				if err != nil {
					logging.Error("rollback failed", logging.Err(err))
				}
			}
		}()

		// gin.Contextにトランザクションを格納
		ctx.Set(ContextKeyTransactor, trns)

		ctx.Next()

		// トランザクション終了
		if trns.IsOpen() {
			// 処理成功はCommit, 失敗はRollback
			if ctx.Writer.Status() == http.StatusOK {
				logging.Debug("commit...")
				err = trns.Commit()
				if err != nil {
					logging.Error("commit failed", logging.Err(err))
				}
			} else {
				logging.Debug("rollback...")
				err = trns.Rollback()
				if err != nil {
					logging.Error("rollback failed", logging.Err(err))
				}
			}
		}
	}
}
