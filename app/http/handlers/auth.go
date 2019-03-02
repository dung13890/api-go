package handlers

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dung13890/api-go/app/http/resources"
	"github.com/dung13890/api-go/app/services"
	"github.com/dung13890/api-go/models"
	"github.com/labstack/echo"
)

type HttpAuthHandler struct {
	AuthService services.AuthService
}

func (a *HttpAuthHandler) Login(c echo.Context) error {
	params := models.User{}
	c.Bind(&params)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	rs, err := a.AuthService.Login(ctx, params)
	if err != nil {
		return c.JSON(http.StatusNotFound, resources.Error("Login failed!"))
	}

	return c.JSON(http.StatusOK, resources.Login(rs, "Login success!"))
}

func (a *HttpAuthHandler) Info(c echo.Context) error {
	ctx := c.Request().Context()
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)

	rs, err := a.AuthService.Info(ctx, claims)
	if err != nil {
		return c.JSON(http.StatusNotFound, resources.Error("Failed!"))
	}

	return c.JSON(http.StatusOK, resources.Model(rs, "Success!"))
}

func (a *HttpAuthHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	rs := a.AuthService.Get(ctx)

	return c.JSON(http.StatusOK, resources.Collection(rs, "Success!"))
}

func (a *HttpAuthHandler) Store(c echo.Context) error {
	params := models.User{}
	c.Bind(&params)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	rs, err := a.AuthService.Store(ctx, params)
	if err != nil {
		return c.JSON(http.StatusNotFound, resources.Error("Error!"))
	}

	return c.JSON(http.StatusOK, resources.Model(rs, "Success!"))
}
