package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MapUserProgram struct {
	User       Users        `gorm:"foreignKey:UserUID;references:UID"`
	Role       RolePrograms `gorm:"foreignKey:RoleUID;references:UID"`
	UserUID    uuid.UUID    `gorm:"column:user_uid;type:uuid;not null;primaryKey" json:"userUid"`
	ProgramUID uuid.UUID    `gorm:"column:program_uid;type:uuid;not null;primaryKey" json:"program_uid"`
	RoleUID    uuid.UUID    `gorm:"column:role_uid;type:uuid;not null;primaryKey" json:"roleUid"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}
