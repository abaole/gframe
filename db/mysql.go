package db

import (
	"fmt"
	"github.com/abaole/gframe/logger"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type DB struct {
	Db *gorm.DB
}

func NewMysql(opts ...Option) (db *DB, err error) {
	o := &Options{
		Host:            DF_Host,
		DbName:          DF_DbName,
		Port:            DF_Port,
		User:            DF_User,
		Password:        DF_Password,
		MaxIdleConn:     DF_MaxIdleConn,
		MaxOpenConn:     DF_MaxOpenConn,
		ConnMaxLifeTime: DF_ConnMaxLifeTime,
		Charset:         DF_Charset,
		IsLog:           DF_IsLog,
	}
	for _, opt := range opts {
		opt(o)
	}

	connHost := fmt.Sprintf("tcp(%s:%d)", o.Host, o.Port)
	s := fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=True&loc=Local", o.User, o.Password, connHost, o.DbName, o.Charset)

	c, err := gorm.Open("mysql", s)
	if err != nil {
		panic("mysql init error")
	}

	c.LogMode(o.IsLog)
	c.DB().SetMaxIdleConns(o.MaxIdleConn)
	c.DB().SetMaxOpenConns(o.MaxOpenConn)
	c.DB().SetConnMaxLifetime(time.Second * time.Duration(o.ConnMaxLifeTime))
	db.Db = c
	SetMysql(c)

	return
}

func (db *DB) Close() {
	db.Db.Close()
}

var migratesOnce sync.Once

// migrate migrates database schemas ...
func (db *DB) Migrate(models []interface{}) error {
	migratesOnce.Do(func() {
		err := db.Db.AutoMigrate(models...).Error
		if err != nil {
			logger.Panicf("auto migrate db table error: %v", err)
		}
	})

	return nil
}

func SetMysql(db *gorm.DB) {
	Db = db
}
