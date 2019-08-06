package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/abaole/gframe/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func NewMysql(opts ...Option) (err error) {
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
	c.SingularTable(true)
	c.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	c.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	SetMysql(c)

	return
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("created_at"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("updated_at"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("updated_at", time.Now())
	}
}

func Close() error {
	return db.Close()
}

var migratesOnce sync.Once

// migrate migrates database schemas ...
func Migrate(models []interface{}) error {
	migratesOnce.Do(func() {
		err := db.AutoMigrate(models...).Error
		if err != nil {
			logger.Panicf("auto migrate db table error: %v", err)
		}
	})

	return nil
}

func SetMysql(gDB *gorm.DB) {
	db = gDB
}

func GetDB() *gorm.DB {
	return db
}
