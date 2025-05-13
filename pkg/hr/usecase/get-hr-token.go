package usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (usecase *hrUseCase) GetHRToken() (*string, error) {

	if usecase.HRToken == nil || usecase.HRTokenExpired.After(time.Now()){
		newToken, err := usecase.repo.GetHrToken()
		if err != nil {
			return nil, err
		}
		customClaims := entity.CustomClaims{}
		jwt.ParseWithClaims(*newToken, &customClaims, nil)
		usecase.HRToken = newToken
		usecase.HRTokenExpired = time.Unix(customClaims.ExpiresAt, 0)
	}

	return usecase.HRToken, nil
}
