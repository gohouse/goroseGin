package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose/utils"
	"github.com/gohouse/goroseGin/helper"
	"github.com/gohouse/goroseGin/model"
	"github.com/gohouse/goroseGin/service/JWTService"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "123456")

	if len(username) > 0 {
		// 验证是否正确
		var user model.User
		_, err := NewConnectionInstance().Table(&user).Where("mobile", username).
			Where("password", helper.Md5(password)).First()
		if err != nil || user.Id == 0 {
			c.JSON(http.StatusOK, utils.FailReturn("帐号或密码错误", 401))
			return
		}

		token, _ := JWTService.GetToken(map[string]interface{}{
			"mobile": username,
		})
		c.JSON(http.StatusOK,
			utils.SuccessReturn(map[string]interface{}{"token": token,
				"userInfo": user}))
		return
	}

	c.JSON(http.StatusOK, utils.FailReturn("账号或密码错误!!!", 401))
}

func PasswordReset(c *gin.Context) {
}

func Register(c *gin.Context) {
}
