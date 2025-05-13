package approvals

import (
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *programApprovalUsecase) ApproveUniversityApprovalByProgramUID(programUID *uuid.UUID, request dto.ApproveApprovalRequestDto) (result *dto.GetApprovalUniversityDto, err error) {
	if err = u.ProgramApprovalRepository.ApproveApproval(*programUID, constant.UNIVERSITY_COUNCIL_APPROVAL_STATUS.EnumIndex(), request); err != nil {
		return nil, err
	}

	return u.ProgramApprovalRepository.GetUniversityApprovalByProgramUID(*programUID)
}
