package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileSystem struct {
	ID           uint           `json:"id"`
	QFType       string         `json:"qf_type"`
	QFMainID     uuid.UUID      `json:"qf_main_id"`
	CategoryType string         `json:"category_type"`
	Attribute    string         `json:"attribute"`
	FilePath     string         `json:"file_path"`
	FileName     string         `json:"file_name"`
	CreatedBy    uuid.UUID      `json:"created_by"`
	UpdatedBy    uuid.UUID      `json:"updated_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type MapFilesSystem struct {
	ID              uint           `json:"id"`
	FileID          uint           `json:"file_id"`
	ChecoID         uint           `json:"checo_id"`
	ApprovalID      uint           `json:"approval_id"`
	GeneralDetailID uint           `json:"general_detail_id"`
	ReferenceID     uint           `json:"reference_id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty"`
}
