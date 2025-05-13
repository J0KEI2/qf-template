package usecase

import (
	"fmt"
	"strings"

	"github.com/zercle/kku-qf-services/pkg/models"
)

func (usecase *userDetailUseCase) SearchLecturerByName(accessToken string, criteria *models.LecturerFetchWithNameRequestModel) (*models.LecturerFetchWithNameResponseModel, error) {

	if criteria == nil {
		return nil, fmt.Errorf("error: %s", "request body is empty")
	}

	nameParts := strings.Fields(*criteria.FullName)

	var page, size int = 1, 10
	var firstName, lastName string

	if len(nameParts) > 0 {
		firstName = nameParts[0]
	}
	if len(nameParts) > 1 {
		lastName = strings.Join(nameParts[1:], " ")
	}

	splitNameRequestModel := models.LecturerFetchWithNameQueryModel{
		FirstName: &firstName,
		LastName:  &lastName,
		Page:      &page,
		Size:      &size,
	}

	output, err := usecase.userDetailRepo.GetAllEmployees(accessToken, &splitNameRequestModel)

	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return output, err
}
