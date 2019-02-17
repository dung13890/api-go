package handlers

import (
	"net/http"

	"github.com/dung13890/api-go/app/contracts"
	"github.com/dung13890/api-go/models"
	"github.com/labstack/echo"
)

type HttpAuthHandler struct {
	UserRepo contracts.UserRepo
}

type resource struct {
	Data     models.User `json:"data,omitempty"`
	Messages []string    `json:"messages,omitempty"`
}

func (a *HttpAuthHandler) Login(c echo.Context) error {
	params := models.User{}
	c.Bind(&params)
	a.UserRepo.Login(params)
	rs := resource{
		Data: params,
		Messages: []string{
			"message ok!",
		},
	}

	return c.JSON(http.StatusOK, rs)
}
