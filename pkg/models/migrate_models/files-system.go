package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileSystem struct {
	ID           uint           `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	QFType       string         `gorm:"column:qf_type;type:varchar(255)" json:"qf_type"`
	QFMainID     uuid.UUID      `gorm:"column:qf_main_id;type:uuid" json:"qf_main_id"`
	CategoryType string         `gorm:"column:category_type;type:varchar(255)" json:"category_type"`
	Attribute    string         `gorm:"column:attribute;type:varchar(255)" json:"attribute"`
	FilePath     string         `gorm:"column:file_path;type:text" json:"file_path"`
	FileName     string         `gorm:"column:file_name;type:varchar(255)" json:"file_name"`
	CreatedBy    uuid.UUID      `gorm:"column:created_by;type:uuid" json:"created_by"`
	UpdatedBy    uuid.UUID      `gorm:"column:updated_by;type:uuid" json:"updated_by"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at,omitempty"`
}

type MapFilesSystem struct {
	ID              uint           `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	FileID          uint           `gorm:"column:file_id;type:bigint" json:"file_id"`
	ChecoID         uint           `gorm:"column:checo_id;type:bigint" json:"checo_id"`
	ApprovalID      uint           `gorm:"column:approval_id;type:bigint" json:"approval_id"`
	GeneralDetailID uint           `gorm:"column:general_detail;type:bigint" json:"general_detail_id"`
	ReferenceID     uint           `gorm:"column:reference_id;type:bigint" json:"reference_id"`
	ReportID        uint           `gorm:"column:report_id;type:bigint" json:"report_id"`
	CreatedAt       time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at,omitempty"`

	FileSystem FileSystem       `gorm:"foreignKey:FileID;references:ID"`
	Checo      CHECO            `gorm:"foreignKey:ChecoID;references:ID"`
	Approval   ProgramApproval  `gorm:"foreignKey:ApprovalID;references:ID"`
	Reference  ProgramReference `gorm:"foreignKey:ReferenceID;references:ID"`
}
