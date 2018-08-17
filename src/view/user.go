package view

import (
	"controller"

	"github.com/labstack/echo"
)

// InitUserView init user view
func InitUserView(user *echo.Group) {
	user.GET("/info", controller.GetUserInfo)
}
