package approvals

import (
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *programApprovalUsecase) ApproveAcademicApprovalByProgramUID(programUID *uuid.UUID, request dto.ApproveApprovalRequestDto) (result *dto.GetApprovalAcademicDto, err error) {
	if err = u.ProgramApprovalRepository.ApproveApproval(*programUID, constant.ACADEMIC_COUNCIL_APPROVAL_STATUS.EnumIndex(), request); err != nil {
		return nil, err
	}

	return u.ProgramApprovalRepository.GetAcademicApprovalByProgramUID(*programUID)
}
