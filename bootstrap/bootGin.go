package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gohouse/goroseGin/middleWare"
)

func BootGin() func(*Booter) {
	return func(srv *Booter) {
		gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
		srv.Router = gin.Default()

		// 配置options请求返回
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		// 配置允许 OPTIONS 请求, 默认是没有的
		config.AddAllowMethods("OPTIONS")
		config.AddAllowHeaders("Authorization")
		//config.AddAllowHeaders("token")

		// 配置静态目录
		srv.Router.Static("static", "/var/tmp/static")

		// 调用允许 OPTIONS 方法访问中间件
		srv.Router.Use(cors.New(config)).Use(middleWare.Cors())
	}
}
