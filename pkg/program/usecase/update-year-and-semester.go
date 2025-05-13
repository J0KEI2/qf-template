package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) UpdateYearAndSemester(request dto.ProgramYearAndSemesterRequestDto, yearAndSemesterId uint) (result *dto.ProgramYearAndSemesterResponseDto, err error) {
	yearAndSemesterStatement := query.ProgramYearAndSemesterQueryEntity{
		ID: &yearAndSemesterId,
	}

	yearAndSemesterUpdate := query.ProgramYearAndSemesterQueryEntity{
		ProgramSubPlanID: request.SubPlanID,
		Year:             request.Year,
		Semester:         request.Semester,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.UpdateYearAndSemesterTransaction(&yearAndSemesterStatement, &yearAndSemesterUpdate))

	if err != nil {
		return nil, err
	}

	result = &dto.ProgramYearAndSemesterResponseDto{
		ID:        yearAndSemesterUpdate.ID,
		Year:      yearAndSemesterUpdate.Year,
		SubPlanID: yearAndSemesterUpdate.ProgramSubPlanID,
		Semester:  yearAndSemesterUpdate.Semester,
		CreatedAt: yearAndSemesterUpdate.CreatedAt,
		UpdatedAt: yearAndSemesterUpdate.UpdatedAt,
		DeletedAt: yearAndSemesterUpdate.DeletedAt,
	}

	return result, err
}

func (u programUsecase) UpdateYearAndSemesterTransaction(statement *query.ProgramYearAndSemesterQueryEntity, update *query.ProgramYearAndSemesterQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
