package JWTService

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gohouse/gupiao/config"
	"time"
)

var (
	jwtConfig = config.JWT
)

func GetToken(claims ...map[string]interface{}) (string, error) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// 是否传了自定义参数
	claimsReal := jwt_lib.MapClaims{}
	// 读取配置中的默认参数
	for k,v := range jwtConfig.Cliams {
		claimsReal[k] = v
	}
	// 读取用户自定义参数
	if len(claims)>0 {
		for _,item := range claims {
			for k,v := range item {
				claimsReal[k] = v
				claimsReal["iat"] = time.Now().Format("2006-01-02 15:04:05")
			}
		}
	}

	// Set some claims
	token.Claims = claimsReal
	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(jwtConfig.Secret))
}

func GetDecodeClaims(c *gin.Context) jwt_lib.Claims {
	return GetDecodeToken(c).Claims
}

func GetCliaimsItem(key string,c *gin.Context) interface{} {
	claims := GetDecodeClaims(c).(jwt_lib.MapClaims)
	return claims[key]
}

func GetDecodeToken(c *gin.Context) *jwt_lib.Token {
	token, _ := request.ParseFromRequest(c.Request, request.OAuth2Extractor,
		func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(config.JWT.Secret))
			return b, nil
		})

	return token
}
