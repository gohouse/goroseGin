package config

const (
	StatusParametersError	= 601
	StatusGoroseError		= 700
	StatusDBError			= 701
	StatusGinError			= 720
)

var statusText = map[int]string{
	StatusParametersError:	"参数有误",
	StatusGoroseError:		"gorose操作有误",
	StatusDBError:			"数据库操作有误",
	StatusGinError:			"gin操作有误",
}

func StatusText(code int) string  {
	return statusText[code]
}