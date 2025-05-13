package approvals

import (
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *programApprovalUsecase) RejectCurriculumCommitteeResult(programUID *uuid.UUID, request dto.RejectApprovalCommitteeResultRequestDto) (result *dto.GetApprovalCurriculumCommitteeDto, err error) {
	if err := u.ProgramApprovalRepository.RejectCurriculumCommitteeResult(programUID, request); err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, err
	}

	return nil, nil
}
