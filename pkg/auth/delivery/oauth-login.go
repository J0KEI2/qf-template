package delivery

import (
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// DESC: redirect to kku oauth
func (h *authHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		redirectURI := url.URL{}
		if len(c.Request().URI().Scheme()) == 0 {
			redirectURI.Scheme = "https"
		} else {
			redirectURI.Scheme = string(c.Request().URI().Scheme())
		}
		redirectURI.Host = string(c.Request().Host())
		redirectURI.Path = "/auth/callback"
		paramVal := url.Values{}
		paramVal.Set("response_type", "code")
		paramVal.Set("client_id", viper.GetString("oauth.id"))
		paramVal.Set("redirect_uri", base64.URLEncoding.EncodeToString([]byte(redirectURI.String())))
		return c.Redirect(viper.GetString("oauth.prd_url")+"/auth?"+paramVal.Encode(), http.StatusFound)
	}
}
