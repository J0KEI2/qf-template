package usecase

import (
	"log"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	permissionConstant "github.com/zercle/kku-qf-services/pkg/constant/permission"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	commonQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	rapQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u programUsecase) CreateMainProgram(userUID uuid.UUID, request dto.ProgramMainRequestDto) (result *dto.ProgramMainPagination, err error) {
	facultyQuery := commonQuery.Faculty{
		ID: &request.FacultyID,
	}
	if err = u.CommonRepository.GetFirst(&facultyQuery); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	queryResult := &programQuery.ProgramMainQueryEntity{}
	if err = helper.ExecuteTransaction(u.CommonRepository, u.CreateMainProgramTransaction(userUID, request, queryResult)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	result = &dto.ProgramMainPagination{
		ProgramMainID: queryResult.ID,
		FacultyNameTH: facultyQuery.FacultyNameTH,
		FacultyNameEN: facultyQuery.FacultyNameEN,
		ProgramNameTH: &request.ProgramNameTH,
		ProgramNameEN: &request.ProgramNameEN,
		BranchNameTH:  &request.BranchNameTH,
		BranchNameEN:  &request.BranchNameEN,
		ProgramType:   &request.ProgramType,
		ProgramTypeID: &request.ProgramTypeID,
		ProgramYear:   &request.ProgramYear,
		ProgramYearID: &request.ProgramYearID,
	}
	return
}

func (u programUsecase) CreateMainProgramTransaction(userUID uuid.UUID, request dto.ProgramMainRequestDto, programMainQuery *programQuery.ProgramMainQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		programGeneralDetail := programQuery.ProgramGeneralDetailQueryEntity{
			UniversityName: pointer.ToString(constant.UNIVERSITY_NAME_TH),
			FacultyID:      pointer.ToUint(request.FacultyID),
			ProgramNameTH:  pointer.ToString(request.ProgramNameTH),
			ProgramNameEN:  pointer.ToString(request.ProgramNameEN),
			BranchNameTH:   pointer.ToString(request.BranchNameTH),
			BranchNameEN:   pointer.ToString(request.BranchNameEN),
			ProgramType:    pointer.ToString(request.ProgramType),
			ProgramTypeID:  pointer.ToInt(request.ProgramTypeID),
			ProgramYear:    pointer.ToUint(request.ProgramYear),
			ProgramYearID:  pointer.ToInt(request.ProgramYearID),
		}

		if err = u.CommonRepository.Create(tx, &programGeneralDetail); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		programPolicyAndStrategic := programQuery.ProgramPolicyAndStrategicQueryEntity{
			ProgramPhilosophy: nil,
			ProgramObjective:  nil,
			ProgramPolicy:     nil,
			ProgramStrategic:  nil,
			ProgramRisk:       nil,
			ProgramFeedback:   nil,
		}

		if err = u.CommonRepository.Create(tx, &programPolicyAndStrategic); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		programQualityAssurance := programQuery.ProgramQualityAssuranceQueryEntity{
			IsHescCheck:      nil,
			HescDescription:  nil,
			IsAunQaCheck:     nil,
			AunQaDescription: nil,
			IsAbetCheck:      nil,
			AbetDescription:  nil,
			IsWfmeCheck:      nil,
			WfmeDescription:  nil,
			IsAacsbCheck:     nil,
			AacsbDescription: nil,
		}

		if err = u.CommonRepository.Create(tx, &programQualityAssurance); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		programSystemAndMechanic := programQuery.ProgramSystemAndMechanicQueryEntity{
			CoursePolicies:          nil,
			CourseStrategies:        nil,
			CourseRisk:              nil,
			CourseStudentComment:    nil,
			CourseExpectedAttribute: nil,
			MainContentAndStructure: nil,
			CourseImprovingPlan:     nil,
		}

		if err = u.CommonRepository.Create(tx, &programSystemAndMechanic); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		programPlanAndEvaluate := programQuery.ProgramPlanAndEvaluateQueryEntity{}

		if err = u.CommonRepository.Create(tx, &programPlanAndEvaluate); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		programUID := uuid.New()
		programMainQuery.ID = &programUID
		programMainQuery.ProgramGeneralDetailID = programGeneralDetail.ID
		programMainQuery.ProgramPolicyAndStrategicID = programPolicyAndStrategic.ID
		programMainQuery.ProgramQualityAssuranceID = programQualityAssurance.ID
		programMainQuery.ProgramPlanAndEvaluateID = programPlanAndEvaluate.ID
		programMainQuery.ProgramSystemAndMechanicID = programSystemAndMechanic.ID

		if err = u.CommonRepository.Create(tx, programMainQuery); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		role := rapQuery.RoleQueryEntity{
			RoleNameEN: pointer.ToString(permissionConstant.ROLE_PROGRAM_OWNER),
		}

		u.CommonRepository.GetFirst(&role)

		programPermission := rapQuery.MapProgramsRolesQueryEntity{
			UserID:    &userUID,
			ProgramID: &programUID,
			RoleID:    role.ID,
		}
		if err = u.CommonRepository.Create(tx, &programPermission); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		// create approvals
		if err := u.InitProgramApprovals(programUID, tx); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		return nil
	}
}
