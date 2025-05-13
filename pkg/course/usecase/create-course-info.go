package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) CreateOrUpdateCourseInfo(data *dto.CourseCreateCourseInfoRequestDto) (res *dto.CourseInfoResponseDto, err error) {
	if err != nil {
		return nil, err
	}

	queryTb := query.CourseMainQueryEntity{
		CourseID: &data.CourseID,
	}

	u.CommonRepository.GetFirst(&queryTb)

	createOrUpdateData := query.CourseInfoQueryEntity{
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

	if queryTb.CourseInfoID == nil {
		err = helper.ExecuteTransaction(u.CommonRepository, u.createCourseInfoAction(&createOrUpdateData, data.CourseID))
	} else {
		err = helper.ExecuteTransaction(u.CommonRepository, u.updateCourseInfoAction(&createOrUpdateData, data.CourseID))
	}

	if err != nil {
		return nil, err
	}

	res = &dto.CourseInfoResponseDto{
		ID:                  createOrUpdateData.ID,
		CategoryName:        createOrUpdateData.CategoryName,
		CourseCode:          createOrUpdateData.CourseCode,
		CourseNameTH:        createOrUpdateData.CourseNameEN,
		CourseNameEN:        createOrUpdateData.CourseNameEN,
		TotalCredit:         createOrUpdateData.TotalCredit,
		Credit1:             createOrUpdateData.Credit1,
		Credit2:             createOrUpdateData.Credit2,
		Credit3:             createOrUpdateData.Credit3,
		CourseTypeID:        createOrUpdateData.CourseTypeID,
		CourseConditionTH:   createOrUpdateData.CourseConditionTH,
		CourseConditionEN:   createOrUpdateData.CourseConditionEN,
		CourseDescriptionTH: createOrUpdateData.CourseDescriptionTH,
		CourseDescriptionEN: createOrUpdateData.CourseDescriptionEN,
		CourseObjective:     createOrUpdateData.CourseObjective,
		Location:            createOrUpdateData.Location,
	}

	return
}

func (u *courseUsecase) createCourseInfoAction(data *query.CourseInfoQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Create course CourseInfo

		err = u.CommonRepository.Create(tx, data)
		if err != nil {
			return err
		}

		// update course assessment id in main table
		mainQuery := query.CourseMainQueryEntity{
			CourseID: &mainUid,
		}

		mainUpdate := query.CourseMainQueryEntity{
			CourseInfoID: data.ID,
		}
		err = u.CommonRepository.Update(tx, mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
