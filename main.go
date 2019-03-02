package main

import (
	"github.com/dung13890/api-go/config"
	"github.com/dung13890/api-go/routes"
	"github.com/labstack/echo"
)

func main() {
	config.Init()
	e := echo.New()
	s := config.NewSession()
	defer s.Close()
	routes.Init(e, s.Copy())
	e.Logger.Fatal(e.Start(config.GetString("server.address")))
}
