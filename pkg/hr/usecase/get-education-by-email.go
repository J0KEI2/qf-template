package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (usecase *hrUseCase) GetEducationByEmail(email string) ([]dto.HREducationDetail, error) {

	hrToken, err := usecase.GetHRToken()

	if err != nil {
		return nil, err
	}

	response, err := usecase.repo.GetEducationByEmail(*hrToken, email)

	if err != nil {
		return nil, err
	}

	return response.Data.Items, nil
}
