package handlers

import (
	"context"
	"net/http"

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

	rs := a.AuthService.Login(ctx, params)

	return c.JSON(http.StatusOK, resources.Login(rs, "Login success!"))
}
