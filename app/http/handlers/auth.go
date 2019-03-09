package handlers

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dung13890/api-go/app/http/resources"
	"github.com/dung13890/api-go/app/http/validators"
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
	errV := validators.Login(&params)
	if len(errV) > 0 {
		return resources.Validate(c, errV)
	}
	rs, err := a.AuthService.Login(ctx, params)
	if err != nil {
		return resources.Error(c, "Login Faild")
	}

	return resources.Login(c, rs)
}

func (a *HttpAuthHandler) Info(c echo.Context) error {
	ctx := c.Request().Context()
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)

	rs, err := a.AuthService.Info(ctx, claims)
	if err != nil {
		return resources.Error(c, err.Error())
	}

	return resources.Model(c, rs)
}

func (a *HttpAuthHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	rs, err := a.AuthService.Get(ctx)
	if err != nil {
		return resources.Error(c, err.Error())
	}

	return resources.Collection(c, rs, "success")
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
		return resources.Error(c, err.Error())
	}

	return resources.Model(c, rs)
}
