package repository

import (
	"context"
	"database/sql"
)

// TransactorProvider Transactorを提供するインターフェース
// 今回Backendはginを前提に構築しており、次の考慮のため、Transactorはgin.Contextに格納する
// * ResourceProviderはリクエスト館で共有されるため、ResourceProviderには格納できない。
// * 複数Repositoryの操作を1つのトランザクションで制御することを可能にするため、Repositoryに管理させるのは適当ではない。
// * ginではリクエストごとにgin.Contextが生成されるため、gin.Contextに格納する
type TransactorProvider interface {
	// NewTransactor Transactorを新規構築する
	// NewTransactor(opts *sql.TxOptions) (Transactor, error)
	// NewTransactorWithFlag トランザクション開始するか判断してTransactorを構築する
	NewTransactorWithFlag(opts *sql.TxOptions, toOpen bool) (Transactor, error)
}

// TransactorProviderImpl TransactorProvider実装モジュール
type transactorProviderImpl struct {
	db   *sql.DB
	ctx  context.Context
	exec Transactor
}

// NewTransactorWithFlag トランザクション開始するか判断してTransactorを構築する
func (eq *transactorProviderImpl) NewTransactorWithFlag(
	opts *sql.TxOptions,
	toOpen bool,
) (trns Transactor, err error) {
	return NewTransactorWithOpts(eq.db, eq.ctx, toOpen, opts)
}
