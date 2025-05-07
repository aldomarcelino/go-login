package models

type LoginRequest struct {
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	Device      *string `json:"device"`
	MacAddress  *string `json:"mac_address"`
	PublicKey   *string `json:"public_key"`
	SSOID       *string `json:"sso_id"`
	SSOPlatform *string `json:"sso_platform"`
	UTM         *string `json:"utm"`
}

type LoginResponse struct {
	SessionID    string `json:"session_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}