package orm

import (
	"time"

	"github.com/daoraimi/dagger/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	initOrm()
}

func R() *gorm.DB {
	if db == nil {
		initOrm()
	}
	return db
}

func initOrm() {
	var err error
	dbConfig := &mysql.Config{
		User:                 config.GetString("mysql.user"),
		Passwd:               config.GetString("mysql.password"),
		Net:                  "tcp",
		Addr:                 config.GetString("mysql.addr"),
		DBName:               config.GetString("mysql.db"),
		ParseTime:            true,
		Loc:                  time.Local,
		MaxAllowedPacket:     config.GetInt("mysql.maxAllowedPacket"),
		AllowNativePasswords: true,
		Params: map[string]string{
			"charset":   config.GetString("mysql.charset"),
			"collation": config.GetString("mysql.collation"),
		},
	}
	db, err = gorm.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	if config.GetBool("debug") {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.GetInt("mysql.maxIdleConns"))
	db.DB().SetMaxOpenConns(config.GetInt("mysql.maxOpenConns"))
}

func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
