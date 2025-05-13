package usecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type courseUsecase struct {
	courseRepository domain.CourseRepository
	CommonRepository domain.CommonRepository
}

func NewCourseUsecase(repo domain.CourseRepository, commonRepo domain.CommonRepository) domain.CourseUsecase {
	return &courseUsecase{
		courseRepository: repo,
		CommonRepository: commonRepo,
	}
}
