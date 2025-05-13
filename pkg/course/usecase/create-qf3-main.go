package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) CreateCourseMain(data *dto.CourseCreateMainRequestDto) (res *dto.CourseMainResponseDto, err error) {

	if err != nil {
		return nil, err
	}

	createData := query.CourseMainQueryEntity{
		CourseID:                        &data.CourseID,
		CourseNumber:                 &data.CourseNumber,
		Version:                      &data.Version,
		FacultyID:                    &data.FacultyID,
		DepartmentName:               &data.DepartmentName,
		EducationYear:                &data.EducationYear,
		CourseInfoID:              data.CourseInfoID,
		CourseLecturerID:                data.CourseLecturerID,
		CourseResultID:                  data.CourseResultID,
		CourseTypeAndManagementID: data.CourseTypeAndManagementID,
		CourseAssessmentID:              data.CourseAssessmentID,
		CourseReferenceID:               data.CourseReferenceID,
		Status:                       &data.Status,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.createCourseMain(&createData))

	if err != nil {
		return nil, err
	}

	res = &dto.CourseMainResponseDto{
		CourseID:                        createData.CourseID,
		CourseNumber:                 createData.CourseNumber,
		Version:                      createData.Version,
		FacultyID:                    createData.FacultyID,
		DepartmentName:               createData.DepartmentName,
		EducationYear:                createData.EducationYear,
		CourseInfoID:              createData.CourseInfoID,
		CourseLecturerID:                createData.CourseLecturerID,
		CourseResultID:                  createData.CourseResultID,
		CourseTypeAndManagementID: createData.CourseTypeAndManagementID,
		CourseAssessmentID:              createData.CourseAssessmentID,
		CourseReferenceID:               createData.CourseReferenceID,
		Status:                       createData.Status,
		CreatedAt:                    createData.CreatedAt,
		UpdatedAt:                    createData.UpdatedAt,
	}

	return
}

func (u *courseUsecase) createCourseMain(query *query.CourseMainQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Create(tx, query)
	}
}
