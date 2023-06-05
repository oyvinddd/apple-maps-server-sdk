package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/oyvinddd/apple-maps-server-sdk/location"
	"net/http"
	"time"
)

const (
	apiURL             string = "https://maps-api.apple.com/v1"
	tokenPath          string = "/token"
	geocodePath        string = "/geocode"
	reverseGeocodePath string = "/reverseGeocode"
	searchPath         string = "/search"
)

type AppleMapsSDK interface {
	// GenerateAccessToken generates a JWT token for accessing Apple APIs
	GenerateAccessToken(keyID, teamID string) (string, error)

	// Geocode geocodes the specified address and returns the location (lat/long)
	Geocode(address string) error

	// ReverseGeocode returns the address located at a specific location (lat/long)
	ReverseGeocode(loc location.Location, lang string) error

	// Search searches for POIs ....
	Search(query string, lang string)
}

type appleMapsSDK struct{}

func (sdk appleMapsSDK) GenerateAccessToken(keyID, teamID string) (string, error) {
	var secret []byte // TODO: ...
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": teamID,
		"kid": "SOMETHING_HERE",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	t, err := token.SignedString(secret)
	return t, err
}

func (sdk appleMapsSDK) Geocode(address string) error {
	_, err := http.NewRequest(http.MethodGet, urlWithPath(geocodePath), nil)
	if err != nil {
		return err
	}
	return nil
}

// New creates a new instance of the Apple Maps Server SDK
func New(keyID, teamID string) *AppleMapsSDK {
	return nil
}

func NewWithToken(token string) *AppleMapsSDK {
	return &appleMapsSDK{}
}

func urlWithPath(path string) string {
	return fmt.Sprintf("%s%s", apiURL, path)
}

func main() {

	jwtToken := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlBDQTY1MjkzTk0ifQ.eyJpc3MiOiIyU01GTE02NlI5IiwiaWF0IjoxNjg1OTcyOTkyLCJleHAiOjE2OTU5NDU2MDB9.zP4GVNw5lWRUiHa1irk1R3yItlYjUC_kBQG4jszU3JKUjR_CxVuZ6Iq9ySD-N4NPFhew1i2MIe9nZDGjVypgfw"

	sdk := NewWithToken(jwtToken)
}
