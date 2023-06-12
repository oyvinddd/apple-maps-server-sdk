package token

type AccessToken struct {
	// AccessToken a JWT access token
	Token string `json:"accessToken"`

	// Expiration the expiration (in seconds) for the access token
	Expiration int `json:"expiresInSeconds"`
}
