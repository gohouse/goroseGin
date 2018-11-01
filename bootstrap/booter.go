package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"
	"sync"
)

type Booter struct {
	Router *gin.Engine
	Connection *gorose.Connection
}

var b *Booter
var once sync.Once
// GetBooterInstance 驱动单例
func GetBooterInstance() *Booter {
	once.Do(func() {
		b = &Booter{}
	})
	return b
}

// Use : 驱动中间件
func (b *Booter) Use(options ...func(*Booter))  {
	for _, option := range options {
		option(b)
	}
}
