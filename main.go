package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/oyvinddd/apple-maps-server-sdk/location"
	"github.com/oyvinddd/apple-maps-server-sdk/place"
	"net/http"
	"net/url"
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
	Geocode(address string, countries []string, lang string, searchLocation location.Location) error

	// ReverseGeocode returns the address located at a specific location (lat/long)
	ReverseGeocode(loc location.Location, lang string) ([]place.Place, error)

	// Search searches for POIs ....
	//Search(query string, lang string)
}

type appleMapsSDK struct {
	accessToken string

	client http.Client
}

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

func (sdk appleMapsSDK) Geocode(address string, limitToCountries []string, lang string, searchLocation location.Location) error {
	_, err := buildAuthenticatedRequest(http.MethodGet, sdk.accessToken)
	if err != nil {
		return err
	}
	return nil
}

func (sdk appleMapsSDK) ReverseGeocode(loc location.Location, lang string) ([]place.Place, error) {
	req, err := buildAuthenticatedRequest(reverseGeocodePath, sdk.accessToken)

	req.URL.RawQuery = url.Values{
		"loc":  {loc.String()},
		"lang": {"en-US"}, // default value
	}.Encode()

	fmt.Printf("FULL URL: %s", req.URL.RawQuery)

	res, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	var places []place.Place
	if err := json.NewDecoder(res.Body).Decode(&places); err != nil {
		fmt.Print(res.StatusCode)
		return nil, err
	}

	return places, nil
}

func NewWithToken(token string) AppleMapsSDK {
	return &appleMapsSDK{token, http.Client{}}
}

func buildAuthenticatedRequest(path, accessToken string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiURL, path), nil)
	if err != nil {
		return nil, err
	}
	// FIXME: we're using the wrong access token here. Access token from API endpoint should be used instead.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	return req, nil
}

func main() {

	jwtToken := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlBDQTY1MjkzTk0ifQ.eyJpc3MiOiIyU01GTE02NlI5IiwiaWF0IjoxNjg1OTcyOTkyLCJleHAiOjE2OTU5NDU2MDB9.zP4GVNw5lWRUiHa1irk1R3yItlYjUC_kBQG4jszU3JKUjR_CxVuZ6Iq9ySD-N4NPFhew1i2MIe9nZDGjVypgfw"

	loc := location.New(60.382778, 5.316600)
	_, err := NewWithToken(jwtToken).ReverseGeocode(loc, "NO")
	if err != nil {
		panic(err)
	}
}
