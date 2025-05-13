package approvals

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *approvalRepository) ApproveApproval(programUID uuid.UUID, approvalStatusLevel uint, request dto.ApproveApprovalRequestDto) (err error) {
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
		if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Where(&migrateModels.ProgramApprovalSubmission{ID: request.SubmissionID}).Updates(resolveApproveCurrSubmission(request, updatedBy, committeeUserID)).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
	} else {
		var maxSubmission migrateModels.ProgramApprovalSubmission
		if err = r.MainDbConn.Raw(getMaxSubmissionQuery, programUID, approvalStatusLevel, programUID, approvalStatusLevel).Scan(&maxSubmission).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
		if err = r.MainDbConn.Model(&migrateModels.ProgramApprovalSubmission{}).Where(&migrateModels.ProgramApprovalSubmission{ID: maxSubmission.ID}).Updates(resolveApproveCurrSubmission(request, updatedBy, committeeUserID)).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
	}

	if approvalStatusLevel < 4 {
		// update current approval
		if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, ApprovalStatusLevel: approvalStatusLevel}).Updates(resolveApproveCurrApproval()).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}

		// get next approval level
		var nextApproval migrateModels.ProgramApproval
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
	} else {
		// update current approval
		if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, ApprovalStatusLevel: approvalStatusLevel}).Updates(resolveApproveCurrApproval()).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
	}

	return
}

// func resolveApproveCurrSubmission(request dto.ApproveApprovalRequestDto, updatedBy, commiteeUserID uuid.UUID) *migrateModels.ProgramApprovalSubmission {
func resolveApproveCurrSubmission(request dto.ApproveApprovalRequestDto, updatedBy, commiteeUserID uuid.UUID) map[string]interface{} {
	approvedDate := time.Now()

	return map[string]interface{}{
		"is_current":  false,
		"is_approved": true,
		"is_rejected": false,
		"comment":     &request.Comment,
		"updated_by":  &updatedBy,
		"committee_user_id": func() *uuid.UUID {
			if commiteeUserID != uuid.Nil {
				return &commiteeUserID
			}
			return nil
		}(),
		"meeting_no": func() *string {
			if request.MeetingNo != nil {
				return request.MeetingNo
			}
			return nil
		}(),
		"approved_date": func() *time.Time {
			if request.ApprovedDate != nil {
				return request.ApprovedDate
			}
			return &approvedDate
		}(),
	}
}

func resolveApproveCurrCommitteeResult(request dto.ApproveApprovalRequestDto) map[string]interface{} {
	approvedDate := time.Now()
	return map[string]interface{}{
		"meeting_no": func() *string {
			if request.MeetingNo != nil {
				return request.MeetingNo
			}
			return nil
		}(),
		"approved_date": func() *time.Time {
			if request.ApprovedDate != nil {
				return request.ApprovedDate
			}
			return &approvedDate
		}(),
	}
}

func resolveApproveCurrApproval() map[string]interface{} {
	return map[string]interface{}{
		"is_current":  false,
		"is_approved": true,
		"is_rejected": false,
	}
}

func resolveApproveNextSubmission() map[string]interface{} {
	return map[string]interface{}{
		"is_current":  true,
		"is_approved": false,
		"is_rejected": false,
	}
}

func resolveApproveNextApproval() map[string]interface{} {
	return map[string]interface{}{
		"is_current":  true,
		"is_approved": false,
		"is_rejected": false,
	}
}
