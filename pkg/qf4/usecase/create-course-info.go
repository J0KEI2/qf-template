package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) CreateOrUpdateQF4CourseInfo(data *dto.QF4CreateCourseInfoRequestDto) (res *dto.QF4CourseInfoResponseDto, err error) {
	queryTb := query.QF4MainQueryEntity{
		QF4ID: &data.QF4ID,
	}

	u.CommonRepository.GetFirst(&queryTb)

	createOrUpdateData := query.QF4CourseInfoQueryEntity{
		CategoryName:        &data.CategoryName,
		CourseCode:          &data.CourseCode,
		CourseNameTH:        &data.CourseNameEN,
		CourseNameEN:        &data.CourseNameEN,
		NumberOfCredits:     &data.NumberOfCredits,
		CourseTypeID:        &data.CourseTypeID,
		CourseConditionTH:   &data.CourseConditionTH,
		CourseConditionEN:   &data.CourseConditionEN,
		CourseDescriptionTH: &data.CourseDescriptionTH,
		CourseDescriptionEN: &data.CourseDescriptionEN,
		CourseObjective:     &data.CourseObjective,
		StudentActivity:     &data.StudentActivity,
		FacilitatorTask:     &data.FacilitatorTask,
		ConsultantTask:      &data.ConsultantTask,
		StudentGuideline:    &data.StudentGuideline,
		Location:            &data.Location,
		StudentSupport:      &data.StudentSupport,
	}

	if queryTb.QF4CourseInfoID == nil {
		err = helper.ExecuteTransaction(u.CommonRepository, u.createCourseInfoAction(&createOrUpdateData, data.QF4ID))
	} else {
		err = helper.ExecuteTransaction(u.CommonRepository, u.updateCourseInfoAction(&createOrUpdateData, data.QF4ID))
	}

	if err != nil {
		return nil, err
	}

	res = &dto.QF4CourseInfoResponseDto{
		ID:                  createOrUpdateData.ID,
		CategoryName:        createOrUpdateData.CategoryName,
		CourseCode:          createOrUpdateData.CourseCode,
		CourseNameTH:        createOrUpdateData.CourseNameEN,
		CourseNameEN:        createOrUpdateData.CourseNameEN,
		NumberOfCredits:     createOrUpdateData.NumberOfCredits,
		CourseTypeID:        createOrUpdateData.CourseTypeID,
		CourseConditionTH:   createOrUpdateData.CourseConditionTH,
		CourseConditionEN:   createOrUpdateData.CourseConditionEN,
		CourseDescriptionTH: createOrUpdateData.CourseDescriptionTH,
		CourseDescriptionEN: createOrUpdateData.CourseDescriptionEN,
		CourseObjective:     createOrUpdateData.CourseObjective,
		StudentActivity:     createOrUpdateData.StudentActivity,
		FacilitatorTask:     createOrUpdateData.FacilitatorTask,
		ConsultantTask:      createOrUpdateData.ConsultantTask,
		StudentGuideline:    createOrUpdateData.StudentGuideline,
		Location:            createOrUpdateData.Location,
		StudentSupport:      createOrUpdateData.StudentSupport,
	}

	return
}

func (u *qf4Usecase) createCourseInfoAction(data *query.QF4CourseInfoQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Create course CourseInfo

		err = u.CommonRepository.Create(tx, data)
		if err != nil {
			return err
		}

		// update course assessment id in main table
		mainQuery := query.QF4MainQueryEntity{
			QF4ID: &mainUid,
		}

		mainUpdate := query.QF4MainQueryEntity{
			QF4CourseInfoID: data.ID,
		}
		err = u.CommonRepository.Update(tx, mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
