package service

import (
	"github.com/balqisgautama/jubelio-chat/dto"
	"github.com/balqisgautama/jubelio-chat/dto/response"
	"github.com/balqisgautama/jubelio-chat/util"
	"github.com/gofiber/fiber/v2"
)

type AbstractService struct {
	FileName string
	FuncName string
}

func (input AbstractService) ReadBody(request *fiber.Request) (stringBody string, output response.APIResponse) {
	var errorS error
	input.FuncName = "ReadBody"
	input.FileName = "AbstractService.go"

	stringBody, _, errorS = util.ReadBody(request)
	if errorS != nil {
		output = dto.GenerateInvalidRequestBody(errorS, input.FileName, input.FuncName)
		return
	}

	return
}
