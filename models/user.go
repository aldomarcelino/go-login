package models

import (
	"database/sql"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	EmailHash   uint32
	Password    string
	SSOID       *string
	UserType    sql.NullString
	UserToken   sql.NullString
	FirstName   sql.NullString
	LastName    sql.NullString
	PhoneNumber sql.NullString
}