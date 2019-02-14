package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose/utils"
	"net/http"
)

func Demo() utils.ApiReturn {
	return utils.SuccessReturn("成功")
}

func Demo2(c *gin.Context) {
	c.String(http.StatusOK, "test success")
}
