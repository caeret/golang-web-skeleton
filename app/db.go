package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

func InitDB() error {
	var err error
	DB, err = xorm.NewEngine("mysql", Config.DSN)
	return err
}
