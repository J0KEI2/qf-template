package usecase

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) CreateOrUpdateQF4Lecturer(data *dto.QF4CreateLecturerRequestDto) (res *dto.QF4LecturerResponseDto, err error) {
	queryTb := query.QF4MainQueryEntity{
		QF4ID: &data.QF4ID,
	}

	u.CommonRepository.GetFirst(&queryTb)

	byteData, err := json.Marshal(data.CourseMapLecturer)
	if err != nil {
		return nil, err
	}

	courseMapLecture := []query.MapQF4Lecturer{}
	err = json.Unmarshal(byteData, &courseMapLecture)
	if err != nil {
		return nil, err
	}

	createOrUpdateData := query.QF4LecturerQueryEntity{
		CategoryName:      &data.CategoryName,
		CourseOwnerID:     &data.CourseOwnerID,
		CourseMapLecturer: courseMapLecture,
	}

	// if queryTb.QF4LecturerID == nil {
	err = helper.ExecuteTransaction(u.CommonRepository, u.createLecturerAction(&createOrUpdateData, data.QF4ID))
	// } else {
	// err = helper.ExecuteTransaction(u.CommonRepository, u.updateLecturerAction(&createOrUpdateData, data.QF4ID))
	// }

	if err != nil {
		return nil, err
	}

	byteData2, err := json.Marshal(createOrUpdateData.CourseMapLecturer)
	if err != nil {
		return nil, err
	}

	courseMapLecture2 := []dto.MapQF4Lecturer{}
	err = json.Unmarshal(byteData2, &courseMapLecture2)
	if err != nil {
		return nil, err
	}

	res = &dto.QF4LecturerResponseDto{
		ID:                 createOrUpdateData.ID,
		CategoryName:       createOrUpdateData.CategoryName,
		CourseOwnerID:      createOrUpdateData.CourseOwnerID,
		CourseMapLecturers: courseMapLecture2,
	}

	return
}

func (u *qf4Usecase) createLecturerAction(data *query.QF4LecturerQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Create course Lecturer

		err = u.CommonRepository.Create(tx, data)
		if err != nil {
			return err
		}

		// update course assessment id in main table
		mainQuery := query.QF4MainQueryEntity{
			QF4ID: &mainUid,
		}

		mainUpdate := query.QF4MainQueryEntity{
			QF4LecturerID: data.ID,
		}
		err = u.CommonRepository.Update(tx, mainQuery, &mainUpdate)

		if err != nil {
			return err
		}

		return
	}
}
