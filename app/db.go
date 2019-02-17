package app

import (
	"database/sql"
	"strings"

	"github.com/go-xorm/core"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gobuffalo/packr"
	migrate "github.com/rubenv/sql-migrate"
)

var DB *xorm.Engine

func InitDB() error {
	var err error
	DB, err = xorm.NewEngine("mysql", Config.DSN)
	if err != nil {
		return err
	}
	DB.SetMapper(core.GonicMapper{})
	return nil
}

func Migrate(sources packr.Box, up bool) (n int, err error) {
	dsn := Config.DSN
	if !strings.Contains(dsn, "parseTime") {
		if strings.Contains(dsn, "?") {
			dsn += "&parseTime=true"
		} else {
			dsn += "?parseTime=true"
		}
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	defer db.Close()
	migrations := &migrate.PackrMigrationSource{
		Box: sources,
	}
	migrate.SetTable("schema_migration")
	if up {
		return migrate.Exec(db, "mysql", migrations, migrate.Up)
	} else {
		return migrate.Exec(db, "mysql", migrations, migrate.Down)
	}
}
