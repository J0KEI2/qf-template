package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramStructureDetailQueryEntity struct {
	ID               *uint                               `gorm:"column:id;primaryKey;autoIncrement;type:uint"`
	ProgramSubPlan   *ProgramSubPlanQueryEntity          `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	ProgramSubPlanID *uint                               `gorm:"column:program_sub_plan_id"`
	Name             *string                             `gorm:"column:name;type:varchar;size:255"`
	Order            *uint                               `gorm:"column:order;type:uint"`
	ParentID         *uint                               `gorm:"column:parent_id;type:uint;default:null"`
	Children         []ProgramStructureDetailQueryEntity `gorm:"foreignKey:ParentID;references:ID"`
	CourseDetail     []ProgramCourseDetailQueryEntity    `gorm:"foreignKey:ProgramStructureID;references:ID"`
	Qualification    *string                             `gorm:"column:qualification;type:varchar;size:255"`
	StructureCredit  *uint                               `gorm:"column:structure_credit;type:uint"`
	CreatedAt        *time.Time                          `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt        *time.Time                          `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        *gorm.DeletedAt                     `gorm:"column:deleted_at;index"`
}

func (ProgramStructureDetailQueryEntity) TableName() string {
	return "program_structure_detail"
}
