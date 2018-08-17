package view

import (
	"controller"

	"github.com/labstack/echo"
)

// InitWeixinView init weixin view
func InitWeixinView(weixin *echo.Group) {
	weixin.GET("/redirect_uri", controller.GetRedirectURI)
	weixin.GET("/code", controller.SetUserInfoByCode)
}
