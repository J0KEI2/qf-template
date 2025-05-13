package delivery

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (h *authHandler) GenUserAccessToken(userClaims dto.UserClaims) (responseForm helpers.OauthResponse, err error) {
	// claims.Issuer = hostname
	// claims.Subject = userIds
	// claims.NameTh = user.NameTh
	// claims.NameEn = user.NameEn
	// claims.Email = strings.ToLower(user.Email)
	// claims.UserConfig = userConfig
	// claims.ExpiresAt = time.Now().Add(time.Hour * 8).Unix()
	// claims.Role = role
	// claims.HighestLevel = highestLevel

	token := jwt.NewWithClaims(h.jwtResources.JwtSigningMethod, userClaims)

	responseForm.AccessToken, err = token.SignedString(h.jwtResources.JwtSignKey)
	if err != nil {
		responseForm.Error = helpers.ServerError
		responseForm.ErrorDesc = err.Error()
		err = fiber.NewError(http.StatusInternalServerError, err.Error())
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		return
	}

	responseForm.TokenType = "Bearer"
	responseForm.ExpiresIn = int(time.Duration(time.Hour * 8).Seconds())

	return
}
