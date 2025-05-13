package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (usecase *regUseCase) GetRegCourseByCourseCode(courseCode string) ([]dto.RegCourseResponse, error) {

	regToken, err := usecase.GetRegToken()

	if err != nil {
		return nil, err
	}

	courses, err := usecase.repo.GetRegCourseByCourseCode(regToken, courseCode)

	if err != nil {
		return nil, err
	}

	return courses, nil
}
