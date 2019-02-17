package routes

import (
	"github.com/dung13890/api-go/app/providers"
	"github.com/dung13890/api-go/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	authHandler := providers.NewHttpAuthHandle()
	e.POST("/login", authHandler.Login)
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte(config.GetString("key"))))
}
