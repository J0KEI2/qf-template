package approvals

import (
	"fmt"
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *approvalRepository) CommitteesSelected(approvalID uint) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	// update current approval
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ID: approvalID}).Updates(&migrateModels.ProgramApproval{IsCommitteeSelected: true}).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	return
}
