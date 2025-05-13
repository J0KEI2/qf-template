package usecase

import (
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/domain"
	models "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

type programUsecase struct {
	domain.ProgramRepository
	domain.CommonRepository
	domain.HRUseCase
	domain.RegUseCase
	approval domain.ProgramApprovalRepository
}

func NewProgramUsecase(repo domain.ProgramRepository, common domain.CommonRepository, hrUsecase domain.HRUseCase, regUsecase domain.RegUseCase, approvalRepo domain.ProgramApprovalRepository) domain.ProgramUsecase {
	return &programUsecase{
		repo,
		common,
		hrUsecase,
		regUsecase,
		approvalRepo,
	}
}

func (u programUsecase) InitProgramApprovals(programUID uuid.UUID, tx *gorm.DB) error {
	// init faculty_approval and set is_current = true
	approval := models.ProgramApproval{
		ProgramUID:          programUID,
		NameTh:              constant.APPROVAL_STATUS.String(1),
		NameEn:              constant.APPROVAL_STATUS.StringEn(1),
		ApprovalStatusLevel: 1,
		IsCurrent:           true,
		IsRejected:          false,
		IsApproved:          false,
	}
	if err := u.CommonRepository.Create(tx, &approval); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return err
	}
	// init approval submission
	approvalSubmission := models.ProgramApprovalSubmission{
		ApprovalID:      approval.ID,
		SubmissionCount: 1,
		IsCurrent:       true,
		IsRejected:      false,
		IsApproved:      false,
	}
	if err := u.CommonRepository.Create(tx, &approvalSubmission); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return err
	}

	// create other approval with is_current = false
	approval.IsCurrent = false
	approvalSubmission.IsCurrent = false
	for i := 2; i <= 4; i++ {
		approval := models.ProgramApproval{
			ProgramUID:          programUID,
			NameTh:              constant.APPROVAL_STATUS.String(constant.APPROVAL_STATUS(i)),
			NameEn:              constant.APPROVAL_STATUS.StringEn(constant.APPROVAL_STATUS(i)),
			ApprovalStatusLevel: uint(i),
			IsCurrent:           false,
			IsRejected:          false,
			IsApproved:          false,
		}
		if err := u.CommonRepository.Create(tx, &approval); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			return err
		}

		if i != 2 { // to skip curriculum committee
			approvalSubmission := models.ProgramApprovalSubmission{
				ApprovalID:      approval.ID,
				SubmissionCount: 1,
				IsCurrent:       true,
				IsRejected:      false,
				IsApproved:      false,
			}
			if err := u.CommonRepository.Create(tx, &approvalSubmission); err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
				return err
			}
		}
	}
	return nil
}
