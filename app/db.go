package app

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/gobuffalo/packr"
	migrate "github.com/rubenv/sql-migrate"
)

var DB *xorm.Engine

func InitDB(logger Logger) error {
	var err error
	DB, err = xorm.NewEngine("mysql", Config.DSN)
	if err != nil {
		return err
	}
	DB.SetMapper(core.GonicMapper{})
	DB.SetLogger(newSimpleLogger(logger, core.LOG_DEBUG))
	DB.ShowExecTime(true)
	DB.ShowSQL(true)
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

type simpleLogger struct {
	logger  Logger
	level   core.LogLevel
	showSQL bool
}

func newSimpleLogger(logger Logger, l core.LogLevel) *simpleLogger {
	return &simpleLogger{
		logger: logger,
		level:  l,
	}
}

func (s *simpleLogger) Error(v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.logger.Error(fmt.Sprint(s.format(v)...))
	}
}

func (s *simpleLogger) Errorf(format string, v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.logger.Error(fmt.Sprintf(format, s.format(v)...))
	}
}

func (s *simpleLogger) Debug(v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.logger.Debug(fmt.Sprint(s.format(v)...))
	}
}

func (s *simpleLogger) Debugf(format string, v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.logger.Debug(fmt.Sprintf(format, s.format(v)...))
	}
}

func (s *simpleLogger) Info(v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.logger.Info(fmt.Sprint(s.format(v)...))
	}
	return
}

func (s *simpleLogger) Infof(format string, v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.logger.Info(fmt.Sprintf(format, s.format(v)...))
	}
	return
}

func (s *simpleLogger) Warn(v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.logger.Warn(fmt.Sprint(s.format(v)...))
	}
	return
}

func (s *simpleLogger) Warnf(format string, v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.logger.Warn(fmt.Sprintf(format, s.format(v)...))
	}
	return
}

func (s *simpleLogger) Level() core.LogLevel {
	return s.level
}

func (s *simpleLogger) SetLevel(l core.LogLevel) {
	s.level = l
	return
}

func (s *simpleLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

func (s *simpleLogger) IsShowSQL() bool {
	return s.showSQL
}

func (s *simpleLogger) format(v []interface{}) []interface{} {
	tmpV := make([]interface{}, len(v))
	copy(tmpV, v)
	if len(tmpV) >= 2 {
		if slice, ok := tmpV[1].([]interface{}); ok {
			tmpSlice := make([]interface{}, len(slice))
			copy(tmpSlice, slice)
			for i, item := range tmpSlice {
				switch raw := item.(type) {
				case []byte:
					tmpSlice[i] = fmt.Sprintf("<%d bytes>", len(raw))
					tmpV[1] = tmpSlice
				}
			}
		}
	}
	return tmpV
}
