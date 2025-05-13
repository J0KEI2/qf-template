package usecase

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
)

func (u *authUsecase) GenerateJwtToken(user *models.UserFetchModel) (string, error) {

	claims := jwt.MapClaims{
		"uid": user.UID.String(),
		"ext": time.Now().Add(5 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte("MySignature"))
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return "", err
	}

	return ss, nil
}
