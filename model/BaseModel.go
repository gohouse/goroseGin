package model

import (
	"github.com/gohouse/gorose"
	"github.com/gohouse/goroseGin/bootstrap"
)

func M(table interface{}) *gorose.Session {
	return NewDBConnection().Table(table)
}

func NewDBConnection() *gorose.Session {
	return bootstrap.GetBooterInstance().Connection.NewSession()
}
