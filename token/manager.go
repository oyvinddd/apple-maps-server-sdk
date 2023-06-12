package token

import (
	"net/http"
)

const accessTokenURL string = "https://maps-api.apple.com/v1/token"

type Manager struct {
	authorizationToken string
	apiAccessToken     AccessToken
}

func NewManager(authorizationToken string) *Manager {
	return &Manager{authorizationToken: authorizationToken}
}

func (m *Manager) GetAccessToken() (AccessToken, error) {
	_, err := http.NewRequest(http.MethodGet, accessTokenURL, nil)
	if err != nil {
		return AccessToken{}, err
	}
	return AccessToken{}, nil
}
