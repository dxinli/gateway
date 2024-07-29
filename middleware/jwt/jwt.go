package jwt

import (
	config "github.com/go-kratos/gateway/api/gateway/config/v1"
	"github.com/go-kratos/gateway/middleware"
	"net/http"
)

func init() {
	middleware.Register("jwt", Middleware)
}

func Middleware(c *config.Middleware) (middleware.Middleware, error) {
	return func(next http.RoundTripper) http.RoundTripper {
		return middleware.RoundTripperFunc(func(req *http.Request) (reply *http.Response, err error) {
			return
		})
	}, nil
}
