package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramMapCurMapRespQueryEntity struct {
	ID                *uint                       `gorm:"column:id"`
	ProgramSubPlan        *ProgramSubPlanQueryEntity      `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID      *uint                       `gorm:"column:program_sub_plan_id"`
	ProgramCourseDetailID *uint                       `gorm:"column:program_course_detail_id"`
	ProgramCourseDetail   *ProgramCourseDetailQueryEntity `gorm:"foreignKey:ProgramCourseDetailID;references:ID"` //has child ref
	ProgramPloID          *uint                       `gorm:"column:program_plo_id"`
	ProgramPlo            *ProgramPloQueryEntity          `gorm:"foreignKey:ProgramPloID;references:ID"`
	Status            *int                        `gorm:"column:status"`
	CreatedAt         *time.Time                  `gorm:"column:created_at"`
	UpdatedAt         *time.Time                  `gorm:"column:updated_at"`
	DeletedAt         *gorm.DeletedAt             `gorm:"index;column:deleted_at"`
}

func (ProgramMapCurMapRespQueryEntity) TableName() string {
	return "program_map_curmap_resp"
}

type ProgramMapCurMapKsaQueryEntity struct {
	ID                *uint                       `gorm:"column:id"`
	ProgramSubPlan        *ProgramSubPlanQueryEntity      `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID      *uint                       `gorm:"column:program_sub_plan_id"`
	ProgramCourseDetailID *uint                       `gorm:"column:program_course_detail_id"`
	ProgramCourseDetail   *ProgramCourseDetailQueryEntity `gorm:"foreignKey:ProgramCourseDetailID;references:ID"` //has child ref
	ProgramPloID          *uint                       `gorm:"column:program_plo_id"`
	ProgramPlo            *ProgramPloQueryEntity          `gorm:"foreignKey:ProgramPloID;references:ID"`
	KsaID             *string                     `gorm:"column:ksa_id"`
	CreatedAt         *time.Time                  `gorm:"column:created_at"`
	UpdatedAt         *time.Time                  `gorm:"column:updated_at"`
	DeletedAt         *gorm.DeletedAt             `gorm:"index;column:deleted_at"`
}

func (ProgramMapCurMapKsaQueryEntity) TableName() string {
	return "program_map_curmap_ksa"
}

type ProgramKsaDetailQueryEntity struct {
	ID           *uint                  `gorm:"column:id"`
	ProgramSubPlan   *ProgramSubPlanQueryEntity `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID *uint                  `gorm:"column:program_sub_plan_id"`
	KsaType      *string                `gorm:"column:ksa_type"`
	Order        *uint                  `gorm:"column:order"`
	ShortCode    *string                `gorm:"column:short_code"`
	KsaDetail    *string                `gorm:"column:ksa_detail"`
	CreatedAt    *time.Time             `gorm:"column:created_at"`
	UpdatedAt    *time.Time             `gorm:"column:updated_at"`
	DeletedAt    *gorm.DeletedAt        `gorm:"index;column:deleted_at"`
}

func (ProgramKsaDetailQueryEntity) TableName() string {
	return "program_ksa_detail"
}
