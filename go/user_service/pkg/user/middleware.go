package user

import (
	// std lib
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	// internal
	"github.com/coding-kiko/user_service/pkg/errors"

	// third party
	"github.com/golang-jwt/jwt"
)

var secretKey = os.Getenv("JWT_SECRET")

// used to parse jwt payload
type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var ctx context.Context = r.Context()

		// validate if header is well formed
		if r.Header["Authorization"] == nil {
			statusCode, resp := errors.CreateResponse(errors.NewJwtBadRequest("malformed header"))
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(resp)
			return
		}
		authorization := strings.Split(r.Header["Authorization"][0], " ")
		if authorization[0] != "Bearer" {
			statusCode, resp := errors.CreateResponse(errors.NewJwtBadRequest("malformed header"))
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(resp)
			return
		}

		// parse token with secret key and extract encoded information
		token := authorization[1]
		tk, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("error parsing jwt"))
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(resp)
			return
		}
		claims, ok := tk.Claims.(*Claims)
		if !ok {
			statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("error parsing jwt"))
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(resp)
			return
		}
		// check expiry date
		if claims.ExpiresAt < time.Now().UTC().Unix() {
			statusCode, resp := errors.CreateResponse(errors.NewJwtAuthorization("token expired"))
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(resp)
			return
		}

		// adding id's to the context in order to pass it in the handler
		if claims.UserId == "" {
			statusCode, resp := errors.CreateResponse(errors.NewJwtBadRequest("missing user id in jwt"))
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(resp)
			return
		}
		ctx = context.WithValue(ctx, UserIdKey{}, claims.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
