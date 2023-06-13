package token

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	accessTokenURL string = "https://maps-api.apple.com/v1/token"

	invalidOrMissingAuthToken tokenState = 0

	invalidOrMissingAccessToken tokenState = 1
)

var (
	fooErr = errors.New("hello")

	missingAccessToken = errors.New("missing access token")

	unableToCreateRequest = errors.New("unable to create http request")

	unableToGetAccessToken = errors.New("unable to get access token")

	responseWasNotOK = errors.New("http response was not 200 ok")

	unableToDecodeAccessToken = errors.New("unable to decode access token")
)

type (
	tokenState uint8

	Manager struct {
		authorizationToken string
		apiAccessToken     AccessToken
	}
)

func NewManager(authorizationToken string) *Manager {
	return &Manager{authorizationToken: authorizationToken}
}

func (m *Manager) GetAccessToken(c *http.Client) (AccessToken, error) {
	if m.apiAccessToken.Empty() {
		return m.apiAccessToken, nil
	}
	req, err := http.NewRequest(http.MethodGet, accessTokenURL, nil)
	if err != nil {
		return AccessToken{}, unableToCreateRequest
	}

	req.Header.Set("Authorization", m.authorizationToken)

	res, err := c.Do(req)
	if err != nil {
		return AccessToken{}, unableToGetAccessToken
	}

	if res.StatusCode != 200 {
		return AccessToken{}, responseWasNotOK
	}

	var accessToken AccessToken
	if err := json.NewDecoder(res.Body).Decode(&accessToken); err != nil {
		return AccessToken{}, unableToDecodeAccessToken
	}

	return accessToken, nil
}
