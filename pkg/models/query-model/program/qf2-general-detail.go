package query

import (
	"time"

	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	"gorm.io/gorm"
)

type ProgramGeneralDetailQueryEntity struct {
	ID                    *uint                     `gorm:"column:id"`
	UniversityName        *string                   `gorm:"column:university_name"`
	FacultyID             *uint                     `gorm:"column:faculty_id"`
	Faculty               *query.Faculty            `gorm:"foreignKey:ID;references:FacultyID"`
	ProgramNameTH         *string                   `gorm:"column:program_name_th"`
	ProgramNameEN         *string                   `gorm:"column:program_name_en"`
	ProgramCode           *string                   `gorm:"column:program_code"`
	BranchNameTH          *string                   `gorm:"column:branch_name_th"`
	BranchNameEN          *string                   `gorm:"column:branch_name_en"`
	DegreeNameTH          *string                   `gorm:"column:degree_name_th"`
	DegreeNameEN          *string                   `gorm:"column:degree_name_en"`
	DegreeNameShortenTH   *string                   `gorm:"column:degree_name_shorten_th"`
	DegreeNameShortenEN   *string                   `gorm:"column:degree_name_shorten_en"`
	OverallCredit         *int                      `gorm:"column:overall_credit"`
	ProgramMajorTypeID    *int                      `gorm:"column:program_major_type_id"`
	ProgramMajorType      *string                   `gorm:"column:program_major_type"`
	ProgramDegreeTypeID   *int                      `gorm:"column:program_degree_type_id"`
	ProgramDegreeType     *string                   `gorm:"column:program_degree_type"`
	NumberOfYear          *int                      `gorm:"column:number_of_year"`
	ProgramLanguageID     *int                      `gorm:"column:program_language_id"`
	ProgramLanguage       *string                   `gorm:"column:program_language"`
	Admission             *string                   `gorm:"column:admission"`
	MOU                   *string                   `gorm:"column:mou"`
	MOUFilepath           *string                   `gorm:"column:mou_filepath"`
	ProgramTypeID         *int                      `gorm:"column:program_type_id"`
	ProgramType           *string                   `gorm:"column:program_type"`
	ProgramYearID         *int                      `gorm:"column:program_year_id"`
	ProgramYear           *uint                     `gorm:"column:program_year"`
	Semester              *int                      `gorm:"column:semester"`
	SemesterYear          *int                      `gorm:"column:semester_year"`
	BoardApproval         *int                      `gorm:"column:board_approval"`
	BoardApprovalDate     *time.Time                `gorm:"column:board_approval_date"`
	AcademicCouncil       *int                      `gorm:"column:academic_council"`
	AcademicCouncilDate   *time.Time                `gorm:"column:academic_council_date"`
	UniversityCouncil     *int                      `gorm:"column:university_council"`
	UniversityCouncilDate *time.Time                `gorm:"column:university_council_date"`
	IsSamePlanMajor       *bool                     `gorm:"column:is_same_plan_major"`
	IsNationalProgram     *bool                     `gorm:"column:is_national_program"`
	IsEnglishProgram      *bool                     `gorm:"column:is_english_program"`
	IsOther               *bool                     `gorm:"column:is_other"`
	OtherName             *string                   `gorm:"column:other_name"`
	ProgramAdjustFrom     *string                   `gorm:"column:program_adjust_from"`
	ProgramAdjustYear     *int                      `gorm:"column:program_adjust_year"`
	ProgramMajor          []ProgramMajorQueryEntity `gorm:"foreignKey:ProgramGeneralDetailID;references:ID"`
	UpdatedAt             *time.Time                `gorm:"column:updated_at"`
	CreatedAt             *time.Time                `gorm:"column:created_at"`
	DeletedAt             gorm.DeletedAt
}

func (q *ProgramGeneralDetailQueryEntity) TableName() string {
	return "program_general_detail"
}

type ProgramMajorQueryEntity struct {
	ID                     *uint                          `gorm:"column:id"`
	ProgramGeneralDetailID *uint                          `gorm:"column:program_general_detail_id"`
	Name                   *string                        `gorm:"column:name"`
	ProgramPlanDetail      []ProgramPlanDetailQueryEntity `gorm:"foreignKey:ProgramMajorID;references:ID"`
	CreatedAt              *time.Time                     `gorm:"column:created_at"`
	UpdatedAt              *time.Time                     `gorm:"column:updated_at"`
	DeletedAt              gorm.DeletedAt
}

// TableName sets the table name for the MajorAndPlan model
func (q *ProgramMajorQueryEntity) TableName() string {
	return "program_major"
}

type ProgramPlanDetailQueryEntity struct {
	ID             *uint                       `gorm:"column:id"`
	ProgramMajorID *uint                       `gorm:"column:program_major_id"`
	PlanName       *string                     `gorm:"column:plan_name"`
	CreditRulesID  *int                        `gorm:"column:credit_rules_id"`
	CreditRules    *string                     `gorm:"column:credit_rules"`
	Credit         *uint16                     `gorm:"column:credit"`
	IsSplitPlan    *bool                       `gorm:"column:is_split_plan"`
	IsActive       *bool                       `gorm:"column:is_active"`
	ProgramSubPlan []ProgramSubPlanQueryEntity `gorm:"foreignKey:ProgramPlanDetailID;references:ID"`
	CreatedAt      *time.Time                  `gorm:"column:created_at"`
	UpdatedAt      *time.Time                  `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt
}

// TableName sets the table name for the Plan model
func (q *ProgramPlanDetailQueryEntity) TableName() string {
	return "program_plan_detail"
}

type ProgramSubPlanQueryEntity struct {
	ID                  *uint                               `gorm:"column:id"`
	ProgramPlanDetailID *uint                               `gorm:"column:program_plan_detail_id"`
	SubPlanName         *string                             `gorm:"column:sub_plan_name"`
	CreditRulesID       *int                                `gorm:"column:credit_rules_id"`
	CreditRules         *string                             `gorm:"column:credit_rules"`
	Credit              *uint16                             `gorm:"column:credit"`
	YearAndSemester     []ProgramYearAndSemesterQueryEntity `gorm:"foreignKey:ProgramSubPlanID;references:ID"`
	CreatedAt           *time.Time                          `gorm:"column:created_at"`
	UpdatedAt           *time.Time                          `gorm:"column:updated_at"`
	DeletedAt           gorm.DeletedAt
}

// TableName sets the table name for the Plan model
func (q *ProgramSubPlanQueryEntity) TableName() string {
	return "program_sub_plan"
}
