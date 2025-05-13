package approvals

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *approvalRepository) ApproveCurriculumCommitteeApproval(programUID uuid.UUID, approvalStatusLevel uint, request dto.ApproveApprovalRequestDto) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	var currApproval, nextApproval migrateModels.ProgramApproval

	// get current approval
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, ApprovalStatusLevel: approvalStatusLevel}).Take(&currApproval).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// update current committee results
	if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Where(&migrateModels.ProgramApprovalSubmission{ApprovalID: currApproval.ID}).Updates(resolveApproveCurrCommitteeResult(request)).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// update current approval
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, ApprovalStatusLevel: approvalStatusLevel}).Updates(resolveApproveCurrApproval()).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// get next approval level
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, ApprovalStatusLevel: approvalStatusLevel + 1}).Take(&nextApproval).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// update next approval
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ID: nextApproval.ID}).Updates(resolveApproveNextApproval()).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// update next submission
	if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Where(&migrateModels.ProgramApprovalSubmission{ApprovalID: nextApproval.ID, SubmissionCount: 1}).Updates(resolveApproveNextSubmission()).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	return
}
