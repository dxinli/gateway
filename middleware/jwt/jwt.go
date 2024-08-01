package jwt

import (
	"errors"
	config "github.com/go-kratos/gateway/api/gateway/config/v1"
	"github.com/go-kratos/gateway/middleware"
	"github.com/go-kratos/gateway/proxy/auth"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var NotAuthN = errors.New("unauthorized: authentication required")

func init() {
	middleware.Register("jwt", Middleware)
}

func ParseJwt(tokenString, secretKey string) (*auth.User, error) {
	t, err := jwt.ParseWithClaims(tokenString, &auth.User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*auth.User); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func Middleware(c *config.Middleware) (middleware.Middleware, error) {
	return func(next http.RoundTripper) http.RoundTripper {
		return middleware.RoundTripperFunc(func(req *http.Request) (resp *http.Response, err error) {
			authHeader := req.Header.Get("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				// 使用SplitN分割字符串，只分割一次
				parts := strings.SplitN(authHeader, " ", 2)
				if len(parts) == 2 {
					// 提取Bearer后的令牌
					token := parts[1]
					// 认证 token
					user, err := ParseJwt(token, "test_jwt")
					if err != nil {
						log.Errorf("jwt parse error: %v", err)
						return nil, errors.Join(NotAuthN, err)
					}
					reqOpt, _ := middleware.FromRequestContext(req.Context())
					reqOpt.Values.Set("user", user)
					return next.RoundTrip(req)
				} else {
					// 如果没有找到有效的Bearer令牌，则返回错误
					return nil, errors.Join(NotAuthN, errors.New("authorization not have credentials"))
				}
			} else {
				// 如果没有找到有效的Bearer令牌，则返回错误
				return nil, errors.Join(NotAuthN, errors.New("authorization not have bearer scheme"))
			}
		})
	}, nil
}
