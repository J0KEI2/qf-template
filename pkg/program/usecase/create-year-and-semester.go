package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateYearAndSemester(request dto.ProgramYearAndSemesterRequestDto, subPlanID uint) (result *dto.ProgramYearAndSemesterResponseDto, err error) {
	query := query.ProgramYearAndSemesterQueryEntity{
		ProgramSubPlanID: &subPlanID,
		Year: request.Year,
		Semester: request.Semester,
	}
	
	err = helper.ExecuteTransaction(u.CommonRepository, u.CreateYearAndSemesterTransaction(&query))

	if err != nil {
		return nil, err
	}

	result = &dto.ProgramYearAndSemesterResponseDto{
		ID: query.ID,
		Year: query.Year,
		SubPlanID: query.ProgramSubPlanID,
		Semester: query.Semester,
		CreatedAt: query.CreatedAt,
		UpdatedAt: query.UpdatedAt,
		DeletedAt: query.DeletedAt,
	}

	return result, err
}

func (u programUsecase) CreateYearAndSemesterTransaction(yearAndSemester *query.ProgramYearAndSemesterQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Create(tx, yearAndSemester)
	}
}
