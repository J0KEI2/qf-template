package usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (usecase *regUseCase) GetRegToken() (*string, error) {

	if usecase.RegToken == nil || usecase.RegTokenExpired.After(time.Now()) {
		newToken, err := usecase.repo.GetRegToken()
		if err != nil {
			return nil, err
		}
		customClaims := entity.CustomClaims{}
		jwt.ParseWithClaims(*newToken, &customClaims, nil)
		usecase.RegToken = newToken
		usecase.RegTokenExpired = time.Unix(customClaims.ExpiresAt, 0)
	}

	if time.Now().After(usecase.RegTokenExpired) {
		newToken, err := usecase.repo.GetRegToken()
		if err != nil {
			return nil, err
		}
		customClaims := entity.CustomClaims{}
		jwt.ParseWithClaims(*newToken, &customClaims, nil)
		usecase.RegToken = newToken
		usecase.RegTokenExpired = time.Unix(customClaims.ExpiresAt, 0)
	}

	return usecase.RegToken, nil
}
