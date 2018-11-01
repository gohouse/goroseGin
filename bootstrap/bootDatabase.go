package bootstrap

import (
	"github.com/gohouse/gorose"
	"github.com/gohouse/goroseGin/config"
	_ "github.com/go-sql-driver/mysql"
)

func BootDatabase() func(*Booter) {
	return func(srv *Booter) {
		// 加载database
		connection,err := gorose.Open(config.DbConfig)

		if err!=nil{
			panic(err)
		}

		srv.Connection = connection
	}
}