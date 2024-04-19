package repository

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// TransactorProvider Transactorを提供するインターフェース
// 今回Backendはginを前提に構築しており、次の考慮のため、Transactorはgin.Contextに格納する
// * ResourceProviderはリクエスト館で共有されるため、ResourceProviderには格納できない。
// * 複数Repositoryの操作を1つのトランザクションで制御することを可能にするため、Repositoryに管理させるのは適当ではない。
// * ginではリクエストごとにgin.Contextが生成されるため、gin.Contextに格納する
type TransactorProvider interface {
	// NewTransactor Transactorを新規構築する
	// NewTransactor(opts *sql.TxOptions) (Transactor, error)

	// GetCurrentTransactor gin.Contextに格納されたデフォルトTransactorを取得する
	GetCurrentTransactor(ctx *gin.Context) Transactor

	// NewTransactorWithFlag トランザクション開始するか判断してTransactorを構築する
	NewTransactorWithFlag(opts *sql.TxOptions, toOpen bool) (Transactor, error)
}

// NewTransactorProvider TransactorProviderを構築する
// 各ResourceProviderが内包することを想定している
func NewTransactorProvider(db *sql.DB, ctx context.Context) TransactorProvider {
	return &transactorProviderImpl{
		db:   db,
		ctx:  ctx,
		exec: nil,
	}
}

// TransactorProviderImpl TransactorProvider実装モジュール
type transactorProviderImpl struct {
	db   *sql.DB
	ctx  context.Context
	exec Transactor
}

// GetCurrentTransactor gin.Contextに格納されたデフォルトTransactorを取得する
func (eq *transactorProviderImpl) GetCurrentTransactor(ctx *gin.Context) Transactor {
	value, found := ctx.Get(ContextKeyTransactor)
	if found {
		// 型アサーション valueがTransactorインターフェースを実装しているかのチェック
		exec, ok := value.(Transactor)
		if ok {
			return exec
		}
	}
	return nil
}

// NewTransactorWithFlag トランザクション開始するか判断してTransactorを構築する
func (eq *transactorProviderImpl) NewTransactorWithFlag(
	opts *sql.TxOptions,
	toOpen bool,
) (trns Transactor, err error) {
	return NewTransactorWithOpts(eq.db, eq.ctx, toOpen, opts)
}
