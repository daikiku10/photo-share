package repository

import (
	"context"
	"database/sql"
	"photo-share/back/sharelib/domain/logging"
)

type Transactor interface{}

// TransactorImpl Transactor実装モジュール
type transactorImpl struct {
	db   *sql.DB
	tx   *sql.Tx
	ctx  context.Context
	opts *sql.TxOptions
}

// NewTransactorWithOpts Transactorを構築する
// トランザクション開始するかを指定できる
// トランザクション開始する場合は、オプションを指定できる
func NewTransactorWithOpts(
	db *sql.DB,
	ctx context.Context,
	toOpen bool,
	opts *sql.TxOptions,
) (trns Transactor, err error) {
	var tx *sql.Tx
	if toOpen {
		tx, err = db.BeginTx(ctx, opts)
		if err != nil {
			return
		}
		logging.Debug("transaction started")
	}

	trns = &transactorImpl{
		db:  db,
		tx:  tx,
		ctx: ctx,
	}
	return
}
