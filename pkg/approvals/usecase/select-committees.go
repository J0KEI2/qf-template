package approvals

import (
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *programApprovalUsecase) SelectCurriculumCommittees(programUID *uuid.UUID, committeeUserIDs []string) (result *dto.GetApprovalCurriculumCommitteeDto, err error) {
	if err := u.ProgramApprovalRepository.SelectCurriculumCommittees(programUID, committeeUserIDs); err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, err
	}

	return nil, nil
}
