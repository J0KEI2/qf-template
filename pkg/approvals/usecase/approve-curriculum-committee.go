package approvals

import (
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *programApprovalUsecase) ApproveCurriculumCommitteeApprovalByProgramUID(programUID *uuid.UUID, request dto.ApproveApprovalRequestDto, userApprovalLevel uint) (result *dto.GetApprovalCurriculumCommitteeDto, err error) {
	if err = u.ProgramApprovalRepository.ApproveCurriculumCommitteeApproval(*programUID, constant.CURRICULUM_COMMITEE_APPROVAL_STATUS.EnumIndex(), request); err != nil {
		return nil, err
	}

	return u.GetCurriculumCommitteesApprovalByProgramUID(programUID, userApprovalLevel)
}
