package query

import (
	"time"

	"github.com/google/uuid"
	commonQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	"gorm.io/gorm"
)

type ProgramReferenceQueryEntity struct {
	ID                   *uint                                  `gorm:"column:id"`
	ProgramID            *uuid.UUID                             `gorm:"column:program_id"`
	ReferenceName        *string                                `gorm:"column:reference_name"`
	ReferenceDescription *string                                `gorm:"column:reference_description"`
	ReferenceFileName    *string                                `gorm:"column:reference_file_name"`
	ReferenceFilePath    *string                                `gorm:"column:reference_file_path"`
	ReferenceTypeID      *int                                   `gorm:"column:reference_type_id"`
	ReferenceOption      commonQuery.ReferenceOptionQueryEntity `gorm:"foreignKey:ReferenceTypeID;references:ID"`
	CreatedAt            *time.Time                             `gorm:"column:created_at"`
	UpdatedAt            *time.Time                             `gorm:"column:updated_at"`
	DeletedAt            gorm.DeletedAt
}

func (ProgramReferenceQueryEntity) TableName() string {
	return "program_references"
}
