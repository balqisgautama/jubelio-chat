package dto

import (
	"github.com/balqisgautama/jubelio-chat/config"
	"github.com/balqisgautama/jubelio-chat/dto/response"
	"strings"
)

func GenerateValidationFailed(err error, filename string, funcName string) (output response.APIResponse) {
	output.Status.Success = false
	output.Status.Code = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID()) + "-370001-VALIDATION"
	output.Status.Message = err.Error()
	output.Status.Detail = []string{filename, funcName}
	return
}

func GenerateInvalidRequestBody(err error, filename string, funcName string) (output response.APIResponse) {
	output.Status.Success = false
	output.Status.Code = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID()) + "-370002-REQUEST-BODY"
	output.Status.Message = err.Error()
	output.Status.Detail = []string{filename, funcName}
	return
}

func GenerateDBServerError(err error, filename string, funcName string, tableName string) (output response.APIResponse) {
	output.Status.Success = false
	output.Status.Code = strings.ToUpper(config.ApplicationConfiguration.GetServerResourceID()) + "-370003-DB-SERVER"
	output.Status.Message = err.Error()
	output.Status.Detail = []string{filename, funcName, tableName}
	return
}
