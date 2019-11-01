# 半小时快速上手golang web编程之用户的增删改查(示例代码)
博客教程页面: https://segmentfault.com/a/1190000020871866

## 说明
本实例快速实现一个用户的增删改查, 采用前后端分离模式, 后端采用`gorose orm` + `gin`框架完成api的开发, 前端使用`原生html+jquery+ajax`交互数据. 

## 安装
```shell script
git clone https://github.com/gohouse/goroseGin.git
cd goroseGin
go mod tidy
```
> 注意: 确保系统已经安装了golang运行环境和vgo,才能正常执行

## 运行
1.运行服务端
```shell script
go run main.go
```

2.访问 [http://localhost:8080/html](http://localhost:8080/html) ,就可以完成用户的增删改查操作了,效果图  
![gorose-gin.png](https://i.loli.net/2019/10/31/B7wkU8XYS6OIRdH.jpg)
