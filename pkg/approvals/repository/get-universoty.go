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

func (r *approvalRepository) GetUniversityApprovalByProgramUID(programUID uuid.UUID) (*dto.GetApprovalUniversityDto, error) {
	var approvalData migrateModels.ProgramApproval

	if r.MainDbConn == nil {
		err := fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return nil, err
	}

	if err := r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, ApprovalStatusLevel: constant.UNIVERSITY_COUNCIL_APPROVAL_STATUS.EnumIndex()}).Take(&approvalData).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, err
	}

	var submissions []dto.ApprovalSubmissionDto
	if err := r.MainDbConn.Raw(submissionQuery, approvalData.ID).Scan(&submissions).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, err
	}

	universityApproval := dto.GetApprovalUniversityDto{
		ID:                  approvalData.ID,
		NameTh:              approvalData.NameTh,
		NameEn:              approvalData.NameEn,
		ApprovalStatusLevel: approvalData.ApprovalStatusLevel,
		ProgramUID:          approvalData.ProgramUID.String(),
		IsCurrent:           approvalData.IsCurrent,
		IsApproved:          approvalData.IsApproved,
		IsRejected:          approvalData.IsRejected,
		Submissions:         submissions,
		ApprovedDate:        submissions[len(submissions)-1].ApprovedDate,
		MeetingNo:           &submissions[len(submissions)-1].MeetingNo,
		CreatedAt:           approvalData.CreatedAt,
		UpdatedAt:           approvalData.UpdatedAt,
	}

	return &universityApproval, nil
}
