package approvals

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

func (r *approvalRepository) GetCurrentApprovalProgress(programUID uuid.UUID) (currProgress *uint, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	var currApproval migrateModels.ProgramApproval
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, IsCurrent: true}).Take(&currApproval).Error; err != nil {
		if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: programUID, IsApproved: true}).Order("approval_status_level DESC").Take(&currApproval).Error; err != gorm.ErrRecordNotFound {
			if err != nil {
				log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
				return nil, err
			}
		} else {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return nil, nil
		}
	}

	currProgressLevel := currApproval.ApprovalStatusLevel
	return &currProgressLevel, nil
}
