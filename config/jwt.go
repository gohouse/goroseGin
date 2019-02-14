package config

import (
	"time"
)

type jwt struct {
	Secret string
	Cliams map[string]interface{}
}

var JWT = jwt{
	Secret: "secret-strings",
	Cliams: map[string]interface{}{
		"iss": "fizzday.net",	// 发行方
		"exp": time.Now().Add(time.Hour * 4).Unix(),	// 过期时间 4 消失
	},
}
