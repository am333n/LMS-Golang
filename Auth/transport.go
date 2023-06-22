package auth

import (
	"context"
	"net/http"
	khttp "github.com/go-kit/kit/transport/http"
)

var TokenContextKey = "ghfdyrdyytrd5"

func extractCookieToken(ctx context.Context, r *http.Request) (context.Context, error) {
	c, err := r.Cookie("logintoken")
	if err != nil {
		if err == http.ErrNoCookie {
			return ctx, nil
		}
		return ctx, err
	}
	authToken := c.Value
	c2 := context.WithValue(ctx, TokenContextKey, authToken)
	return c2, nil
}

//	func extractHeaderToken(ctx context.Context, r *http.Request) (context.Context, error) {
//		tokenHeader := r.Header.Get("Authorization")
//		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
//			return ctx, errors.New("invalid")
//		}
//		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
//		if len(splitted) != 2 {
//			return ctx, errors.New("invalid")
//		}
//		authToken := &Token{
//			Name:  "token",
//			Value: splitted[1],
//		}
//		c2 := context.WithValue(ctx, TokenContextKey, authToken)
//		return c2, nil
//	}
var ExtractTokenData khttp.RequestFunc = func(ctx context.Context, r *http.Request) context.Context {
	context, err := extractCookieToken(ctx, r)
	if err != nil {
		return context
	}
	return context
}
