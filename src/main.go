/*
Package main package is the entry file
*/
package main

import (
	"config"

	mid "middleware"
	"view"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	validator "gopkg.in/go-playground/validator.v9"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// 参数验证器
	e.Validator = &mid.DefaultValidator{Validator: validator.New()}

	v1 := e.Group("/api/v1")
	view.InitIndexView(v1)

	weixin := v1.Group("/weixin")
	view.InitWeixinView(weixin)

	user := v1.Group("/user")
	view.InitUserView(user)

	// 启动
	e.Logger.Fatal(e.Start(config.Conf.AppInfo.Addr))
}
