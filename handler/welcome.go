package handler

import (
	"errors"
	"fmt"
	"jwt_auth/security"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var ErrNoAuthHeader = errors.New("jwt: Authorization Header not present in the request")

func Welcome(w http.ResponseWriter, r *http.Request) {

	authHeader, err := ExtractJwtTokenFromHeader(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := &security.Claims{}

	tkn, err := jwt.ParseWithClaims(*authHeader, claims, func(token *jwt.Token) (interface{}, error) {
		return security.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
