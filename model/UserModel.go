package model

import "database/sql"

type UserModel struct {
	UserID    sql.NullInt64  `json:"user_id"`
	Email     sql.NullString `json:"email"`
	Password  sql.NullString `json:"password"`
	ClientID  sql.NullString `json:"client_id"`
	Status    sql.NullInt16  `json:"status"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
}
