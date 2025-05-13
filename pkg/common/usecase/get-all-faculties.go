package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
)

func (u commonUsecase) GetAllFaculties() (result []dto.FacultyResponseDto, err error) {
	record := make([]query.Faculty, 0);

	err = u.CommonRepository.GetList(query.Faculty{}, &record, nil)
	if err != nil {
		return nil, err
	}

	result = make([]dto.FacultyResponseDto, 0)

	for _, faculty := range record {
		result = append(result, dto.FacultyResponseDto{
			ID:            faculty.ID,
			FacultyNameEN: faculty.FacultyNameEN,
			FacultyNameTH: faculty.FacultyNameTH,
			University:    faculty.University,
		})
	}

	return result, nil
}
