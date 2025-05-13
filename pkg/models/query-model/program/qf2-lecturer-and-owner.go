package query

import (
	"time"

	"github.com/google/uuid"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	"gorm.io/gorm"
)

type ProgramOwnerQueryEntity struct {
	ID             *uint                  `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramMainUID *uuid.UUID             `gorm:"column:program_main_uid;type:uuid"`
	OwnerID        *uuid.UUID             `gorm:"column:owner_id;type:uuid"`
	Owner          *query.EmployeeDetails `gorm:"foreignKey:UID;references:OwnerID"`
	CreatedAt      *time.Time             `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt      *time.Time             `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt      *gorm.DeletedAt        `gorm:"index;column:deleted_at"`
}

func (ProgramOwnerQueryEntity) TableName() string {
	return "program_owner"
}

type ProgramThesisLecturerQueryEntity struct {
	ID               *uint                  `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramMainUID   *uuid.UUID             `gorm:"column:program_main_uid;type:uuid"`
	ThesisLecturerID *uuid.UUID             `gorm:"column:thesis_lecturer_id;type:uuid"`
	ThesisLecturer   *query.EmployeeDetails `gorm:"foreignKey:UID;references:ThesisLecturerID"`
	CreatedAt        *time.Time             `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt        *time.Time             `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        *gorm.DeletedAt        `gorm:"index;column:deleted_at"`
}

func (ProgramThesisLecturerQueryEntity) TableName() string {
	return "program_thesis_lecturer"
}

type ProgramLecturerQueryEntity struct {
	ID             *uint                  `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramMainUID *uuid.UUID             `gorm:"column:program_main_uid;type:uuid"`
	LecturerID     *uuid.UUID             `gorm:"column:lecturer_id;type:uuid"`
	Lecturer       *query.EmployeeDetails `gorm:"foreignKey:UID;references:LecturerID"`
	CreatedAt      *time.Time             `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt      *time.Time             `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt      *gorm.DeletedAt        `gorm:"index;column:deleted_at"`
}

func (ProgramLecturerQueryEntity) TableName() string {
	return "program_lecturer"
}
