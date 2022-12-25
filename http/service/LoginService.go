package service

import (
	"encoding/json"
	"errors"
	"github.com/balqisgautama/jubelio-chat/constanta"
	"github.com/balqisgautama/jubelio-chat/dao"
	req "github.com/balqisgautama/jubelio-chat/dto/request"
	res "github.com/balqisgautama/jubelio-chat/dto/response"
	"github.com/balqisgautama/jubelio-chat/model"
	"github.com/balqisgautama/jubelio-chat/util"
	"github.com/balqisgautama/jubelio-chat/util/jwt"
	"github.com/gofiber/fiber/v2"
)

type loginService struct {
	AbstractService
}

var LoginService = loginService{}.New()

func (input loginService) New() (output loginService) {
	output.FileName = "LoginService.go"
	return
}

func (input loginService) Login(request *fiber.Request) (output res.APIResponse, header map[string]string, err error) {
	input.FuncName = "Login"

	result, output := input.readBodyAndValidateLogin(request, req.ValidateReqLogin)
	if output.Status.Code != "" {
		return
	}

	userFound, output := dao.UserDAO.GetUserByEmail(result.Email)
	if output.Status.Code != "" {
		return
	}

	if util.CheckSumWithSha256([]byte(result.Password)) != userFound.Password.String {
		output = model.GenerateValidationFailed(input.FileName, input.FuncName, errors.New(constanta.DescInvalidPassword))
		return
	}

	token, output_ := jwt.GenerateJWTUser(userFound.UserID.Int64, userFound.ClientID.String, "")
	if output_.Status.Code != "" {
		output = output_
		return
	}

	output.Status.Success = true
	output.Status.Message = constanta.RequestSuccess
	output.Data = struct {
		UserID   int64
		ClientID string
	}{
		UserID:   userFound.UserID.Int64,
		ClientID: userFound.ClientID.String,
	}

	headerData := map[string]string{
		"user-token": token,
	}
	header = headerData

	return
}

func (input loginService) readBodyAndValidateLogin(request *fiber.Request, validation func(input *req.LoginRequest) (output res.APIResponse)) (inputStruct req.LoginRequest, output res.APIResponse) {
	input.FuncName = "readBodyAndValidateLogin"
	var stringBody string

	stringBody, output = input.ReadBody(request)
	if output.Status.Code != "" {
		return
	}

	if stringBody != "" {
		errorS := json.Unmarshal([]byte(stringBody), &inputStruct)
		if errorS != nil {
			output = model.GenerateValidationFailed(input.FileName, input.FuncName, errorS)
			return
		}
	}

	output = validation(&inputStruct)
	if output.Status.Code != "" {
		return
	}

	regexPassword, err1, err2 := util.IsPasswordStandardValid(inputStruct.Password)
	if !regexPassword {
		err := errors.New(constanta.DescIncorrectFormat + " (" + err1 + " " + err2 + ")")
		output = model.GenerateValidationFailed(input.FileName, input.FuncName, err)
		return
	}
	return
}
