package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgramApproval struct {
	ID                  uint           `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	ProgramUID          uuid.UUID      `gorm:"column:program_uid;type:uuid;"`
	NameTh              string         `gorm:"column:name_th;type:varchar;size:255"`
	NameEn              string         `gorm:"column:name_en;type:varchar;size:255"`
	ApprovalStatusLevel uint           `gorm:"column:approval_status_level;type:smallint"`
	IsCurrent           bool           `gorm:"column:is_current;type:boolean"`
	IsApproved          bool           `gorm:"column:is_approved;type:boolean"`
	IsRejected          bool           `gorm:"column:is_rejected;type:boolean"`
	IsCommitteeSelected bool           `gorm:"column:is_committee_selected;type:boolean"`
	CreatedAt           time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type ProgramApprovalSubmission struct {
	ID                uint            `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	ApprovalID        uint            `gorm:"column:approval_id;type:bigint"`
	SubmissionCount   uint            `gorm:"column:submission_count;type:bigint"`
	IsCurrent         bool            `gorm:"column:is_current;type:boolean"`
	IsApproved        bool            `gorm:"column:is_approved;type:boolean"`
	IsRejected        bool            `gorm:"column:is_rejected;type:boolean"`
	Comment           *string         `gorm:"column:comment;type:varchar"`
	CommitteeUserID   *uuid.UUID      `gorm:"column:committee_user_id;type:uuid"`
	UpdatedBy         *uuid.UUID      `gorm:"column:updated_by;type:uuid"`
	ApprovedDate      *time.Time      `gorm:"column:approved_date;type:timestamp"`
	MeetingNo         *string         `gorm:"column:meeting_no;type:varchar"`
	CreatedAt         time.Time       `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	CommitteeUserIDFK Users           `gorm:"foreignKey:CommitteeUserID;references:UID"`
	UpdatedByUserIDFK Users           `gorm:"foreignKey:UpdatedBy;references:UID"`
	ApprovalIDFK      ProgramApproval `gorm:"foreignKey:ApprovalID;references:ID"`
}

type CHECO struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	ProgramUID uuid.UUID `gorm:"column:program_uid;type:uuid;"`
	NameEN     string    `gorm:"column:name_en;type:varchar;size:255"`
	NameTH     string    `gorm:"column:name_th;type:varchar;size:255"`
	StatusID   uint      `gorm:"column:status_id;type:int"`
	// StatusLevel  uint       `gorm:"column:status_level;type:smallint"`
	ApprovedDate *time.Time `gorm:"column:approved_date;type:timestamp"`
	FileSystem   FileSystem `gorm:"foreignKey:FileID;references:ID"`
	FileID       *uint      `gorm:"column:file_id;type:bigint" json:"file_id"`
	CreatedAt    time.Time  `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}
