package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		jwtSecret := os.Getenv("JWT_SECRET")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Write([]byte(`{"error": "invalid token"}`))
			return
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(jwtSecret), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Write([]byte(`{"error": "invalid token"}`))
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "JWT", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value(props).(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Write([]byte(`{"error": "invalid token"}`))
				return
			}
		}
	})
}
