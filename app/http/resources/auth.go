package resources

import (
	"github.com/dung13890/api-go/models"
)

type resource struct {
	Data     interface{} `json:"data"`
	Messages []string    `json:"messages"`
}

func Error(m string) resource {
	return resource{
		Data:     nil,
		Messages: []string{m},
	}
}

func Login(rs map[string]string, m string) resource {
	return resource{
		Data:     rs,
		Messages: []string{m},
	}
}

func Collection(rs []models.User, m string) resource {
	return resource{
		Data:     rs,
		Messages: []string{m},
	}
}

func Model(rs models.User, m string) resource {
	return resource{
		Data:     rs,
		Messages: []string{m},
	}
}
