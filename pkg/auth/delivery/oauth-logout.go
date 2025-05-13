package delivery

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/utils"
)

func (h *authHandler) Logout() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		// c := &helper.c{c: c}
		sess, err := utils.SessStore.Get(c)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				helpers.ResponseError{
					Code:   http.StatusInternalServerError,
					Title:  err.Error(),
					Source: helpers.WhereAmI(),
				},
			)
		}
		defer sess.Save()

		// sess.Delete("userId")
		// sess.Delete("nameTh")
		// sess.Delete("nameEn")

		sess.Destroy()

		paramVal := url.Values{}
		scheme := string(c.Request().URI().Scheme())
		if len(scheme) == 0 {
			scheme = "http"
		}
		paramVal.Set("client_id", viper.GetString("oauth.id"))
		paramVal.Set("redirect_uri", scheme+"://"+string(c.Request().Host()))
		// data := map[string]interface{}{
		// 	"redirect_uri": scheme + "://" + string(c.Request().Host()),
		// 	"oauth_logout": viper.GetString("oauth.prd_url") + "/logout?" + paramVal.Encode(),
		// }
		// return c.Render("logout", data)

		return c.Redirect(viper.GetString("oauth.prd_url") + "/logout?" + paramVal.Encode())
	}
}
