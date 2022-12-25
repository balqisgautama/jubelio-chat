package dao

import (
	"github.com/balqisgautama/jubelio-chat/config/server"
	"github.com/balqisgautama/jubelio-chat/dto/response"
	"github.com/balqisgautama/jubelio-chat/model"
)

type userDAO struct {
	AbstractDAO
}

var UserDAO = userDAO{}.New()

func (input userDAO) New() (output userDAO) {
	output.FileName = "UserDAO.go"
	output.TableName = "users"
	return
}

// GetUserByEmail ------------------------------------------------------------------------------------------
func (input userDAO) GetUserByEmail(email string) (result model.UserModel, output response.APIResponse) {
	input.FuncName = "GetUserByEmail"

	sqlStatement := `SELECT * FROM ` + input.TableName + ` WHERE email=$1 AND ` +
		`EXTRACT(EPOCH FROM deleted_at) is NULL`

	row := server.ServerConfig.DBConnection.QueryRow(sqlStatement, email)

	err := row.Scan(&result.UserID, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)

	if err != nil {
		output = model.GenerateDBServerError(input.FileName, input.FuncName, input.TableName, err)
		return
	}

	return
}
