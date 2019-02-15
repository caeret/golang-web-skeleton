package app

import "github.com/go-xorm/xorm"

type Container interface {
	DB() *xorm.Engine
}
