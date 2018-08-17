package view

import (
	"controller"

	"github.com/labstack/echo"
)

// InitIndexView init index view
func InitIndexView(index *echo.Group) {
	index.GET("/health", controller.CheckHealthy)
}
