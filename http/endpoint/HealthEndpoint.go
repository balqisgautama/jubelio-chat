package endpoint

import (
	"github.com/balqisgautama/jubelio-chat/http/service"
	"github.com/gofiber/fiber/v2"
)

type healthEndpoint struct {
	AbstractEndpoint
}

var HealthEndpoint = healthEndpoint{}.New()

func (input healthEndpoint) New() (output healthEndpoint) {
	output.FileName = "HealthEndpoint.go"
	return
}

func (input healthEndpoint) CheckingHealth(c *fiber.Ctx) (err error) {
	input.ServeEndpoint(service.HealthService.CheckingHealth, c)
	return
}
