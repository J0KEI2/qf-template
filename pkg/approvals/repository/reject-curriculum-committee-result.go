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

func (r *approvalRepository) RejectCurriculumCommitteeResult(programUID *uuid.UUID, request dto.RejectApprovalCommitteeResultRequestDto) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	// get current approval
	var currApproval migrateModels.ProgramApproval
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: *programUID, ApprovalStatusLevel: constant.CURRICULUM_COMMITEE_APPROVAL_STATUS.EnumIndex()}).Take(&currApproval).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// update curr sub with user id
	committeeUserID, err := uuid.Parse(request.CommitteeUserID)
	if err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}
	updatedBy, err := uuid.Parse(request.UpdatedBy)
	if err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	var maxSubmission migrateModels.ProgramApprovalSubmission
	if err = r.MainDbConn.Raw(getMaxCommitteeSubmissionQuery, programUID, constant.CURRICULUM_COMMITEE_APPROVAL_STATUS.EnumIndex(), committeeUserID, programUID, constant.CURRICULUM_COMMITEE_APPROVAL_STATUS.EnumIndex(), committeeUserID).Scan(&maxSubmission).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	fmt.Printf("\n max submission: %+v \n\n", maxSubmission)

	if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Where(&migrateModels.ProgramApprovalSubmission{
		ApprovalID:      currApproval.ID,
		CommitteeUserID: &committeeUserID,
		SubmissionCount: maxSubmission.SubmissionCount,
	}).Updates(resolveRejectCurrentCurriculumSubmission(request.Comment, updatedBy)).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// create next submission
	if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Create(resolveRejectNextCurriculumSubmission(currApproval.ID, maxSubmission.SubmissionCount+1, committeeUserID)).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	return
}
func resolveRejectCurrentCurriculumSubmission(comment string, updatedBy uuid.UUID) map[string]interface{} {
	return map[string]interface{}{
		"is_current":  false,
		"is_approved": false,
		"is_rejected": true,
		"comment":     &comment,
		"updated_by":  &updatedBy,
	}
}

func resolveRejectNextCurriculumSubmission(approvalID, submissionCount uint, committeeUserID uuid.UUID) *migrateModels.ProgramApprovalSubmission {
	return &migrateModels.ProgramApprovalSubmission{
		ApprovalID:      approvalID,
		SubmissionCount: submissionCount,
		IsCurrent:       true,
		IsApproved:      false,
		IsRejected:      false,
		CommitteeUserID: &committeeUserID,
	}
}
