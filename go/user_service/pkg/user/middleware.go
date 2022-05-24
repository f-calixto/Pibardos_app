package user

import (
	// std lib
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	// third party
	"github.com/golang-jwt/jwt"
)

var secretKey = os.Getenv("JWT_SECRET")

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context = r.Context()

		// validate if header is well formed
		if r.Header["Authorization"] == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		authorization := strings.Split(r.Header["Authorization"][0], " ")
		if authorization[0] != "Bearer" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// parse token with secret key and extract encoded information
		token := authorization[1]
		tk, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			fmt.Println("3")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		claims, ok := tk.Claims.(*Claims)
		if !ok {
			fmt.Println("4")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// check expiry date
		if claims.ExpiresAt < time.Now().UTC().Unix() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// adding id's to the context in order to pass it in the handler
		if claims.UserId == "" {
			fmt.Println("5")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx = context.WithValue(ctx, UserIdKey{}, claims.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
