package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (usecase *hrUseCase) GetLecturerBySSN(ssn string) (*dto.GetSingleEmployeeResponseDto, error) {

	hrToken, err := usecase.GetHRToken()

	if err != nil {
		return nil, err
	}

	HrEmployees, err := usecase.repo.GetLecturerBySSN(*hrToken, ssn)

	if err != nil {
		return nil, err
	}

	return HrEmployees, nil
}
