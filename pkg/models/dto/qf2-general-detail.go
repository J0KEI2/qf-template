package dto

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgramGeneralDetailGetResponseDto struct {
	ID                    uint              `json:"id"`
	UniversityName        *string           `json:"university_name"`
	FacultyID             *uint             `json:"faculty_id"`
	ProgramNameTH         *string           `json:"program_name_th"`
	ProgramNameEN         *string           `json:"program_name_en"`
	ProgramCode           *string           `json:"program_code"`
	DegreeNameTH          *string           `json:"degree_name_th"`
	DegreeNameEN          *string           `json:"degree_name_en"`
	DegreeNameShortenTH   *string           `json:"degree_name_shorten_th"`
	DegreeNameShortenEN   *string           `json:"degree_name_shorten_en"`
	OverallCredit         *int              `json:"overall_credit"`
	ProgramMajorTypeID    *int              `json:"program_major_type_id"`
	ProgramMajorType      *string           `json:"program_major_type"`
	ProgramDegreeTypeID   *int              `json:"program_degree_type_id"`
	ProgramDegreeType     *string           `json:"program_degree_type"`
	NumberOfYear          *int              `json:"number_of_year"`
	ProgramLanguageID     *int              `json:"program_language_id"`
	ProgramLanguage       *string           `json:"program_language"`
	Admission             *string           `json:"admission"`
	MOU                   *string           `json:"mou"`
	MOUFilepath           *string           `json:"mou_filepath"`
	ProgramTypeID         *int              `json:"program_type_id"`
	ProgramType           *string           `json:"program_type"`
	ProgramYearID         *int              `json:"program_year_id"`
	ProgramYear           *uint             `json:"program_year"`
	Semester              *int              `json:"semester"`
	SemesterYear          *int              `json:"semester_year"`
	BoardApproval         *int              `json:"board_approval"`
	BoardApprovalDate     *time.Time        `json:"board_approval_date"`
	AcademicCouncil       *int              `json:"academic_council"`
	AcademicCouncilDate   *time.Time        `json:"academic_council_date"`
	UniversityCouncil     *int              `json:"university_council"`
	UniversityCouncilDate *time.Time        `json:"university_council_date"`
	IsSamePlanMajor       *bool             `json:"is_same_plan_major"`
	ProgramMajor          []ProgramMajorDto `json:"major"`
	IsNationalProgram     *bool             `json:"is_national_program"`
	IsEnglishProgram      *bool             `json:"is_english_program"`
	IsOther               *bool             `json:"is_other"`
	OtherName             *string           `json:"other_name"`
	ProgramAdjustFrom     *string           `json:"program_adjust_from"`
	ProgramAdjustYear     *int              `json:"program_adjust_year"`
}

type ProgramGeneralDetailRequestDto struct {
	ProgramMainID         uuid.UUID         `json:"program_main_id"`
	UniversityName        *string           `json:"university_name"`
	FacultyID             *uint             `json:"faculty_id"`
	ProgramNameTH         *string           `json:"program_name_th"`
	ProgramNameEN         *string           `json:"program_name_en"`
	ProgramCode           *string           `json:"program_code"`
	DegreeNameTH          *string           `json:"degree_name_th"`
	DegreeNameEN          *string           `json:"degree_name_en"`
	BranchNameTH          *string           `json:"branch_name_th"`
	BranchNameEN          *string           `json:"branch_name_en"`
	DegreeNameShortenTH   *string           `json:"degree_name_shorten_th"`
	DegreeNameShortenEN   *string           `json:"degree_name_shorten_en"`
	OverallCredit         *int              `json:"overall_credit"`
	ProgramMajorTypeID    *int              `json:"program_major_type_id"`
	ProgramMajorType      *string           `json:"program_major_type"`
	ProgramDegreeTypeID   *int              `json:"program_degree_type_id"`
	ProgramDegreeType     *string           `json:"program_degree_type"`
	NumberOfYear          *int              `json:"number_of_year"`
	ProgramLanguageID     *int              `json:"program_language_id"`
	ProgramLanguage       *string           `json:"program_language"`
	Admission             *string           `json:"admission"`
	MOU                   *string           `json:"mou"`
	MOUFilepath           *string           `json:"mou_filepath"`
	ProgramTypeID         *int              `json:"program_type_id"`
	ProgramType           *string           `json:"program_type"`
	ProgramYearID         *int              `json:"program_year_id"`
	ProgramYear           *uint             `json:"program_year"`
	Semester              *int              `json:"semester"`
	SemesterYear          *int              `json:"semester_year"`
	BoardApproval         *int              `json:"board_approval"`
	BoardApprovalDate     *time.Time        `json:"board_approval_date"`
	AcademicCouncil       *int              `json:"academic_council"`
	AcademicCouncilDate   *time.Time        `json:"academic_council_date"`
	UniversityCouncil     *int              `json:"university_council"`
	UniversityCouncilDate *time.Time        `json:"university_council_date"`
	IsSamePlanMajor       *bool             `json:"is_same_plan_major"`
	ProgramMajor          []ProgramMajorDto `json:"major"`
	IsNationalProgram     *bool             `json:"is_national_program"`
	IsEnglishProgram      *bool             `json:"is_english_program"`
	IsOther               *bool             `json:"is_other"`
	OtherName             *string           `json:"other_name"`
	ProgramAdjustFrom     *string           `json:"program_adjust_from"`
	ProgramAdjustYear     *int              `json:"program_adjust_year"`
	CreatedAt             time.Time         `json:"created_at"`
	UpdatedAt             time.Time         `json:"updated_at"`
	DeletedAt             gorm.DeletedAt    `json:"deleted_at,omitempty"`
}

type ProgramMajorDto struct {
	ID                     *uint                  `json:"id"`
	ProgramGeneralDetailID uint                   `json:"program_general_detail_id,omitempty"`
	Name                   *string                `json:"name"`
	ProgramPlanDetail      []ProgramPlanDetailDto `json:"plan_detail"`
}

type ProgramPlanDetailDto struct {
	ID                *uint               `json:"id"`
	ProgramMajorID    *uint               `json:"major_id,omitempty"`
	PlanName          *string             `json:"plan_name"`
	CreditRulesID     *int                `json:"credit_rules_id"`
	CreditRules       *string             `json:"credit_rules"`
	Credit            *uint16             `json:"credit"`
	IsSplitPlan       *bool               `json:"is_split_plan"`
	IsActive          *bool               `json:"is_active"`
	ProgramSubPlanDto []ProgramSubPlanDto `json:"sub_plan"`
}

type ProgramSubPlanDto struct {
	ID                *uint   `json:"id"`
	ProgramPlanDetail *uint   `json:"plan_detail_id,omitempty"`
	SubPlanName       *string `json:"sub_plan_name"`
	CreditRulesID     *int    `json:"credit_rules_id"`
	CreditRules       *string `json:"credit_rules"`
	Credit            *uint16 `json:"credit"`
}

type MouFileDto struct {
	FileID *uint
	File   *multipart.FileHeader
}

type GetMouFileResponseDto struct {
	Items []MouFileResponse `json:"items"`
}

type MouFileResponse struct {
	FileID   *uint   `json:"file_id"`
	FileName *string `json:"file_name"`
}
