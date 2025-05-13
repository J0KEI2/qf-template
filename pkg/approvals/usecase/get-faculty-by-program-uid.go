package approvals

import (
	"log"
	"sort"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *programApprovalUsecase) GetFacultyApprovalByProgramUID(programUID *uuid.UUID, userApprovalLevel uint) (result *dto.GetApprovalFacultyDto, err error) {
	result, err = u.ProgramApprovalRepository.GetFacultyApprovalByProgramUID(*programUID)
	if err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, err
	}

	result.IsEditable = u.CheckIfApprovalEditable(uint(constant.FACULTY_APPROVAL_LEVEL), userApprovalLevel)

	sort.SliceStable(result.Submissions, func(i, j int) bool {
		return result.Submissions[i].ID < result.Submissions[j].ID
	})

	return result, nil
}
