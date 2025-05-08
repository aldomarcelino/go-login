package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID      `db:"id"`
	EmailHash   uint32         `db:"email_hash"`
	Password    string         `db:"password"`
	UserType    sql.NullString `db:"user_type"`
	UserToken   sql.NullString `db:"user_token"`
	FirstName   sql.NullString `db:"first_name"`
	LastName    sql.NullString `db:"last_name"`
	PhoneNumber sql.NullString `db:"phone_number"`
}
