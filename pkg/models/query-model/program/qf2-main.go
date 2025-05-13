package query

import (
	"time"

	"github.com/google/uuid"
	permissionQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"gorm.io/gorm"
)

type ProgramMainQueryEntity struct {
	ID                          *uuid.UUID                                     `gorm:"column:id"`
	ProgramGeneralDetail        *ProgramGeneralDetailQueryEntity               `gorm:"foreignKey:ProgramGeneralDetailID;references:ID"`
	ProgramGeneralDetailID      *uint                                          `gorm:"column:program_general_detail_id"`
	ProgramPolicyAndStrategic   *ProgramPolicyAndStrategicQueryEntity          `gorm:"foreignKey:ProgramPolicyAndStrategicID;references:ID"`
	ProgramPolicyAndStrategicID *uint                                          `gorm:"column:program_policy_and_strategic_id"`
	ProgramPlanAndEvaluate      *ProgramPlanAndEvaluateQueryEntity             `gorm:"foreignKey:ProgramPlanAndEvaluateID;references:ID"`
	ProgramPlanAndEvaluateID    *uint                                          `gorm:"column:program_plan_and_evaluate_id"`
	ProgramOwner                []ProgramOwnerQueryEntity                      `gorm:"foreignKey:ProgramMainUID;references:ID"`
	ProgramLecturer             []ProgramLecturerQueryEntity                   `gorm:"foreignKey:ProgramMainUID;references:ID"`
	ProgramThesisLecturer       []ProgramThesisLecturerQueryEntity             `gorm:"foreignKey:ProgramMainUID;references:ID"`
	ProgramQualityAssurance     *ProgramQualityAssuranceQueryEntity            `gorm:"foreignKey:ProgramQualityAssuranceID;references:ID"`
	ProgramQualityAssuranceID   *uint                                          `gorm:"column:program_quality_assurance_id"`
	ProgramSystemAndMechanic    *ProgramSystemAndMechanicQueryEntity           `gorm:"foreignKey:ProgramSystemAndMechanicID;references:ID"`
	ProgramSystemAndMechanicID  *uint                                          `gorm:"column:program_system_and_mechanic_id"`
	ProgramPermission           []permissionQuery.PermissionProgramQueryEntity `gorm:"foreignKey:ProgramUID;references:ID"`
	UpdatedAt                   *time.Time                                     `gorm:"column:updated_at"`
	CreatedAt                   *time.Time                                     `gorm:"column:created_at"`
	DeletedAt                   gorm.DeletedAt
}

// TableName specifies the table name for the Program struct
func (q *ProgramMainQueryEntity) TableName() string {
	return "program_main"
}
