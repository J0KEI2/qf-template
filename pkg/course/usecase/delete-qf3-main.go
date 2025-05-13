package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (useCase courseUsecase) DeleteCourseMain(data *dto.CourseDeleteMainRequestDto) (err error) {
	convertUid, _ := uuid.Parse(data.CourseUID)

	helper.ExecuteTransaction(useCase.CommonRepository, func(tx *gorm.DB) error {
		deleteQuery := query.CourseMainQueryEntity{
			CourseID: &convertUid,
		}
		return useCase.CommonRepository.DeleteMainQFWithWhereClause(tx, &deleteQuery, "course_id", convertUid)
	})

	return nil
}
