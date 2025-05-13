package approvals

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *approvalRepository) SelectCurriculumCommittees(programUID *uuid.UUID, committeeUserIDs []string) (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	// get current approval
	var currApproval migrateModels.ProgramApproval
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ProgramUID: *programUID, ApprovalStatusLevel: constant.CURRICULUM_COMMITEE_APPROVAL_STATUS.EnumIndex()}).Take(&currApproval).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	// create submission followed by commitees
	for _, committeeUserIDString := range committeeUserIDs {
		committeeUserID, err := uuid.Parse(committeeUserIDString)
		if err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
		if err = r.MainDbConn.Create(&migrateModels.ProgramApprovalSubmission{
			ApprovalID:      currApproval.ID,
			SubmissionCount: 1,
			IsCurrent:       true,
			CommitteeUserID: &committeeUserID,
		}).Error; err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
			return err
		}
	}

	// update curr approval to selected
	if err = r.MainDbConn.Model(&migrateModels.ProgramApproval{}).Where(&migrateModels.ProgramApproval{ID: currApproval.ID}).Updates(&migrateModels.ProgramApproval{IsCommitteeSelected: true}).Error; err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return err
	}

	return
}
