package handler

import (
	"net/http"
	"strings"
)

func ExtractJwtTokenFromHeader(r *http.Request) (*string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, ErrNoAuthHeader
	}
	splitToken := strings.Split(authHeader, "Bearer")
	if len(splitToken) != 2 {
		return nil, ErrNoAuthHeader
	}
	token := strings.TrimSpace(splitToken[1])
	return &token, nil
}
