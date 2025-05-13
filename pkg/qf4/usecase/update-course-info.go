package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) UpdateQF4CourseInfo(data *dto.QF4UpdateCourseInfoRequestDto) (res *dto.QF4CourseInfoResponseDto, err error) {

	updateData := query.QF4CourseInfoQueryEntity{
		CategoryName:        data.CategoryName,
		CourseCode:          data.CourseCode,
		CourseNameTH:        data.CourseNameEN,
		CourseNameEN:        data.CourseNameEN,
		NumberOfCredits:     data.NumberOfCredits,
		CourseTypeID:        data.CourseTypeID,
		CourseConditionTH:   data.CourseConditionTH,
		CourseConditionEN:   data.CourseConditionEN,
		CourseDescriptionTH: data.CourseDescriptionTH,
		CourseDescriptionEN: data.CourseDescriptionEN,
		CourseObjective:     data.CourseObjective,
		StudentActivity:     data.StudentActivity,
		FacilitatorTask:     data.FacilitatorTask,
		ConsultantTask:      data.ConsultantTask,
		StudentGuideline:    data.StudentGuideline,
		Location:            data.Location,
		StudentSupport:      data.StudentSupport,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateCourseInfoAction(&updateData, *data.QF4ID))

	res = &dto.QF4CourseInfoResponseDto{
		ID:                  updateData.ID,
		CategoryName:        updateData.CategoryName,
		CourseCode:          updateData.CourseCode,
		CourseNameTH:        updateData.CourseNameEN,
		CourseNameEN:        updateData.CourseNameEN,
		NumberOfCredits:     updateData.NumberOfCredits,
		CourseTypeID:        updateData.CourseTypeID,
		CourseConditionTH:   updateData.CourseConditionTH,
		CourseConditionEN:   updateData.CourseConditionEN,
		CourseDescriptionTH: updateData.CourseDescriptionTH,
		CourseDescriptionEN: updateData.CourseDescriptionEN,
		CourseObjective:     updateData.CourseObjective,
		StudentActivity:     updateData.StudentActivity,
		FacilitatorTask:     updateData.FacilitatorTask,
		ConsultantTask:      updateData.ConsultantTask,
		StudentGuideline:    updateData.StudentGuideline,
		Location:            updateData.Location,
		StudentSupport:      updateData.StudentSupport,
	}

	return
}

func (u *qf4Usecase) updateCourseInfoAction(dataUpdate *query.QF4CourseInfoQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		mainQF4 := query.QF4MainQueryEntity{
			QF4ID: &mainUid,
		}
		err = u.CommonRepository.GetFirst(&mainQF4)

		if err != nil {
			return err
		}

		if mainQF4.QF4CourseInfoID == nil {
			return errors.New("can't find assessment for this qf")
		}

		queryUpdate := query.QF4CourseInfo{
			ID: *mainQF4.QF4CourseInfoID,
		}
		err = u.CommonRepository.Update(tx, queryUpdate, dataUpdate)

		if err != nil {
			return err
		}

		return
	}
}
