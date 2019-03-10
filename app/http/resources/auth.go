package resources

import (
	"net/http"

	"github.com/dung13890/api-go/models"
	"github.com/labstack/echo"
)

func Login(c echo.Context, rs map[string]string) error {
	return c.JSON(http.StatusOK, resource{
		Data:     rs,
		Messages: []string{"Login Success!"},
	})
}

func Collection(c echo.Context, rs []models.User, m ...string) error {
	messages := []string{}
	if len(m) == 0 {
		messages = []string{"Success!"}
	}
	if len(m) > 0 {
		for _, i := range m {
			messages = append(messages, i)
		}
	}

	return c.JSON(http.StatusOK, resource{
		Data:     rs,
		Messages: m,
	})
}

func Model(c echo.Context, rs models.User, m ...string) error {
	messages := []string{}
	if len(m) == 0 {
		messages = []string{"Success!"}
	}
	if len(m) > 0 {
		for _, i := range m {
			messages = append(messages, i)
		}
	}
	return c.JSON(http.StatusOK, resource{
		Data:     rs,
		Messages: messages,
	})
}
