package model

import "github.com/balqisgautama/jubelio-chat/dto/response"

func GenerateDBServerError(fileName string, funcName string, tableName string, causedBy error) (output response.APIResponse) {
	output.Status.Success = false
	output.Status.Code = "JUBELIO-370001-DB-SERVER"
	output.Status.Message = causedBy.Error()
	output.Status.Detail = []string{fileName, funcName, tableName}
	return
}

func GenerateValidationFailed(fileName string, funcName string, causedBy error) (output response.APIResponse) {
	output.Status.Success = false
	output.Status.Code = "JUBELIO-370002-VALIDATION"
	output.Status.Message = causedBy.Error()
	output.Status.Detail = []string{fileName, funcName}
	return
}

func GenerateJWTError(fileName string, funcName string, causedBy error) (output response.APIResponse) {
	output.Status.Success = false
	output.Status.Code = "JUBELIO-370002-JWT"
	output.Status.Message = causedBy.Error()
	output.Status.Detail = []string{fileName, funcName}
	return
}
