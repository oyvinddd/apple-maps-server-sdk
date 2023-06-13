package token

type AccessToken struct {
	// AccessToken a JWT access token
	Token string `json:"accessToken"`

	// Expiration the expiration (in seconds) for the access token
	Expiration int `json:"expiresInSeconds"`
}

func (at AccessToken) Empty() bool {
	return len(at.Token) == 0 && at.Expiration == 0
}
