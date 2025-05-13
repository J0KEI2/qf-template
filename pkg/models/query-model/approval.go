package query

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgramApprovalQueryEntity struct {
	ID                  uint
	ProgramUID          uuid.UUID
	NameTh              string
	NameEn              string
	ApprovalStatusLevel uint
	IsCurrent           bool
	IsApproved          bool
	IsRejected          bool
	IsCommitteeSelected bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

func (e *ProgramApprovalQueryEntity) TableName() string {
	return "program_approvals"
}

type ProgramApprovalSubmissionQueryEntity struct {
	ID              uint       `json:"id"`
	ApprovalID      uint       `json:"approval_id"`
	SubmissionCount uint       `json:"submission_count"`
	IsCurrent       bool       `json:"is_current"`
	IsApproved      bool       `json:"is_approved"`
	IsRejected      bool       `json:"is_rejected"`
	Comment         *string    `json:"comment"`
	CommitteeUserID *uuid.UUID `json:"committee_user_id"`
	ApprovalDate    *time.Time `json:"approval_date"`
	UpdatedBy       *uuid.UUID `json:"updated_by"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`

	CommitteeUserIDFK UserQueryEntity            `gorm:"foreignKey:CommitteeUserID;references:UID"`
	UpdatedByUserIDFK UserQueryEntity            `gorm:"foreignKey:UpdatedBy;references:UID"`
	SubmissionFK      ProgramApprovalQueryEntity `gorm:"foreignKey:ApprovalID;references:ID"`
}

type ProgramApprovalSubmissionQueryEntityJoinUser struct {
	ID                  uint
	ApprovalID          uint
	SubmissionCount     uint
	IsCurrent           bool
	IsApproved          bool
	IsRejected          bool
	Comment             *string
	CommitteeUserID     *uuid.UUID
	CommitteeUserNameTh *string
	CommitteeUserNameEn *string
	CommitteeUserEmail  *string
	ApprovalDate        *time.Time
	UpdatedBy           *uuid.UUID
	ApproverNameTh      *string
	ApproverNameEn      *string
	ApproverEmail       *string
	CreatedAt           time.Time
	UpdatedAt           time.Time

	CommitteeUserIDFK UserQueryEntity            `gorm:"foreignKey:CommitteeUserID;references:UID"`
	UpdatedByUserIDFK UserQueryEntity            `gorm:"foreignKey:UpdatedBy;references:UID"`
	SubmissionFK      ProgramApprovalQueryEntity `gorm:"foreignKey:ApprovalID;references:ID"`
}

func (e *ProgramApprovalSubmissionQueryEntity) TableName() string {
	return "program_approval_submissions"
}

type CHECOQueryEntity struct {
	ID           *uint
	ProgramUID   *uuid.UUID
	NameEN       *string
	NameTH       *string
	StatusID     *int
	ApprovedDate *time.Time
	FileID       *uint
	FileSystem   FileSystemQueryEntity `gorm:"foreignKey:FileID;references:ID"`
	CreatedAt    *time.Time
}

func (e *CHECOQueryEntity) TableName() string {
	return "checos"
}
