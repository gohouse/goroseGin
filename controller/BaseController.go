package controller

import (
	"github.com/gohouse/gorose"
	"github.com/gohouse/gupiao/bootstrap"
)

func M(table string) *gorose.Session {
	db := NewConnectionInstance()
	return db.Table(table)
}

func NewConnectionInstance() *gorose.Session {
	return bootstrap.GetBooterInstance().Connection.NewSession()
}
