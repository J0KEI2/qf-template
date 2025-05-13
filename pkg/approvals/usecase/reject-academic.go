package approvals

import (
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *programApprovalUsecase) RejectAcademicApprovalByProgramUID(programUID *uuid.UUID, request dto.RejectApprovalRequestDto) (result *dto.GetApprovalAcademicDto, err error) {
	if err = u.ProgramApprovalRepository.RejectApproval(*programUID, constant.ACADEMIC_COUNCIL_APPROVAL_STATUS.EnumIndex(), request); err != nil {
		return nil, err
	}

	return u.ProgramApprovalRepository.GetAcademicApprovalByProgramUID(*programUID)
}
