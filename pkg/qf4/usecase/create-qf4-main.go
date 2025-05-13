package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) CreateQF4Main(data *dto.QF4CreateMainRequestDto) (res *dto.QF4MainResponseDto, err error) {
	createData := query.QF4MainQueryEntity{
		QF4ID:                        &data.QF4ID,
		CourseNumber:                 &data.CourseNumber,
		Version:                      &data.Version,
		FacultyID:                    &data.FacultyID,
		DepartmentName:               &data.DepartmentName,
		EducationYear:                &data.EducationYear,
		QF4CourseInfoID:              data.QF4CourseInfoID,
		QF4LecturerID:                data.QF4LecturerID,
		QF4ResultID:                  data.QF4ResultID,
		QF4CourseTypeAndManagementID: data.QF4CourseTypeAndManagementID,
		QF4AssessmentID:              data.QF4AssessmentID,
		QF4ReferenceID:               data.QF4ReferenceID,
		Status:                       &data.Status,
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, u.createQF4Main(&createData)); err != nil {
		return nil, err
	}

	return &dto.QF4MainResponseDto{
		QF4ID:                        createData.QF4ID,
		CourseNumber:                 createData.CourseNumber,
		Version:                      createData.Version,
		FacultyID:                    createData.FacultyID,
		DepartmentName:               createData.DepartmentName,
		EducationYear:                createData.EducationYear,
		QF4CourseInfoID:              createData.QF4CourseInfoID,
		QF4LecturerID:                createData.QF4LecturerID,
		QF4ResultID:                  createData.QF4ResultID,
		QF4CourseTypeAndManagementID: createData.QF4CourseTypeAndManagementID,
		QF4AssessmentID:              createData.QF4AssessmentID,
		QF4ReferenceID:               createData.QF4ReferenceID,
		Status:                       createData.Status,
		CreatedAt:                    createData.CreatedAt,
		UpdatedAt:                    createData.UpdatedAt,
	}, nil
}

func (u *qf4Usecase) createQF4Main(query *query.QF4MainQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Create(tx, query)
	}
}
