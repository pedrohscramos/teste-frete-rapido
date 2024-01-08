package middlewares

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/spf13/viper"
)

func Csrf() func(handler http.Handler) http.Handler {
	secure := viper.GetString("api.environment") == "production"

	return csrf.Protect(
		[]byte(viper.GetString("api.secret")),
		csrf.RequestHeader("X-CSRF-Token"),
		csrf.CookieName("xsrf-token"),
		csrf.Secure(secure),
		csrf.Path("/"),
	)
}
