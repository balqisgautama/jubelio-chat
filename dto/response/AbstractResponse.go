package response

import (
	"github.com/balqisgautama/jubelio-chat/util"
)

type APIResponse struct {
	Timestamp string         `json:"timestamp"`
	Status    StatusResponse `json:"status"`
	Data      interface{}
}

type StatusResponse struct {
	Success bool     `json:"success"`
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Detail  []string `json:"detail"`
}

func (input APIResponse) String() string {
	return util.StructToJSON(input)
}
