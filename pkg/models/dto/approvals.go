package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetAllApprovalDto struct {
	ID            uint      `json:"id"`
	ProgramUID    uuid.UUID `json:"program_uid"`
	NameTh        string    `json:"name_th"`
	NameEn        string    `json:"name_en"`
	ApprovalLevel uint      `json:"approval_level"`
	IsCurrent     bool      `json:"is_current"`
	IsApproved    bool      `json:"is_approved"`
	IsRejected    bool      `json:"is_rejected"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ApprovalSubmissionDto struct {
	ID              uint       `json:"id"`
	SubmissionCount uint       `json:"submission_count"`
	IsCurrent       bool       `json:"is_current"`
	IsApproved      bool       `json:"is_approved"`
	IsRejected      bool       `json:"is_rejected"`
	Comment         string     `json:"comment"`
	ApproverNameTh  string     `json:"approver_name_th,omitempty"` // updated_by
	ApproverNameEn  string     `json:"approver_name_en,omitempty"` // updated_by
	ApproverEmail   string     `json:"approver_email,omitempty"`   // updated_by
	ApprovedDate    *time.Time `json:"approved_date"`
	MeetingNo       string     `json:"meeting_no,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type ApprovalSubmissionWithCommitteeDto struct {
	ID              uint   `json:"id"`
	SubmissionCount uint   `json:"submission_count"`
	IsCurrent       bool   `json:"is_current"`
	IsApproved      bool   `json:"is_approved"`
	IsRejected      bool   `json:"is_rejected"`
	Comment         string `json:"comment"`
	ApproverNameTh  string `json:"approver_name_th,omitempty"`  // updated_by
	ApproverNameEn  string `json:"approver_name_en,omitempty"`  // updated_by
	ApproverEmail   string `json:"approver_email,omitempty"`    // updated_by
	CommitteeNameTh string `json:"committee_name_th,omitempty"` // committee
	CommitteeNameEn string `json:"committee_name_en,omitempty"` // committee
	CommitteeEmail  string `json:"committee_email,omitempty"`   // committee
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type GetApprovalFacultyDto struct {
	ID                  uint                    `json:"id"`
	NameTh              string                  `json:"name_th"`
	NameEn              string                  `json:"name_en"`
	ApprovalStatusLevel uint                    `json:"approval_level"`
	ProgramUID          string                  `json:"program_uid"`
	IsCurrent           bool                    `json:"is_current"`
	IsApproved          bool                    `json:"is_approved"`
	IsRejected          bool                    `json:"is_rejected"`
	IsEditable          bool                    `json:"is_editable"`
	Submissions         []ApprovalSubmissionDto `json:"submissions"`
	ApprovedDate        *time.Time              `json:"approved_date,omitempty"`
	MeetingNo           *string                 `json:"meeting_no,omitempty"`
	CreatedAt           time.Time               `json:"created_at"`
	UpdatedAt           time.Time               `json:"updated_at"`
}

type SelectCurriculumCommittees struct {
	CommitteeUserIDs []string `json:"committee_user_ids"`
	UpdatedBy        string   `json:"updated_by"`
}

type GetApprovalCurriculumCommitteeDto struct {
	ID                          uint                       `json:"id"`
	ApprovalStatusLevel         uint                       `json:"approval_level"`
	NameTh                      string                     `json:"name_th"`
	NameEn                      string                     `json:"name_en"`
	ProgramUID                  string                     `json:"program_uid"`
	IsCurrent                   bool                       `json:"is_current"`
	IsApproved                  bool                       `json:"is_approved"`
	IsRejected                  bool                       `json:"is_rejected"`
	IsEditable                  bool                       `json:"is_editable"`
	IsCommitteeSelected         bool                       `json:"is_committee_selected"`
	Committees                  []GetApprovalCommitteesDto `json:"committees"`
	CurriculumCommitteeFileList []ApprovalFileList         `json:"file_list"`
	CreatedAt                   time.Time                  `json:"created_at"`
	UpdatedAt                   time.Time                  `json:"updated_at"`
}

type GetApprovalCommitteesDto struct {
	UserID           uuid.UUID               `json:"user_id"`
	NameTh           string                  `json:"name_th"`
	NameEn           string                  `json:"name_en"`
	Email            string                  `json:"email"`
	IsCurrent        bool                    `json:"is_current"`  // same with latest result
	IsApproved       bool                    `json:"is_approved"` // same with latest result
	IsRejected       bool                    `json:"is_rejected"` // same with latest result
	CommitteeResults []ApprovalSubmissionDto `json:"committee_results"`
	ApprovedDate     time.Time               `json:"approved_date"`
	MeetingNo        string                  `json:"meeting_no"`
	CreatedAt        time.Time               `json:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at"`
}

type CommitteeResultDto struct {
	ID              uint      `json:"id"`
	ApprovalID      uint      `json:"approval_id"`
	SubmissionCount uint      `json:"submission_count"`
	IsCurrent       bool      `json:"is_current"`
	IsApproved      bool      `json:"is_approved"`
	IsRejected      bool      `json:"is_rejected"`
	Comment         *string   `json:"comment"`
	CommitteeUserID uuid.UUID `json:"committee_user_id"`
	CommitteeNameTh *string   `json:"committee_name_th"`
	CommitteeNameEn *string   `json:"committee_name_en"`
	CommitteeEmail  *string   `json:"committee_email"`
	ApproverUserID  uuid.UUID `json:"approver_user_id"`
	ApproverNameTh  *string   `json:"approver_name_th"`
	ApproverNameEn  *string   `json:"approver_name_en"`
	ApproverEmail   *string   `json:"approver_email"`
	ApprovedDate    time.Time `json:"approved_date"`
	MeetingNo       string    `json:"meeting_no"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ApproveApprovalCommitteeResultRequestDto struct {
	Comment         string `json:"comment"`
	SubmissionID    uint   `json:"submission_id"`
	UpdatedBy       string `json:"updated_by"`
	CommitteeUserID string `json:"committee_user_id"`
}

type RejectApprovalCommitteeResultRequestDto struct {
	Comment         string `json:"comment"`
	SubmissionID    uint   `json:"submission_id"`
	UpdatedBy       string `json:"updated_by"`
	CommitteeUserID string `json:"committee_user_id"`
}

type GetApprovalAcademicDto struct {
	ID                  uint                    `json:"id"`
	NameTh              string                  `json:"name_th"`
	NameEn              string                  `json:"name_en"`
	ApprovalStatusLevel uint                    `json:"approval_level"`
	ProgramUID          string                  `json:"program_uid"`
	IsCurrent           bool                    `json:"is_current"`
	IsApproved          bool                    `json:"is_approved"`
	IsRejected          bool                    `json:"is_rejected"`
	IsEditable          bool                    `json:"is_editable"`
	Submissions         []ApprovalSubmissionDto `json:"submissions"`
	ApprovedDate        *time.Time              `json:"approved_date,omitempty"`
	MeetingNo           *string                 `json:"meeting_no,omitempty"`
	AcademicFileList    []ApprovalFileList      `json:"file_list"`
	CreatedAt           time.Time               `json:"created_at"`
	UpdatedAt           time.Time               `json:"updated_at"`
}
type GetApprovalUniversityDto struct {
	ID                  uint                    `json:"id"`
	NameTh              string                  `json:"name_th"`
	NameEn              string                  `json:"name_en"`
	ApprovalStatusLevel uint                    `json:"approval_level"`
	ProgramUID          string                  `json:"program_uid"`
	IsCurrent           bool                    `json:"is_current"`
	IsApproved          bool                    `json:"is_approved"`
	IsRejected          bool                    `json:"is_rejected"`
	IsEditable          bool                    `json:"is_editable"`
	Submissions         []ApprovalSubmissionDto `json:"submissions"`
	ApprovedDate        *time.Time              `json:"approved_date,omitempty"`
	MeetingNo           *string                 `json:"meeting_no,omitempty"`
	UniversityFileList  []ApprovalFileList      `json:"file_list"`
	CreatedAt           time.Time               `json:"created_at"`
	UpdatedAt           time.Time               `json:"updated_at"`
}

type ApproveApprovalRequestDto struct {
	Comment         string     `json:"comment"`
	SubmissionID    uint       `json:"submission_id"`
	ApprovedDate    *time.Time `json:"approved_date"`
	MeetingNo       *string    `json:"meeting_no"`
	UpdatedBy       string     `json:"updated_by"`
	CommitteeUserID string     `json:"committee_user_id"`
}

type RejectApprovalRequestDto struct {
	Comment      string `json:"comment"`
	SubmissionID uint   `json:"submission_id"`
	// ApprovedDate    *time.Time `json:"approved_date"`
	// MeetingNo       *string    `json:"meeting_no"`
	UpdatedBy       string `json:"updated_by"`
	CommitteeUserID string `json:"committee_user_id"`
}

type ApprovalFileList struct {
	FileID   *uint   `json:"file_id"`
	FileName *string `json:"file_name"`
}

type CreateCHECOStatusRequestDto struct {
	ProgramUID   *uuid.UUID `json:"program_uid"`
	StatusID     *int       `json:"status_id"`
	ApprovedDate *time.Time `json:"approved_date"`
}

type GetCHECOByUIDResponseDto struct {
	ProgramUID uuid.UUID          `json:"program_uid"`
	Items      []CHECODetailsList `json:"items"`
}

type CHECODetailsList struct {
	ID            uint            `json:"id"`
	NameEn        string          `json:"name_en"`
	NameTN        string          `json:"name_th"`
	Status        string          `json:"status"`
	ApprovedDate  time.Time       `json:"approved_date"`
	CHECOFileList []CHECOFileList `json:"checo_file_list"`
	CreatedAt     time.Time       `json:"created_at"`
}

type GetCHECORequestDto struct {
	ProgramUID *uuid.UUID `json:"program_uid"`
}

type CHECOFileList struct {
	FileID   *uint   `json:"file_id"`
	FileName *string `json:"file_name"`
}
