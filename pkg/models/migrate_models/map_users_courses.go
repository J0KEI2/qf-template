package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MapUserCourse struct {
	User      Users       `gorm:"foreignKey:UserUID;references:UID"`
	Role      RoleCourses `gorm:"foreignKey:RoleUID;references:UID"`
	UserUID   uuid.UUID   `gorm:"column:user_uid;type:uuid;not null;primaryKey" json:"userUid"`
	CourseUID uuid.UUID   `gorm:"column:course_uid;type:uuid;not null;primaryKey" json:"courseUid"`
	RoleUID   uuid.UUID   `gorm:"column:role_uid;type:uuid;not null;primaryKey" json:"roleUid"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}
