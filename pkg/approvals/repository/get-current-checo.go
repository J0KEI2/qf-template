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

func (r *approvalRepository) GetCurrentCHECOProgress(programUID uuid.UUID) (currProgress *uint, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	var currCHECO migrateModels.CHECO
	var checoProgress uint
	if err = r.MainDbConn.Model(&migrateModels.CHECO{}).Where(&migrateModels.CHECO{ProgramUID: programUID}).Order("id DESC").Take(&currCHECO).Error; err != gorm.ErrRecordNotFound {
		if err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return nil, err
		}
	} else {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		checoProgress = uint(0)
		return &checoProgress, nil
	}

	checoProgress = currCHECO.StatusID

	return &checoProgress, nil
}
