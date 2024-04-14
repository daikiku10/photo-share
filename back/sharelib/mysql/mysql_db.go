package mysql

import (
	"database/sql"
	"photo-share/back/sharelib/domain/logging"
	"photo-share/back/sharelib/domain/secret"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	dbName string
	secret secret.SecretInterface
	db     *sql.DB
}

// NewMysqlDB MysqlDB構造体を構築する
func NewMysqlDB(dbName string, secret secret.SecretInterface) *MysqlDB {
	return &MysqlDB{
		dbName: dbName,
		secret: secret,
		db:     nil,
	}
}

func (md *MysqlDB) GetDB() *sql.DB {
	return md.db
}

func (md *MysqlDB) Open() error {
	user := md.secret.User()
	password := md.secret.Password()
	host := md.secret.Host()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	config := mysqlDriver.Config{
		DBName:               md.dbName,
		User:                 user,
		Passwd:               password,
		Addr:                 host,
		Net:                  "tcp",
		ParseTime:            true,
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	logging.Debug("sql open前", logging.Var("user", user), logging.Var("host", host))
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return err
	}

	md.db = db
	logging.Debug("Open成功", logging.Var("MysqlDB", md), logging.Var("db", db))
	return nil
}

func (md *MysqlDB) Close() error {
	return md.db.Close()
}
