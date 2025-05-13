package approvals

import (
	"sync"

	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type programApprovalUsecase struct {
	domain.ProgramApprovalRepository
	domain.CommonRepository
	mutex sync.Mutex
}

func NewProgramApprovalUsecase(repo domain.ProgramApprovalRepository, common domain.CommonRepository) domain.ProgramApprovalUsecase {
	return &programApprovalUsecase{
		ProgramApprovalRepository: repo,
		CommonRepository:          common,
	}
}

func (u *programApprovalUsecase) CheckIfApprovalEditable(programApprovalLevel, userApprovalLevel uint) (isEditable bool) {
	if programApprovalLevel == userApprovalLevel {
		return true
	} else if userApprovalLevel == uint(constant.ADMIN_APPROVAL_LEVEL) {
		return true
	} else {
		return false
	}
}
