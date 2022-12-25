package request

import (
	"github.com/balqisgautama/jubelio-chat/dto/response"
	"github.com/balqisgautama/jubelio-chat/model"
	"github.com/go-playground/validator/v10"
)

func init() {
	filename = "LoginRequest.go"
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=7"`
}

func ValidateReqLogin(inputStruct *LoginRequest) (output response.APIResponse) {
	funcName = "ValidateReqLogin"
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output = model.GenerateValidationFailed(filename, funcName, err)
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output = model.GenerateValidationFailed(filename, funcName, errV)
			return
		}
		return
	}
	return
}
