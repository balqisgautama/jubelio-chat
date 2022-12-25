package endpoint

import (
	"github.com/balqisgautama/jubelio-chat/http/service"
	"github.com/gofiber/fiber/v2"
)

type loginEndpoint struct {
	AbstractEndpoint
}

var LoginEndpoint = loginEndpoint{}.New()

func (input loginEndpoint) New() (output loginEndpoint) {
	output.FileName = "LoginEndpoint.go"
	return
}

func (input loginEndpoint) Login(c *fiber.Ctx) (err error) {
	input.ServeEndpoint(service.LoginService.Login, c)
	return
}
