package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oyvinddd/apple-maps-server-sdk/location"
	"github.com/oyvinddd/apple-maps-server-sdk/token"
	"log"
	"net/http"
	"net/url"
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
	GenerateAccessToken() (token.AccessToken, error)

	// Geocode geocodes the specified address and returns the location (lat/long)
	Geocode(address string, countries []string, lang string, searchLocation location.Location) error

	// ReverseGeocode returns the address located at a specific location (lat/long)
	ReverseGeocode(loc location.Location, lang string) ([]location.Place, error)
}

type appleMapsSDK struct {
	authorizationToken string

	accessToken string

	client http.Client
}

func (sdk appleMapsSDK) GenerateAccessToken() (token.AccessToken, error) {
	req, err := buildAuthenticatedRequest(tokenPath, sdk.authorizationToken)
	if err != nil {
		return token.AccessToken{}, err
	}

	res, err := sdk.client.Do(req)
	if err != nil {
		return token.AccessToken{}, err
	}

	var accessToken token.AccessToken
	if err := json.NewDecoder(res.Body).Decode(&accessToken); err != nil {
		return token.AccessToken{}, err
	}
	return accessToken, nil
}

func (sdk appleMapsSDK) Geocode(query string, limitToCountries []string, lang string, searchLocation location.Location) error {
	_, err := buildAuthenticatedRequest(http.MethodGet, sdk.accessToken)
	if err != nil {
		return err
	}
	return nil
}

func (sdk appleMapsSDK) ReverseGeocode(loc location.Location, lang string) ([]location.Place, error) {
	req, err := buildAuthenticatedRequest(reverseGeocodePath, sdk.accessToken)

	req.URL.RawQuery = url.Values{
		"loc":  {loc.String()},
		"lang": {lang}, // default = en-US
	}.Encode()

	res, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	var places []location.Place
	if err := json.NewDecoder(res.Body).Decode(&places); err != nil {
		fmt.Print(res.StatusCode)
		return nil, err
	}

	return places, nil
}

func NewWithToken(token string) AppleMapsSDK {
	return &appleMapsSDK{token, "", http.Client{}}
}

func buildAuthenticatedRequest(path, token string) (*http.Request, error) {
	if token == "" {
		return nil, errors.New("unauthorized - no access token present")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiURL, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return req, nil
}

func main() {

	jwtToken := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlpHSk5IWTQ4N1MifQ.eyJpc3MiOiIyU01GTE02NlI5IiwiaWF0IjoxNjg2MTM5MzkzLCJleHAiOjE2ODg3MzExODh9.NhMY58eABMdHw366XCX5dlH2nWFUjqJ20Pye7UTk3gy9ADH3eFhGBvJAIue3SCdKkPOPfqBYjitFIM4V67ES0g"
	appleMapsSDK := NewWithToken(jwtToken)

	places, err := appleMapsSDK.ReverseGeocode(location.New(60.0, 5.0), "en-US")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(places)
}
