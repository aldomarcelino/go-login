package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID      `db:"id"`
	EmailHash   int         `db:"email_hash"`
	Password    *string         `db:"password"`
	UserType    *string `db:"user_type"`
	UserToken   *string `db:"user_token"`
	FirstName   *string `db:"first_name"`
	LastName    *string `db:"last_name"`
	PhoneNumber *string `db:"phone_number"`
}
