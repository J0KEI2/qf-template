package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramPloFormatQueryEntity struct {
	ID               *uint                      `gorm:"primaryKey;autoIncrement"`
	ProgramSubPlanID *uint                      `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramSubPlan   *ProgramSubPlanQueryEntity `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	PLOFormat        *string                    `gorm:"column:plo_format;type:text"`
	ProgramPLO       []ProgramPloQueryEntity    `gorm:"foreignKey:ProgramPloFormatID;references:ID"`
	CreatedAt        *time.Time                 `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt        *time.Time                 `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt        *gorm.DeletedAt            `gorm:"index;column:deleted_at"`
}

func (q *ProgramPloFormatQueryEntity) TableName() string {
	return "program_plo_format"
}

type ProgramPloQueryEntity struct {
	ID                               *uint                                     `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramPloFormatID               *uint                                     `gorm:"column:program_plo_format_id;type:uint"`
	Order                            *uint                                     `gorm:"column:order;type:uint"`
	ParentID                         *uint                                     `gorm:"column:parent_id;type:uint;default:null"`
	Children                         []ProgramPloQueryEntity                   `gorm:"foreignKey:ParentID;references:ID"`
	PLOPrefix                        *string                                   `gorm:"column:plo_prefix;type:varchar;size:255"`
	PLODetail                        *string                                   `gorm:"column:plo_detail;type:varchar;size:255"`
	LearningSolution                 []ProgramPLOLearningSolutionQueryEntity   `gorm:"foreignKey:PloID;references:ID"`
	LearningEvaluation               []ProgramPLOLearningEvaluationQueryEntity `gorm:"foreignKey:PloID;references:ID"`
	ProgramMapPloWithKsecQueryEntity []ProgramMapPloWithKsecQueryEntity        `gorm:"foreignKey:PloID;references:ID"`
	CreatedAt                        *time.Time                                `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt                        *time.Time                                `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt                        *gorm.DeletedAt                           `gorm:"index;column:deleted_at"`
}

func (ProgramPloQueryEntity) TableName() string {
	return "program_plo"
}

type ProgramKsecDetailQueryEntity struct {
	ID                      *uint                              `gorm:"column:id;primaryKey;autoIncrement"`
	ProgramSubPlanID        *uint                              `gorm:"column:program_sub_plan_id;type:uint"`
	ProgramSubPlan          *ProgramSubPlanQueryEntity         `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	Type                    *string                            `gorm:"column:type;type:varchar;size:2;index"`
	Order                   *uint                              `gorm:"column:order;type:uint;index"`
	Detail                  *string                            `gorm:"column:detail;type:text"`
	ProgramMapPloWithKsec   []ProgramMapPloWithKsecQueryEntity `gorm:"foreignKey:KsecID;references:ID"`
	IsChecked               *bool                              `gorm:"-"`
	ProgramMapPloWithKsecID *uint                              `gorm:"-"`
	CreatedAt               *time.Time                         `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt               *time.Time                         `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt               *gorm.DeletedAt                    `gorm:"index;column:deleted_at"`
}

func (ProgramKsecDetailQueryEntity) TableName() string {
	return "program_ksec_detail"
}

type ProgramMapPloWithKsecQueryEntity struct {
	ID        *uint                         `gorm:"column:id;primaryKey;autoIncrement"`
	PloID     *uint                         `gorm:"column:plo_id;type:uint;not null"`
	KsecID    *uint                         `gorm:"column:ksec_id;type:uint;not null"`
	PLO       *ProgramPloQueryEntity        `gorm:"foreignKey:PloID;references:ID"`
	KSEC      *ProgramKsecDetailQueryEntity `gorm:"foreignKey:KsecID;references:ID"`
	CreatedAt *time.Time                    `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt *time.Time                    `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt *gorm.DeletedAt               `gorm:"index;column:deleted_at"`
}

func (ProgramMapPloWithKsecQueryEntity) TableName() string {
	return "program_map_plo_with_ksec"
}

type ProgramPLOLearningSolutionQueryEntity struct {
	ID        *uint           `gorm:"column:id;primaryKey;autoIncrement;index"`
	PloID     *uint           `gorm:"column:plo_id;type:uint;not null;index"`
	Key       *string         `gorm:"column:key;type:varchar;size:255;"`
	Detail    *string         `gorm:"column:detail;type:varchar;size:255;"`
	Order     *uint           `gorm:"column:order;type:uint"`
	CreatedAt *time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt *time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt *gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramPLOLearningSolutionQueryEntity) TableName() string {
	return "program_plo_learning_solution"
}

type ProgramPLOLearningEvaluationQueryEntity struct {
	ID        *uint           `gorm:"column:id;primaryKey;autoIncrement;index"`
	PloID     *uint           `gorm:"column:plo_id;type:uint;not null;index"`
	Key       *string         `gorm:"column:key;type:varchar;size:255;"`
	Detail    *string         `gorm:"column:detail;type:varchar;size:255;"`
	Order     *uint           `gorm:"column:order;type:uint"`
	CreatedAt *time.Time      `gorm:"column:created_at;default:current_timestamp"`
	UpdatedAt *time.Time      `gorm:"column:updated_at;default:current_timestamp;onUpdate:current_timestamp"`
	DeletedAt *gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (ProgramPLOLearningEvaluationQueryEntity) TableName() string {
	return "program_plo_learning_evaluation"
}
