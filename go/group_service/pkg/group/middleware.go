package group

import (
	// std lib
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	// third party
	"github.com/dgrijalva/jwt-go"
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
			// manage error
			return
		}
		claims, ok := tk.Claims.(*Claims)
		if !ok {
			// manage error
			return
		}
		// check expiry date
		if claims.ExpiresAt < time.Now().UTC().Unix() {
			// manage error
			return
		}

		// adding id's to the context in order to pass it in the handler
		if claims.userId == "" {
			// manage error
			return
		}
		ctx = context.WithValue(ctx, UserIdKey{}, claims.userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
