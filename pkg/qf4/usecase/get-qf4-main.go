package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"
)

// TODO: create the new func don't forget to embed in domain use case interface
func (useCase qf4Usecase) GetQF4Main(data *dto.QF4GetMainRequestDto) (response *dto.QF4MainResponseDto, err error) {
	convertUid, err := uuid.Parse(data.QF4UID)

	if err != nil {
		return nil, err
	}

	res := query.QF4MainQueryEntity{
		QF4ID: &convertUid,
	}

	err = useCase.CommonRepository.GetFirst(&res, "Faculty")

	if err != nil {
		return nil, err
	}

	course := dto.QF4MainResponseDto{
		CreatedAt:                    res.CreatedAt,
		UpdatedAt:                    res.UpdatedAt,
		QF4ID:                        res.QF4ID,
		CourseNumber:                 res.CourseNumber,
		Version:                      res.Version,
		FacultyID:                    res.FacultyID,
		DepartmentName:               res.DepartmentName,
		EducationYear:                res.EducationYear,
		QF4CourseInfoID:              res.QF4CourseInfoID,
		QF4LecturerID:                res.QF4LecturerID,
		QF4ResultID:                  res.QF4ResultID,
		QF4CourseTypeAndManagementID: res.QF4CourseTypeAndManagementID,
		QF4AssessmentID:              res.QF4AssessmentID,
		QF4ReferenceID:               res.QF4ReferenceID,
		Status:                       res.Status,
	}
	return &course, nil
}
