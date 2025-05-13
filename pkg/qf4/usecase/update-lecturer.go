package usecase

import (
	"encoding/json"

	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"
)

func (u *qf4Usecase) UpdateQF4Lecturer(data *dto.QF4UpdateLecturerRequestDto) (res *dto.QF4LecturerResponseDto, err error) {

	byteData, err := json.Marshal(data.CourseMapLecturers)
	if err != nil {
		return nil, err
	}

	courseMapLecture := []query.MapQF4Lecturer{}
	err = json.Unmarshal(byteData, &courseMapLecture)
	if err != nil {
		return nil, err
	}

	updateData := query.QF4LecturerQueryEntity{
		CategoryName:      data.CategoryName,
		CourseOwnerID:     data.CourseOwnerID,
		CourseMapLecturer: courseMapLecture,
	}

	helper.ExecuteTransaction(u.CommonRepository, u.createLecturerAction(&updateData, *data.QF4ID))

	byteData2, err := json.Marshal(updateData.CourseMapLecturer)
	if err != nil {
		return nil, err
	}

	courseMapLecture2 := []dto.MapQF4Lecturer{}
	err = json.Unmarshal(byteData2, &courseMapLecture2)
	if err != nil {
		return nil, err
	}

	res = &dto.QF4LecturerResponseDto{
		ID:                 updateData.ID,
		CategoryName:       updateData.CategoryName,
		CourseOwnerID:      data.CourseOwnerID,
		CourseMapLecturers: courseMapLecture2,
		CreatedAt:          updateData.CreatedAt,
		UpdatedAt:          updateData.UpdatedAt,
	}

	return
}

// func (u *qf4Usecase) updateLecturerAction(dataUpdate *query.QF4LecturerQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
// 	return func(tx *gorm.DB) (err error) {
// 		// mainQF4 := query.QF4MainQueryEntity{
// 		// 	QF4ID: &mainUid,
// 		// }
// 		// err = u.CommonRepository.GetFirst(&mainQF4)

// 		// if err != nil {
// 		// 	return err
// 		// }

// 		// if mainQF4.QF4LecturerID == nil {
// 		// 	return errors.New("can't find assessment for this qf")
// 		// }

// 		// queryUpdate := query.QF4Lecturer{
// 		// 	ID: *mainQF4.QF4LecturerID,
// 		// }

// 		if err = u.CommonRepository.Create(tx, dataUpdate); err != nil {
// 			return err
// 		}

// 		return
// 	}
// }
