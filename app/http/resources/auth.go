package resources

import (
	"github.com/dung13890/api-go/models"
)

type resource struct {
	Data     interface{} `json:"data,omitempty"`
	Messages []string    `json:"messages,omitempty"`
}

func Login(rs models.User, m string) resource {
	return resource{
		Data:     rs,
		Messages: []string{m},
	}
}
