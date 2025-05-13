package usecase

import (
	"encoding/json"

	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

func (u *courseUsecase) UpdateCourseLecturer(data *dto.CourseUpdateLecturerRequestDto) (res *dto.CourseLecturerResponseDto, err error) {

	byteData, err := json.Marshal(data.CourseMapLecturers)
	if err != nil {
		return nil, err
	}

	courseMapLecture := []query.MapCourseLecturer{}
	err = json.Unmarshal(byteData, &courseMapLecture)
	if err != nil {
		return nil, err
	}

	updateData := query.CourseLecturerQueryEntity{
		CategoryName:      data.CategoryName,
		CourseOwnerID:     data.CourseOwnerID,
		CourseMapLecturer: courseMapLecture,
	}

	helper.ExecuteTransaction(u.CommonRepository, u.createLecturerAction(&updateData, *data.CourseID))

	byteData2, err := json.Marshal(updateData.CourseMapLecturer)
	if err != nil {
		return nil, err
	}

	courseMapLecture2 := []dto.MapCourseLecturer{}
	err = json.Unmarshal(byteData2, &courseMapLecture2)
	if err != nil {
		return nil, err
	}

	res = &dto.CourseLecturerResponseDto{
		ID:                 updateData.ID,
		CategoryName:       updateData.CategoryName,
		CourseOwnerID:      data.CourseOwnerID,
		CourseMapLecturers: courseMapLecture2,
		CreatedAt:          updateData.CreatedAt,
		UpdatedAt:          updateData.UpdatedAt,
	}

	return
}

// func (u *courseUsecase) updateLecturerAction(dataUpdate *query.CourseLecturerQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
// 	return func(tx *gorm.DB) (err error) {
// 		mainCourse := query.CourseMainQueryEntity{
// 			CourseID: &mainUid,
// 		}
// 		err = u.CommonRepository.GetFirst(&mainCourse)

// 		if err != nil {
// 			return err
// 		}

// 		if mainCourse.CourseLecturerID == nil {
// 			return errors.New("can't find assessment for this qf")
// 		}

// 		queryUpdate := query.CourseLecturer{
// 			ID: *mainCourse.CourseLecturerID,
// 		}
// 		err = u.CommonRepository.Update(tx, queryUpdate, dataUpdate)

// 		if err != nil {
// 			return err
// 		}

// 		return
// 	}
// }
