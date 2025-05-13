package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetQualityAssurance(ProgramID uuid.UUID) (result *dto.ProgramQualityAssurance, err error) {

	programMain := query.ProgramMainQueryEntity{
		ID: &ProgramID,
	}

	if err = u.CommonRepository.GetFirst(&programMain, "ProgramQualityAssurance"); err != nil {
		return
	}

	result = &dto.ProgramQualityAssurance{
		HESC: dto.QualityAssuranceData{
			IsCheck:     programMain.ProgramQualityAssurance.IsHescCheck,
			Description: programMain.ProgramQualityAssurance.HescDescription,
		},
		AunQa: dto.QualityAssuranceData{
			IsCheck:     programMain.ProgramQualityAssurance.IsAunQaCheck,
			Description: programMain.ProgramQualityAssurance.AunQaDescription,
		},
		ABET: dto.QualityAssuranceData{
			IsCheck:     programMain.ProgramQualityAssurance.IsAbetCheck,
			Description: programMain.ProgramQualityAssurance.AbetDescription,
		},
		WFME: dto.QualityAssuranceData{
			IsCheck:     programMain.ProgramQualityAssurance.IsWfmeCheck,
			Description: programMain.ProgramQualityAssurance.WfmeDescription,
		},
		AACSB: dto.QualityAssuranceData{
			IsCheck:     programMain.ProgramQualityAssurance.IsAacsbCheck,
			Description: programMain.ProgramQualityAssurance.AacsbDescription,
		},
	}

	return
}
