package models

type User struct {
	ID          string `db:"id"`
	EmailHash   string `db:"email_hash"`
	Password    string `db:"password"`
	UserType    string `db:"user_type"`
	UserToken   string `db:"user_token"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	PhoneNumber string `db:"phone_number"`
}
