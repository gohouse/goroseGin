package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/goroseGin/bootstrap"
	"net/http"
	"time"
)

func AdminRun() {
	var router gin.IRouter = bootstrap.GetBooterInstance().Router

	admin := router.Group("admin")

	admin.GET("/", func(c *gin.Context) {
		format := "2006-01-02 15:04:05"
		c.String(http.StatusOK, "Admin works "+time.Now().Format(format))
	})

	// 用户管理
	user := router.Group("user")
	user.Any("getlist")
	user.Any("addoredit")
	user.Any("delete")

	// 统计
	calc := router.Group("calc")
	// 统计周期:今天,昨天,本周,上周,本月,上月,本年度, 上年度
	// 统计维度:进出库材料分类统计(数量,金额,盈亏)
	calc.Any("/")

	// 计划任务
	plan := router.Group("plan")
	// 晚上0点统计 昨天,上周,上月,上年度,本周,本月,本年度 的数据
	// 如果是周一就统计上周,如果是1号就统计上月,如果是元旦就统计去年
	plan.Any("calc")
}