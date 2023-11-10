package router

import (
	"github.com/Math-dev264/AutoTrade/Backend/src/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initialeRoutes(router *echo.Echo) {
	basePath := "/api"
	e := router.Group(basePath + "/auth")
	{
		e.Use(middleware.Logger())
		e.POST("/signup", handlers.Signup)
		e.POST("/login", handlers.Login)
	}
}
