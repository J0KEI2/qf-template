package usecase

import (
	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	"gorm.io/gorm"
)

func (u commonUsecase) GetFacultyByFacultyName(facultyName string) (result *dto.FacultyResponseDto, err error) {
	record := query.Faculty{
		FacultyNameTH: &facultyName,
	}

	err = u.CommonRepository.GetFirst(&record)
	if err == gorm.ErrRecordNotFound {
		record = query.Faculty{
			ID: pointer.ToUint(0),
		}
		err = u.CommonRepository.GetFirst(&record)
	}

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
