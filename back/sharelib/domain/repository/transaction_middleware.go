package repository

import (
	"database/sql"
	"net/http"
	"photo-share/back/sharelib/domain/logging"

	"github.com/gin-gonic/gin"
)

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
	return func(c *gin.Context) {
		// デフォルトトランザクションを開始するか判断
		toOpen, opts := filter(c.Request.Method, c.Request.URL.Path)

		// デフォルトTransactorの構築とトランザクション開始
		trns, err := provider.NewTransactorWithFlag(opts, toOpen)
		if err != nil {
			logging.Error("start transaction failed", logging.Err(err))
			c.AbortWithStatus(http.StatusInternalServerError)
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
	}
}
