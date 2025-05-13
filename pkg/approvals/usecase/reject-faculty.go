package approvals

import (
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *programApprovalUsecase) RejectFacultyApprovalByProgramUID(programUID *uuid.UUID, request dto.RejectApprovalRequestDto) (result *dto.GetApprovalFacultyDto, err error) {
	if err = u.ProgramApprovalRepository.RejectApproval(*programUID, constant.FACULTY_APPROVAL_STATUS.EnumIndex(), request); err != nil {
		return nil, err
	}

	result, err = u.ProgramApprovalRepository.GetFacultyApprovalByProgramUID(*programUID)
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	return result, nil
}
