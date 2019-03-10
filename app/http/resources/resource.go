package resources

import (
	"net/http"

	"github.com/labstack/echo"
)

type resource struct {
	Data     interface{} `json:"data"`
	Messages []string    `json:"messages"`
}

func Validate(c echo.Context, m []string) error {
	return c.JSON(http.StatusUnprocessableEntity, resource{
		Data:     nil,
		Messages: m,
	})
}

func Error(c echo.Context, m string, code ...int) error {
	if m == "" {
		m = "Response Failed!"
	}
	if len(code) == 0 {
		code[0] = http.StatusNotFound
	}

	return c.JSON(code[0], resource{
		Data:     nil,
		Messages: []string{m},
	})
}
