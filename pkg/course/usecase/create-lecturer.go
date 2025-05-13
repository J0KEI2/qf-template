package usecase

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) CreateOrUpdateCourseLecturer(data *dto.CourseCreateLecturerRequestDto) (res *dto.CourseLecturerResponseDto, err error) {
	if err != nil {
		return nil, err
	}

	queryTb := query.CourseMainQueryEntity{
		CourseID: &data.CourseID,
	}

	u.CommonRepository.GetFirst(&queryTb)

	byteData, err := json.Marshal(data.CourseMapLecturers)
	if err != nil {
		return nil, err
	}

	courseMapLecture := []query.MapCourseLecturer{}
	err = json.Unmarshal(byteData, &courseMapLecture)
	if err != nil {
		return nil, err
	}

	createOrUpdateData := query.CourseLecturerQueryEntity{
		CategoryName:      &data.CategoryName,
		CourseOwnerID:     &data.CourseOwnerID,
		CourseMapLecturer: courseMapLecture,
	}

	// if queryTb.CourseLecturerID == nil {
	err = helper.ExecuteTransaction(u.CommonRepository, u.createLecturerAction(&createOrUpdateData, data.CourseID))
	// } else {
	// 	err = helper.ExecuteTransaction(u.CommonRepository, u.createLecturerAction(&createOrUpdateData, data.CourseID))
	// }

	if err != nil {
		return nil, err
	}

	byteData2, err := json.Marshal(createOrUpdateData.CourseMapLecturer)
	if err != nil {
		return nil, err
	}

	courseMapLecture2 := []dto.MapCourseLecturer{}
	err = json.Unmarshal(byteData2, &courseMapLecture2)
	if err != nil {
		return nil, err
	}

	res = &dto.CourseLecturerResponseDto{
		ID:                 createOrUpdateData.ID,
		CategoryName:       createOrUpdateData.CategoryName,
		CourseOwnerID:      createOrUpdateData.CourseOwnerID,
		CourseMapLecturers: courseMapLecture2,
	}

	return
}

func (u *courseUsecase) createLecturerAction(data *query.CourseLecturerQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Create course Lecturer

		err = u.CommonRepository.Create(tx, data)
		if err != nil {
			return err
		}

		// update course assessment id in main table
		mainQuery := query.CourseMainQueryEntity{
			CourseID: &mainUid,
		}

		mainUpdate := query.CourseMainQueryEntity{
			CourseLecturerID: data.ID,
		}
		err = u.CommonRepository.Update(tx, mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
