package domain

import "github.com/zercle/kku-qf-services/pkg/models/dto"

type RegUseCase interface {
	GetRegToken() (*string, error)
	GetRegCourseByCourseCode(courseCode string) ([]dto.RegCourseResponse, error)
}

type RegRepository interface {
	GetRegToken() (*string, error)
	GetRegCourseByCourseCode(regToken *string, courseCode string) ([]dto.RegCourseResponse, error)
}
