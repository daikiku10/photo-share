package repository

import (
	"context"
	"database/sql"
	"fmt"
	"photo-share/back/sharelib/domain/logging"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Transactor interface {
	// IsOpen トランザクション中か判断する
	IsOpen() bool
	// Rollback トランザクションをロールバックする
	Rollback() error
	// Commit トランザクションをコミットする
	Commit() error
	// Context 紐づくContextを返す
	Context() context.Context
	// Get 有効なContextTransactorを返す
	// トランザクション開始前ならsql.DBそのもの、トランザクション中はsql.Txを返す
	Get() boil.ContextExecutor
}

// TransactorImpl Transactor実装モジュール
type transactorImpl struct {
	db   *sql.DB
	tx   *sql.Tx
	ctx  context.Context
	opts *sql.TxOptions
}

// IsOpen トランザクション中ならtrue、そうでなければfalse
func (trns *transactorImpl) IsOpen() bool {
	return trns.tx != nil
}

// Rollback トランザクションをロールバックする
func (trns *transactorImpl) Rollback() (err error) {
	if !trns.IsOpen() {
		err = fmt.Errorf("トランザクションは開始していません")
		return
	}

	err = trns.tx.Rollback()
	if err == nil {
		logging.Debug("transaction rollback")
		trns.tx = nil
	}
	return
}

// Commit トランザクションをコミットする
func (trns *transactorImpl) Commit() error {
	if !trns.IsOpen() {
		err := fmt.Errorf("トランザクションは開始していません。")
		return err
	}

	err := trns.tx.Commit()
	if err != nil {
		return err
	}

	logging.Debug("transaction committed")
	return nil
}

// Context 紐づくContextを返す
func (trns *transactorImpl) Context() context.Context {
	return trns.ctx
}

// Get 有効なContextTransactorを返す
// トランザクション開始前ならsql.DBそのもの、トランザクション中はsql.Txを返す
func (trns *transactorImpl) Get() boil.ContextExecutor {
	if trns.IsOpen() {
		return trns.tx
	} else {
		return trns.db
	}
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
