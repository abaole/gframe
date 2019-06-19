package gframe

import (
	"github.com/abaole/gframe/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDBManger(opts *db.Options) error {
	db.NewMysql(
		db.WithHost(opts.Host),
		db.WithDbName(opts.DbName),
		db.WithPort(opts.Port),
		db.WithUser(opts.User),
		db.WithPassword(opts.Password),
		db.WithMaxIdleConn(opts.MaxIdleConn),
		db.WithMaxOpenConn(opts.MaxOpenConn),
		db.WithMaxLifeTime(opts.ConnMaxLifeTime),
		db.WithCharset(opts.Charset),
		db.WithIsLog(opts.IsLog),
	)
	return nil
}

func closeDb() error {
	return db.Close()
}
