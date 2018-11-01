package main

import (
	"github.com/gohouse/gupiao/bootstrap"
	"github.com/gohouse/gupiao/routes"
)

func main() {
	// 驱动基本服务
	var booter = bootstrap.GetBooterInstance()
	// 驱动数据库
	booter.Use(bootstrap.BootDatabase())
	// 驱动gin框架
	booter.Use(bootstrap.BootGin())

	// 延迟关闭数据库
	defer booter.Connection.Close()

	// 加载路由
	routes.Run()

	//监听端口
	booter.Router.Run(":8003")
}
