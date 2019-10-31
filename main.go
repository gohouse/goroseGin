package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose/v2"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"sync"
)

func main() {
	// 加载gorose
	BootGorose()

	// 初始化用户表
	UserInit()

	// 加载gin
	BootGin()
}

// =============== 用户管理部分 ===============
// UserAdd 用户增加
func UserAdd(c *gin.Context) {
	// 用户名
	username := c.Query("username")
	// 年龄
	age := c.DefaultQuery("age", "0")
	// 接受参数组装入库数据
	var data = make(map[string]interface{})
	if username == "" {
		c.JSON(http.StatusOK, FailReturn("用户名不能为空"))
		return
	}
	data["username"] = username
	if age != "0" {
		data["age"] = age
	}
	// 执行入库
	affected_rows, err := DB().Table("users").Data(data).Insert()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	// api接口返回
	c.JSON(http.StatusOK, SuccessReturn(affected_rows))
}

// UserDelete 用户删除
func UserDelete(c *gin.Context) {
	// 按主键删除
	uid := c.Query("uid")
	if uid == "" {
		c.JSON(http.StatusOK, FailReturn("用户id不能为空"))
		return
	}
	affected_rows, err := DB().Table("users").Where("uid", uid).Delete()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(affected_rows))
}

// UserEdit 用户编辑
func UserEdit(c *gin.Context) {
	// 按照主键编辑
	uid := c.Query("uid")
	username := c.DefaultQuery("username", "")
	age := c.DefaultQuery("age", "0")
	if uid == "" {
		c.JSON(http.StatusOK, FailReturn("用户id不能为空"))
		return
	}
	if username == "" && age == "0" {
		c.JSON(http.StatusOK, FailReturn("未修改"))
		return
	}

	var data = make(map[string]interface{})
	if username != "" {
		data["username"] = username
	}
	if age != "0" {
		data["age"] = age
	}
	// 执行入库操作
	affected_rows, err := DB().Table("users").Where("uid", uid).Data(data).Update()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(affected_rows))
}

// UserList 获取用户列表
func UserList(c *gin.Context) {
	// 默认查询50条数据
	userList, err := DB().Table("users").OrderBy("uid desc").Limit(50).Get()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(userList))
}

// =============== 初始化部分 ============
var once sync.Once
var engin *gorose.Engin

//BootGorose 初始化gorose, 单例模式
func BootGorose() {
	var err error
	once.Do(func() {
		engin, err = gorose.Open(&gorose.Config{
			Driver: "sqlite3",
			Dsn:    "db.sqlite",
		})
		if err != nil {
			panic(err.Error())
		}
	})
}

//UserInit 初始化用户表
func UserInit() {
	dbSql := `CREATE TABLE IF NOT EXISTS "users" (
	 "uid" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	 "username" TEXT NOT NULL default "",
	 "age" integer NOT NULL default 0
)`
	affected_rows, err := DB().Execute(dbSql)
	if err != nil {
		panic(err.Error())
	}
	if affected_rows == 0 {
		return
	}
}

//BootGin 初始化gin
func BootGin() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

//	router.GET("/", func(c *gin.Context) {
//		c.Header("Content-Type", "text/html; charset=utf-8")
//		c.String(http.StatusOK,
//			`<br><br><center>
//<h1>欢迎来到golang入门用户管理api服务系统</h1>
//</center>`)
//	})
	router.Use(Cors())
	router.GET("/UserAdd", UserAdd)
	router.GET("/UserList", UserList)
	router.GET("/UserEdit", UserEdit)
	router.GET("/UserDelete", UserDelete)

	// 静态文件服务
	router.Static("/html","./")

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

//============= 工具函数部分 ==============
// SuccessReturn api正确返回函数
func SuccessReturn(msg interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = msg
	res["code"] = http.StatusOK
	res["msg"] = "success"

	return res
}

// FailReturn api错误返回函数
func FailReturn(msg interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = ""
	res["code"] = http.StatusBadRequest
	res["msg"] = msg

	return res
}

// DB orm快捷使用函数
func DB() gorose.IOrm {
	return engin.NewOrm()
}

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
