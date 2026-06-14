package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(secretKey string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}

			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, errors.New("invalid signing method")
				}

				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			userID := claims["user_id"]

			ctx := context.WithValue(r.Context(), "user_id", userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
