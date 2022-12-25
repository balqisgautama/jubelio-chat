package service

import (
	"github.com/balqisgautama/jubelio-chat/config/server"
	"github.com/balqisgautama/jubelio-chat/dto/response"
	"github.com/gofiber/fiber/v2"
)

type healthService struct {
	AbstractService
}

var HealthService = healthService{}.New()

func (input healthService) New() (output healthService) {
	output.FileName = "HealthService.go"
	return
}

func (input healthService) CheckingHealth(request *fiber.Request) (output response.APIResponse, header map[string]string, err error) {
	output.Status.Success = true
	output.Status.Message = "UP"
	if server.ServerConfig.DBConnection.Ping() != nil {
		output.Status.Message = "DOWN"
		output.Status.Detail = []string{"PostgreSQL"}
	}
	return
}
