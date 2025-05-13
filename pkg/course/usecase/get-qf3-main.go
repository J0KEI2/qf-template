package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

// TODO: create the new func don't forget to embed in domain use case interface
func (useCase courseUsecase) GetCourseMain(data *dto.CourseGetMainRequestDto) (response *dto.CourseMainResponseDto, err error) {
	convertUid, err := uuid.Parse(data.CourseUID)

	if err != nil {
		return nil, err
	}

	res := query.CourseMainQueryEntity{
		CourseID: &convertUid,
	}

	err = useCase.CommonRepository.GetFirst(&res, "Faculty")

	if err != nil {
		return nil, err
	}

	course := dto.CourseMainResponseDto{
		CreatedAt:                    res.CreatedAt,
		UpdatedAt:                    res.UpdatedAt,
		CourseID:                        res.CourseID,
		CourseNumber:                 res.CourseNumber,
		Version:                      res.Version,
		FacultyID:                    res.FacultyID,
		DepartmentName:               res.DepartmentName,
		EducationYear:                res.EducationYear,
		CourseInfoID:              res.CourseInfoID,
		CourseLecturerID:                res.CourseLecturerID,
		CourseResultID:                  res.CourseResultID,
		CourseTypeAndManagementID: res.CourseTypeAndManagementID,
		CourseAssessmentID:              res.CourseAssessmentID,
		CourseReferenceID:               res.CourseReferenceID,
		Status:                       res.Status,
	}
	return &course, nil
}
