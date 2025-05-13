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

func (r *approvalRepository) GetCurriculumCommitteeApprovalByProgramUID(programUID, userUID uuid.UUID) (*migrateModels.ProgramApproval, []dto.GetApprovalCommitteesDto, error) {
	var curriculumApproval migrateModels.ProgramApproval
	committees := make([]dto.GetApprovalCommitteesDto, 0)
	if r.MainDbConn == nil {
		err := fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return nil, committees, err
	}

	// get approval
	if err := r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, ApprovalStatusLevel: constant.CURRICULUM_COMMITEE_APPROVAL_STATUS.EnumIndex()}).Take(&curriculumApproval).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, committees, err
	}

	// get committees -> id, 2name, email
	var currSubmissionQuery []dto.CommitteeResultDto
	if err := r.MainDbConn.Raw(submissionQueryWithCommittee, curriculumApproval.ID, userUID).Find(&currSubmissionQuery).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, committees, err
	}

	committeeResults := map[string][]dto.CommitteeResultDto{}
	for _, committeeKey := range currSubmissionQuery {
		committeeResults[committeeKey.CommitteeUserID.String()] = append(committeeResults[committeeKey.CommitteeUserID.String()], committeeKey)
	}

	for _, committeeResult := range committeeResults {
		committee := dto.GetApprovalCommitteesDto{
			UserID:       committeeResult[0].CommitteeUserID,
			NameTh:       *committeeResult[0].CommitteeNameTh,
			NameEn:       *committeeResult[0].CommitteeNameEn,
			Email:        *committeeResult[0].CommitteeEmail,
			CreatedAt:    committeeResult[0].CreatedAt,
			ApprovedDate: committeeResult[0].ApprovedDate,
			MeetingNo:    committeeResult[0].MeetingNo,
		}
		for _, resultValue := range committeeResult {
			committee.CommitteeResults = append(committee.CommitteeResults, dto.ApprovalSubmissionDto{
				ID:              resultValue.ID,
				SubmissionCount: resultValue.SubmissionCount,
				IsCurrent:       resultValue.IsCurrent,
				IsApproved:      resultValue.IsApproved,
				IsRejected:      resultValue.IsRejected,
				Comment: func() string {
					if resultValue.Comment == nil {
						return ""
					}
					return *resultValue.Comment
				}(),
				ApproverNameTh: func() string {
					if resultValue.ApproverNameTh == nil {
						return ""
					}
					return *resultValue.ApproverNameTh
				}(),
				ApproverNameEn: func() string {
					if resultValue.ApproverNameEn == nil {
						return ""
					}
					return *resultValue.ApproverNameEn
				}(),
				ApproverEmail: func() string {
					if resultValue.ApproverEmail == nil {
						return ""
					}
					return *resultValue.ApproverEmail
				}(),
				ApprovedDate: &resultValue.ApprovedDate,
				MeetingNo:    resultValue.MeetingNo,
				CreatedAt:    resultValue.CreatedAt,
				UpdatedAt:    resultValue.UpdatedAt,
			})
		}
		committee.IsCurrent = committeeResult[len(committeeResult)-1].IsCurrent
		committee.IsApproved = committeeResult[len(committeeResult)-1].IsApproved
		committee.IsRejected = committeeResult[len(committeeResult)-1].IsRejected
		committee.UpdatedAt = committeeResult[len(committeeResult)-1].UpdatedAt
		committees = append(committees, committee)
	}

	return &curriculumApproval, committees, nil
}
