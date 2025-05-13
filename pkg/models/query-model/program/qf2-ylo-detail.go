package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramYLODetailQueryEntity struct {
	ID               *uint
	ProgramSubPlanID *uint
	ProgramSubPlan   *ProgramSubPlanQueryEntity `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	Order            *uint8
	ParentID         *uint                        `gorm:"column:parent_id"`
	Parent           *ProgramYLODetailQueryEntity `gorm:"foreignKey:ParentID;references:ID"`
	YLODetail        *string
	YLOYear          *uint
	IsActive         *bool
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	DeletedAt        *gorm.DeletedAt
}

// TableName sets the table name for the YLO model
func (q *ProgramYLODetailQueryEntity) TableName() string {
	return "program_ylo_detail"
}

type ProgramYloKsecQueryEntity struct {
	ID                       *uint                              `gorm:"primaryKey;autoIncrement" json:"id"`
	ProgramSubPlanID         *uint                              `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramSubPlan           *ProgramSubPlanQueryEntity         `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramYearAndSemesterID *uint                              `gorm:"column:program_year_and_semester_id;type:uint"`
	ProgramYearAndSemester   *ProgramYearAndSemesterQueryEntity `gorm:"foreignKey:ProgramYearAndSemesterID;references:ID"`
	Knowledge                *string                            `gorm:"column:knowledge;type:varchar;size:255;"`
	Skill                    *string                            `gorm:"column:skill;type:varchar;size:255;"`
	Ethic                    *string                            `gorm:"column:ethic;type:varchar;size:255;"`
	Character                *string                            `gorm:"column:character;type:varchar;size:255;"`
	CreatedAt                *time.Time                         `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                *time.Time                         `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                *gorm.DeletedAt                    `gorm:"index;column:deleted_at"`
}

// TableName sets the table name for the YLO model
func (ProgramYloKsecQueryEntity) TableName() string {
	return "program_ylo_ksec"
}

type ProgramYloWithKsecQueryEntity struct {
	ID                       *uint                              `gorm:"primaryKey;autoIncrement" json:"id"`
	ProgramYearAndSemesterID *uint                              `gorm:"column:program_year_and_semester_id;type:uint"`
	ProgramYearAndSemester   *ProgramYearAndSemesterQueryEntity `gorm:"foreignKey:ProgramYearAndSemesterID;references:ID"`
	ProgramMapPloWithKsecID  *uint                              `gorm:"column:program_map_plo_with_ksec_id;type:varchar;size:255;"`
	ProgramMapPloWithKsec    *ProgramMapPloWithKsecQueryEntity  `gorm:"foreignKey:ProgramMapPloWithKsecID;references:ID"`
	Remark                   *string                            `gorm:"column:remark;type:varchar;size:255;"`
	IsChecked                *bool                              `gorm:"column:is_checked" json:"is_checked"`
	CreatedAt                *time.Time                         `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                *time.Time                         `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                *gorm.DeletedAt                    `gorm:"index;column:deleted_at"`
}

func (ProgramYloWithKsecQueryEntity) TableName() string {
	return "program_ylo_with_ksec"
}

type ProgramYloWithPloQueryEntity struct {
	ID                       *uint                              `gorm:"primaryKey;autoIncrement" json:"id"`
	ProgramYearAndSemesterID *uint                              `gorm:"column:program_year_and_semester_id;type:uint"`
	ProgramYearAndSemester   *ProgramYearAndSemesterQueryEntity `gorm:"foreignKey:ProgramYearAndSemesterID;references:ID"`
	ProgramPloID             *uint                              `gorm:"column:program_plo_id;type:varchar;size:255;"`
	ProgramPlo               *ProgramPloQueryEntity             `gorm:"foreignKey:ProgramPloID;references:ID"`
	Remark                   *string                            `gorm:"column:remark;type:varchar;size:255;"`
	IsChecked                *bool                              `gorm:"column:is_checked" json:"is_checked"`
	CreatedAt                *time.Time                         `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                *time.Time                         `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                *gorm.DeletedAt                    `gorm:"index;column:deleted_at"`
}

func (ProgramYloWithPloQueryEntity) TableName() string {
	return "program_ylo_with_plo"
}
