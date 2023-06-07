package token

type AccessToken struct {
	Token      string `json:"accessToken"`
	Expiration int    `json:"expiresInSeconds"`
}

func GenerateJWTToken() (string, error) {
	return "", nil
}
