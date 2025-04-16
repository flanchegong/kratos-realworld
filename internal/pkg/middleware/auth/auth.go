package auth

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
)

const (
	bearerWord string = "Token"
)

func GenerateToken(username string) string {
	// Create the token using your secret key and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// JWTAuth is a middleware that checks the JWT token in the request header.
func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				auths := strings.SplitN(tokenString, " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
					return nil, errors.New("missing jwt token")
				}
				spew.Dump(tokenString)
				token, err := jwt.Parse(auths[1], func(token *jwt.Token) (interface{}, error) {
					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return []byte(secret), nil
				}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
				if err != nil {
					log.Fatal(err)
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					spew.Dump(claims["username"])
				} else {
					return nil, errors.New("invalid token")
				}
			}
			return handler(ctx, req)
		}
	}
}
