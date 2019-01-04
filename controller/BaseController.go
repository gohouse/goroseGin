package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"
	"github.com/gohouse/gorose/utils"
	"github.com/gohouse/goroseGin/bootstrap"
	"net/http"
)

func M(table string) *gorose.Session {
	db := NewConnectionInstance()
	return db.Table(table)
}

func NewConnectionInstance() *gorose.Session {
	return bootstrap.GetBooterInstance().Connection.NewSession()
}

func Demo() utils.ApiReturn {
	return utils.SuccessReturn("成功")
}

func Demo2(c *gin.Context) {
	c.String(http.StatusOK, "test success")
}
