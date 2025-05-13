package approvals

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/domain"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

type approvalRepository struct {
	MainDbConn *gorm.DB
}

func NewProgramApprovalRepository(mainDbConn *gorm.DB) domain.ProgramApprovalRepository {
	return &approvalRepository{
		MainDbConn: mainDbConn,
	}
}

// args [approval_id]
const submissionQuery = `SELECT pas.id, pas.approval_id, pas.submission_count, pas.is_current, pas.is_approved, pas.is_rejected, pas."comment", pas.committee_user_id, u1.name_th as approver_name_th, u1.name_en as approver_name_en, u1.email as approver_email, pas.created_at, pas.updated_at, pas.approved_date, pas.meeting_no FROM program_approval_submissions pas LEFT join users u1 on u1.uid = pas.updated_by WHERE approval_id = ? order by pas.submission_count asc`

// args [approval_id]
const submissionQueryWithCommittees = `SELECT pas.id, pas.approval_id, pas.submission_count, pas.is_current, pas.is_approved, pas.is_rejected, pas."comment",
pas.committee_user_id,
u1.name_th as committee_name_th, u1.name_en as committee_name_en, u1.email as committee_email,
u2.name_th as approver_name_th, u2.name_en as approver_name_en, u2.email as approver_email,
pas.updated_by as approver_email, 
pas.created_at, pas.updated_at,
pas.approved_date, pas.meeting_no 
FROM program_approval_submissions pas 
LEFT join users u1 on u1.uid = pas.committee_user_id 
LEFT join users u2 on u2.uid = pas.updated_by 
WHERE approval_id = ? 
order by pas.submission_count asc`

const submissionQueryWithCommittee = `SELECT pas.id, pas.approval_id, pas.submission_count, pas.is_current, pas.is_approved, pas.is_rejected, pas."comment",
pas.committee_user_id,
u1.name_th as committee_name_th, u1.name_en as committee_name_en, u1.email as committee_email,
u2.name_th as approver_name_th, u2.name_en as approver_name_en, u2.email as approver_email,
pas.updated_by as approver_email, 
pas.created_at, pas.updated_at,
pas.approved_date, pas.meeting_no 
FROM program_approval_submissions pas 
LEFT join users u1 on u1.uid = pas.committee_user_id 
LEFT join users u2 on u2.uid = pas.updated_by 
WHERE approval_id = ? and committee_user_id = ? 
order by pas.submission_count asc`

// args [approval_id, program_uid, approval_status_level]
const getMaxSubmissionQuery = `select pas.id, pas.submission_count from program_approval_submissions pas where pas.submission_count = (select MAX(pas.submission_count) from program_approval_submissions pas where approval_id = (select id from program_approvals pa where pa.program_uid = ? and pa.approval_status_level = ?)) and approval_id = (select id from program_approvals pa where pa.program_uid = ? and pa.approval_status_level = ?)`

// args [program_uid, approval_status_level, committee_user_id]
const getMaxCommitteeSubmissionQuery = `select * from program_approval_submissions pas where pas.submission_count = (select MAX(pas.submission_count) 
from program_approval_submissions pas where approval_id = (select id from program_approvals pa 
where pa.program_uid = ? and pa.approval_status_level = ?) and pas.committee_user_id = ?) 
and pas.approval_id = (select id from program_approvals pa where pa.program_uid = ? and pa.approval_status_level = ?)
and pas.committee_user_id = ?`

// const getMaxSubmissionCountQuery = `select submission_count from program_approval_submissions pas where pas.submission_count = (select MAX(pas.submission_count)
// from program_approval_submissions pas where approval_id = (select id from program_approvals pa
// where pa.program_uid = ? and pa.approval_status_level = ?))
// and pas.approval_id = (select id from program_approvals pa where pa.program_uid = ? and pa.approval_status_level = ?)`

func (r *approvalRepository) DbApprovalVCMigrator() (err error) {

	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	if err := r.MainDbConn.AutoMigrate(
		&migrateModels.ProgramApproval{},
		&migrateModels.ProgramApprovalSubmission{},
		&migrateModels.FileSystem{},
		&migrateModels.CHECO{},
		&migrateModels.MapFilesSystem{},
	); err != nil {
		return err
	}

	return
}
