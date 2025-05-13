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

func (r *approvalRepository) RejectApproval(programUID uuid.UUID, approvalStatusLevel uint, request dto.RejectApprovalRequestDto) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	updatedBy, err := uuid.Parse(request.UpdatedBy)
	if err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	var committeeUserID uuid.UUID
	if request.CommitteeUserID != "" {
		committeeUserID, err = uuid.Parse(request.CommitteeUserID)
		if err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
	}

	// update current submission
	if request.SubmissionID != 0 {
		if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Where(&migrateModels.ProgramApprovalSubmission{ID: request.SubmissionID}).Updates(resolveRejectCurrSubmission(request, updatedBy, committeeUserID)).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
	} else {
		var maxSubmission migrateModels.ProgramApprovalSubmission
		if err = r.MainDbConn.Raw(getMaxSubmissionQuery, programUID, approvalStatusLevel, programUID, approvalStatusLevel).Scan(&maxSubmission).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
		if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Where(&migrateModels.ProgramApprovalSubmission{ID: maxSubmission.ID}).Updates(resolveRejectCurrSubmission(request, updatedBy, committeeUserID)).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
	}

	// get curr approval level
	var currApproval migrateModels.ProgramApproval
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, ApprovalStatusLevel: approvalStatusLevel}).Take(&currApproval).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// create next submission
	var maxSubmission migrateModels.ProgramApprovalSubmission
	if err = r.MainDbConn.Raw(getMaxSubmissionQuery, programUID, approvalStatusLevel, programUID, approvalStatusLevel).Scan(&maxSubmission).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}
	if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Create(resolveCreateNextSubmission(currApproval.ID, maxSubmission.SubmissionCount+1, updatedBy, committeeUserID)).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	return
}

func resolveRejectCurrSubmission(request dto.RejectApprovalRequestDto, updatedBy, commiteeUserID uuid.UUID) map[string]interface{} {
	return map[string]interface{}{
		"is_current":  false,
		"is_approved": false,
		"is_rejected": true,
		"comment":     &request.Comment,
		"updated_by":  &updatedBy,
		"committee_user_id": func() *uuid.UUID {
			if commiteeUserID != uuid.Nil {
				return &commiteeUserID
			}
			return nil
		}(),
	}
}

func resolveCreateNextSubmission(approvalID, submissionCount uint, updatedBy, committeeUserID uuid.UUID) *migrateModels.ProgramApprovalSubmission {
	return &migrateModels.ProgramApprovalSubmission{
		ApprovalID:      approvalID,
		SubmissionCount: submissionCount,
		IsCurrent:       true,
		IsApproved:      false,
		IsRejected:      false,
		CommitteeUserID: func() *uuid.UUID {
			if committeeUserID == uuid.Nil {
				return nil
			}
			return &committeeUserID
		}(),
		UpdatedBy: &updatedBy,
	}
}
