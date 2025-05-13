package query

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeDetails struct {
	UID                 *uuid.UUID      `json:"uid"`
	Email               *string         `json:"email"`
	TitleTh             *string         `json:"title_th"`
	TitleEn             *string         `json:"title_en"`
	FirstnameTh         *string         `json:"firstname_th"`
	FirstnameEn         *string         `json:"firstname_en"`
	LastnameTh          *string         `json:"lastname_th"`
	LastnameEn          *string         `json:"lastname_en"`
	EducationBackGround *string         `json:"education_back_ground"`
	Position            *string         `json:"position"`
	UpdatedAt           *time.Time      `json:"updated_at"`
	CreatedAt           *time.Time      `json:"created_at"`
	DeletedAt           *gorm.DeletedAt `json:"deleted_at"`
}

func (EmployeeDetails) TableName() string {
	return "common_employee_details"
}
