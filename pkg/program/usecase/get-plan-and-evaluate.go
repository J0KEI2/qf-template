package usecase

import (
	"encoding/json"

	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) GetPlanAndEvaluate(programSubPlanID uint) (result *dto.ProgramPlanAndEvaluateResponseDto, err error) {
	// Get Program year from programSubPlanID
	programSubPlanQuery := query.ProgramSubPlanQueryEntity{
		ID: &programSubPlanID,
	}

	if err := u.CommonRepository.GetFirst(&programSubPlanQuery); err != nil {
		return nil, err
	}

	programPlanDetailQuery := query.ProgramPlanDetailQueryEntity{
		ID: programSubPlanQuery.ProgramPlanDetailID,
	}

	if err := u.CommonRepository.GetFirst(&programPlanDetailQuery); err != nil {
		return nil, err
	}

	programMajorQuery := query.ProgramMajorQueryEntity{
		ID: programPlanDetailQuery.ProgramMajorID,
	}

	if err := u.CommonRepository.GetFirst(&programMajorQuery); err != nil {
		return nil, err
	}

	programGeneralDetailQuery := query.ProgramGeneralDetailQueryEntity{
		ID: programMajorQuery.ProgramGeneralDetailID,
	}

	if err := u.CommonRepository.GetFirst(&programGeneralDetailQuery); err != nil {
		return nil, err
	}

	// Get Plan and evaluate Detail
	programPlanAndEvaluate := query.ProgramPlanAndEvaluateQueryEntity{
		ProgramSubPlanID: &programSubPlanID,
	}

	if err = u.CommonRepository.GetFirst(&programPlanAndEvaluate); err != nil {
		if err == gorm.ErrRecordNotFound {
			err = helper.ExecuteTransaction(u.CommonRepository, u.InitPlanAndEvaluate(programSubPlanID, *programGeneralDetailQuery.NumberOfYear))
			if err != nil {
				return nil, err
			}
			err = u.CommonRepository.GetFirst(&programPlanAndEvaluate)
			if err != nil {
				return nil, err
			}
		} else {
			return
		}
	}

	var studentCharacteristic interface{}
	if err = json.Unmarshal([]byte(*programPlanAndEvaluate.StudentCharacteristic), &studentCharacteristic); err != nil {
		return nil, err
	}

	var receiveStudentPlan interface{}
	if err = json.Unmarshal([]byte(*programPlanAndEvaluate.ReceiveStudentPlan), &receiveStudentPlan); err != nil {
		return nil, err
	}

	var programIncome interface{}
	if err = json.Unmarshal([]byte(*programPlanAndEvaluate.ProgramIncome), &programIncome); err != nil {
		return nil, err
	}

	var programOutcome interface{}
	if err = json.Unmarshal([]byte(*programPlanAndEvaluate.ProgramOutcome), &programOutcome); err != nil {
		return nil, err
	}

	result = &dto.ProgramPlanAndEvaluateResponseDto{
		ID:                                programPlanAndEvaluate.ID,
		ProgramSubPlanID:                  &programSubPlanID,
		ProgramDegreeTypeID:               programGeneralDetailQuery.ProgramDegreeTypeID,
		ProgramDegreeType:                 programGeneralDetailQuery.ProgramDegreeType,
		ProgramYearID:                     programGeneralDetailQuery.ProgramYearID,
		ProgramYear:                       programGeneralDetailQuery.ProgramYear,
		StudentCharacteristics:            studentCharacteristic,
		ReceiveStudentPlan:                receiveStudentPlan,
		ProgramIncome:                     programIncome,
		ProgramOutcome:                    programOutcome,
		AcademicEvaluation:                programPlanAndEvaluate.AcademicEvaluation,
		GraduationCriteria:                programPlanAndEvaluate.GraduationCriteria,
		ProgramUniversityTransferStandard: programPlanAndEvaluate.ProgramUniversityTransferStandard,
		ProgramPreparation:                programPlanAndEvaluate.ProgramPreparation,
		UpdatedAt:                         programPlanAndEvaluate.UpdatedAt,
		CreatedAt:                         programPlanAndEvaluate.CreatedAt,
	}
	return result, nil
}
