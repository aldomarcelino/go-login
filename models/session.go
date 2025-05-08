package models

// type Session struct {
// 	ID           uuid.UUID
// 	UserID       uuid.UUID
// 	AccessToken  string
// 	RefreshToken string
// 	Device       *string
// 	MacAddress   *string
// 	Active       bool
// }

type Session struct {
	ID              string `db:"id"`
	UserID          string `db:"user_id"`
	ClientVersion   string `db:"client_version"`
	Device          string `db:"device"`
	MacAddress      string `db:"mac_address"`
	PublicKey       string `db:"public_key"`
	ChallengeString string `db:"challenge_string,omitempty"`
	Active          bool   `db:"active"`
	IP              string `db:"ip,omitempty"`
	UserAgent       string `db:"user_agent,omitempty"`
	Location        string `db:"location,omitempty"`
}
