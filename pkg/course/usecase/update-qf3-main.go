package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) UpdateCourseMain(data *dto.CourseUpdateMainRequestDto) (res *dto.CourseMainResponseDto, err error) {

	updateData := query.CourseMainQueryEntity{
		CourseID:                        data.CourseID,
		CourseNumber:                 data.CourseNumber,
		Version:                      data.Version,
		FacultyID:                    data.FacultyID,
		DepartmentName:               data.DepartmentName,
		EducationYear:                data.EducationYear,
		CourseInfoID:              data.CourseInfoID,
		CourseLecturerID:                data.CourseLecturerID,
		CourseResultID:                  data.CourseResultID,
		CourseTypeAndManagementID: data.CourseTypeAndManagementID,
		CourseAssessmentID:              data.CourseAssessmentID,
		CourseReferenceID:               data.CourseReferenceID,
		Status:                       data.Status,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateMainAction(&updateData, *data.CourseID))

	res = &dto.CourseMainResponseDto{
		CourseID:                        data.CourseID,
		CourseNumber:                 data.CourseNumber,
		Version:                      data.Version,
		FacultyID:                    data.FacultyID,
		DepartmentName:               data.DepartmentName,
		EducationYear:                data.EducationYear,
		CourseInfoID:              data.CourseInfoID,
		CourseLecturerID:                data.CourseLecturerID,
		CourseResultID:                  data.CourseResultID,
		CourseTypeAndManagementID: data.CourseTypeAndManagementID,
		CourseAssessmentID:              data.CourseAssessmentID,
		CourseReferenceID:               data.CourseReferenceID,
		Status:                       data.Status,
		CreatedAt:                    updateData.CreatedAt,
		UpdatedAt:                    updateData.UpdatedAt,
	}

	return
}

func (u *courseUsecase) updateMainAction(dataUpdate *query.CourseMainQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		mainCourse := query.CourseMainQueryEntity{
			CourseID: &mainUid,
		}

		// userUpdateData := map[string]interface{}{
		// 	"CourseNumber":                 dataUpdate.CourseNumber,
		// 	"Version":                      dataUpdate.Version,
		// 	"FacultyID":                    dataUpdate.FacultyID,
		// 	"DepartmentName":               dataUpdate.DepartmentName,
		// 	"EducationYear":                dataUpdate.EducationYear,
		// 	"CourseInfoID":              dataUpdate.CourseInfoID,
		// 	"CourseLecturerID":                dataUpdate.CourseLecturerID,
		// 	"CourseResultID":                  dataUpdate.CourseResultID,
		// 	"CourseTypeAndManagementID": dataUpdate.CourseTypeAndManagementID,
		// 	"CourseAssessmentID":              dataUpdate.CourseAssessmentID,
		// 	"CourseReferenceID":               dataUpdate.CourseReferenceID,
		// 	"Status":                       dataUpdate.Status,
		// }

		err = u.CommonRepository.Update(tx, mainCourse, dataUpdate)

		if err != nil {
			return err
		}

		return
	}
}
