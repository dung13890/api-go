package routes

import (
	"net/http"

	"github.com/dung13890/api-go/app/http/resources"
	"github.com/dung13890/api-go/app/providers"
	"github.com/dung13890/api-go/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mgo "gopkg.in/mgo.v2"
)

func Init(e *echo.Echo, session *mgo.Session) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.HTTPErrorHandler = customHTTPErrorHandler
	authHandler := providers.NewHttpAuthHandle(session)
	e.POST("api/login", authHandler.Login)
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte(config.GetString("key"))))
	r.GET("/users", authHandler.Get)
	r.GET("/info", authHandler.Info)
	r.POST("/user", authHandler.Store)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Server Error!"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}

	resources.Error(c, message, code)
}
