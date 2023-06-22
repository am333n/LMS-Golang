package auth

import (
	"context"
	"errors"
	"lms/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
)

// Middleware Auth Middleware
func Middleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			td := ctx.Value(TokenContextKey)
			if td == nil {
				return nil, common.ErrLoginInvalid
			}
			tokenData := td.(string)
			tokenString := tokenData
			claims := jwt.MapClaims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid || err == jwt.ErrInvalidKey {
					return nil, errors.New("invalid key")
				}
				v, _ := err.(*jwt.ValidationError)
				if v.Errors == jwt.ValidationErrorExpired {
					return nil, common.ErrTokenExpired
				}
				if v.Errors == jwt.ValidationErrorSignatureInvalid {
					return nil, errors.New("Signature invalid")
				}
				return nil, errors.New("jwt parse failed")
			}
			if !token.Valid {
				return nil, common.ErrTokenCorrupted
			}

			// c2 := context.WithValue(ctx, common.CurrentUserNameKey, claims.UserEmail)
			c2 := context.WithValue(ctx, "UserID", claims["username"])
			c2 = context.WithValue(ctx, "exp", claims["exp"])
			c2 = context.WithValue(c2, "userType", claims["userType"])
			// c2 = context.WithValue(c2, common.CurrentTenantUserIDKey, claims.TenantUserID)
			// c2 = context.WithValue(c2, common.CurrentTenantIDKey, claims.TenantID)
			// c2 = context.WithValue(c2, common.CurrentEditionIDKey, claims.EditionID)
			// c2 = context.WithValue(c2, common.CurrentUserCompanyProfileIDKey, claims.CompanyProfileID)
			// c2 = context.WithValue(c2, common.CurrentUserTimeZoneKey, claims.TimeZone)
			// c2 = context.WithValue(c2, common.CurrentUserWeekStartDayKey, claims.WeekStartDay)
			// c2 = context.WithValue(c2, common.CurrentDBMapperKey, claims.Mapper)
			return next(c2, request)
		}
	}
}
func EmployeeEntryMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			usertype:=ctx.Value("userType")
			switch usertype{
			case 1: 
				return next(ctx,nil)
			default:
				return next(ctx,common.ErrLoginInvalid)
			}
		}
	}
}
