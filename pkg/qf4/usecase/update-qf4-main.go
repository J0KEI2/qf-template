package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) UpdateQF4Main(data *dto.QF4UpdateMainRequestDto) (res *dto.QF4MainResponseDto, err error) {

	updateData := query.QF4MainQueryEntity{
		QF4ID:                        data.QF4ID,
		CourseNumber:                 data.CourseNumber,
		Version:                      data.Version,
		FacultyID:                    data.FacultyID,
		DepartmentName:               data.DepartmentName,
		EducationYear:                data.EducationYear,
		QF4CourseInfoID:              data.QF4CourseInfoID,
		QF4LecturerID:                data.QF4LecturerID,
		QF4ResultID:                  data.QF4ResultID,
		QF4CourseTypeAndManagementID: data.QF4CourseTypeAndManagementID,
		QF4AssessmentID:              data.QF4AssessmentID,
		QF4ReferenceID:               data.QF4ReferenceID,
		Status:                       data.Status,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateMainAction(&updateData, *data.QF4ID))

	res = &dto.QF4MainResponseDto{
		QF4ID:                        data.QF4ID,
		CourseNumber:                 data.CourseNumber,
		Version:                      data.Version,
		FacultyID:                    data.FacultyID,
		DepartmentName:               data.DepartmentName,
		EducationYear:                data.EducationYear,
		QF4CourseInfoID:              data.QF4CourseInfoID,
		QF4LecturerID:                data.QF4LecturerID,
		QF4ResultID:                  data.QF4ResultID,
		QF4CourseTypeAndManagementID: data.QF4CourseTypeAndManagementID,
		QF4AssessmentID:              data.QF4AssessmentID,
		QF4ReferenceID:               data.QF4ReferenceID,
		Status:                       data.Status,
		CreatedAt:                    updateData.CreatedAt,
		UpdatedAt:                    updateData.UpdatedAt,
	}

	return
}

func (u *qf4Usecase) updateMainAction(dataUpdate *query.QF4MainQueryEntity, mainUid uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		mainQF4 := query.QF4MainQueryEntity{
			QF4ID: &mainUid,
		}

		err = u.CommonRepository.Update(tx, mainQF4, dataUpdate)

		if err != nil {
			return err
		}

		return
	}
}
