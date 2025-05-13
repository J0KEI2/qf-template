package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
)

func (u commonUsecase) GetFacultyByID(id uint) (result *dto.FacultyResponseDto, err error) {
	record := query.Faculty{
		ID: &id,
	}

	err = u.CommonRepository.GetFirst(&record)

	if err != nil {
		return nil, err
	}

	result = &dto.FacultyResponseDto{
		ID:            record.ID,
		FacultyNameEN: record.FacultyNameEN,
		FacultyNameTH: record.FacultyNameTH,
		University:    record.University,
	}

	return result, nil
}
