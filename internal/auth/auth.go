package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
GetAPIKey extracts an API Key from
the headers of an HTTP request.
Example:
Authorization: ApiKey {API_KEY}
*/
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Authentication header")

	}
	if vals[0] != "ApiKey" {

		return "", errors.New("malformed first part of Authentication header. should start with ApiKey")
	}
	return vals[1], nil
}
