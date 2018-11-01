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
		"iss": "fizzday.net",
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	},
}
