package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"gorm.io/gorm"
)

type ProgramReferenceDto struct {
	ID                   *uint           `json:"id"`
	ProgramID            *uuid.UUID      `json:"program_id"`
	ReferenceName        *string         `json:"reference_name"`
	ReferenceDescription *string         `json:"reference_description"`
	ReferenceFileName    *string         `json:"reference_file_name"`
	ReferenceFilePath    *string         `json:"reference_file_path"`
	ReferenceTypeID      *int            `json:"reference_type_id"`
	ReferenceTypeName    *string         `json:"reference_type_name"`
	CreatedAt            *time.Time      `json:"created_at,omitempty"`
	UpdatedAt            *time.Time      `json:"updated_at,omitempty"`
	DeletedAt            *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type CreateOrUpdateReferenceDto struct {
	References []ProgramReferenceDto `json:"references"`
}

type ProgramReferenceResponseDto struct {
	ID                   *uint           `json:"id"`
	ProgramID            *uuid.UUID      `json:"program_id"`
	ReferenceName        *string         `json:"reference_name"`
	ReferenceDescription *string         `json:"reference_description"`
	ReferenceFileName    *string         `json:"reference_file_name"`
	ReferenceFileID      *uint           `json:"reference_file_id"`
	ReferenceTypeID      *int            `json:"reference_type_id"`
	ReferenceTypeName    *string         `json:"reference_type_name"`
	CreatedAt            *time.Time      `json:"created_at,omitempty"`
	UpdatedAt            *time.Time      `json:"updated_at,omitempty"`
	DeletedAt            *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type GetReferenceResponseDto struct {
	Items []ProgramReferenceResponseDto `json:"items"`
	*models.PaginationOptions
}
