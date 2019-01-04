# goroseGin
gorose + gin demo

## 运行
1. 更改数据库配置文件 `path/to/goroseGin/config/databse.go`

2. 启动服务
```go
cd path/to/goroseGin
go run main.go
```

3. 访问测试 
```go
http://localhost:8005
http://localhost:8005/test
http://localhost:8005/api
```

## 目录说明
```sh
.
├── README.md
├── bootstrap               // 数据库框架驱动目录(gorose,gin驱动)
│   ├── bootDatabase.go
│   ├── bootGin.go
│   └── booter.go
├── config                  // 配置目录
│   ├── config.go
│   ├── database.go
│   ├── jwt.go
│   └── statusCode.go
├── controller              // 控制器controller目录
│   ├── BaseController.go
│   └── LoginController.go
├── doc
├── helper                  // 工具函数目录
│   └── helpers.go
├── main.go                 // 入口文件
├── middleWare              // 中间件
│   └── ginCors.go
├── model                   // 数据库model
│   └── model.go
├── routes                  // 路由
│   ├── admin.go
│   └── api.go
└── service                 // 公共服务service
    └── JWTService
        └── JWT.go
```