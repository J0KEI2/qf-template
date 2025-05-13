package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) UpdateCourseInfo(data *dto.CourseUpdateCourseInfoRequestDto) (res *dto.CourseInfoResponseDto, err error) {

	updateData := query.CourseInfoQueryEntity{
		CategoryName:        data.CategoryName,
		CourseCode:          data.CourseCode,
		CourseNameTH:        data.CourseNameEN,
		CourseNameEN:        data.CourseNameEN,
		TotalCredit:         data.TotalCredit,
		Credit1:             data.Credit1,
		Credit2:             data.Credit2,
		Credit3:             data.Credit3,
		CourseTypeID:        data.CourseTypeID,
		CourseConditionTH:   data.CourseConditionTH,
		CourseConditionEN:   data.CourseConditionEN,
		CourseDescriptionTH: data.CourseDescriptionTH,
		CourseDescriptionEN: data.CourseDescriptionEN,
		CourseObjective:     data.CourseObjective,
		Location:            data.Location,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateCourseInfoAction(&updateData, *data.CourseID))

	res = &dto.CourseInfoResponseDto{
		ID:                  updateData.ID,
		CategoryName:        updateData.CategoryName,
		CourseCode:          updateData.CourseCode,
		CourseNameTH:        updateData.CourseNameEN,
		CourseNameEN:        updateData.CourseNameEN,
		TotalCredit:         updateData.TotalCredit,
		Credit1:             updateData.Credit1,
		Credit2:             updateData.Credit2,
		Credit3:             updateData.Credit3,
		CourseTypeID:        updateData.CourseTypeID,
		CourseConditionTH:   updateData.CourseConditionTH,
		CourseConditionEN:   updateData.CourseConditionEN,
		CourseDescriptionTH: updateData.CourseDescriptionTH,
		CourseDescriptionEN: updateData.CourseDescriptionEN,
		CourseObjective:     updateData.CourseObjective,
		Location:            updateData.Location,
		CreatedAt:           updateData.CreatedAt,
		UpdatedAt:           updateData.UpdatedAt,
	}

	return
}

func (u *courseUsecase) updateCourseInfoAction(dataUpdate *query.CourseInfoQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		mainCourse := query.CourseMainQueryEntity{
			CourseID: &mainUid,
		}
		err = u.CommonRepository.GetFirst(&mainCourse)

		if err != nil {
			return err
		}

		if mainCourse.CourseInfoID == nil {
			return errors.New("can't find assessment for this qf")
		}

		queryUpdate := query.CourseInfoQueryEntity{
			ID: mainCourse.CourseInfoID,
		}
		err = u.CommonRepository.Update(tx, queryUpdate, dataUpdate)

		if err != nil {
			return err
		}

		return
	}
}
