package query

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

type Course struct {
	CourseID           uuid.UUID                `gorm:"column:course_id"`
	CourseNumber    int                      `gorm:"column:course_number"`
	Version         string                   `gorm:"column:version"`
	Faculty         query.FacultiesJoinQuery `gorm:"foreignKey:FacultyID;references:ID"`
	FacultyID       uint                     `gorm:"column:faculty_id"`
	DepartmentName  string                   `gorm:"column:department_name"`
	EducationYear   string                   `gorm:"column:education_year"`
	CourseInfo      CourseInfoQueryEntity `gorm:"foreignKey:CourseInfoID;references:ID"`
	CourseInfoID int                      `gorm:"column:course_info_id"`
	Lecturer        CourseLecturer              `gorm:"foreignKey:CourseLecturerID;references:ID"`
	CourseLecturerID   int                      `gorm:"column:course_lecturer_id"`
	Result          CourseResult                `gorm:"foreignKey:CourseResultID;references:ID"`
	CourseResultID     int                      `gorm:"column:course_result_id"`
	CourseManagementID int                      `gorm:"column:course_management_id"`
	Assessment      CourseAssessment            `gorm:"foreignKey:CourseAssessmentID;references:ID"`
	CourseAssessmentID int                      `gorm:"column:course_plan_and_assessment_id"`
	CourseReferenceID  int                      `gorm:"column:course_reference_id"`
	Status          enums.UserStatus         `gorm:"column:status"`
	CreatedAt       time.Time                `gorm:"column:created_at"`
	UpdatedAt       time.Time                `gorm:"column:updated_at"`
}

type CourseQueryEntity struct {
	CourseID                        *uuid.UUID                `gorm:"column:course_id"`
	CourseNumber                 *int                      `gorm:"column:course_number"`
	Version                      *string                   `gorm:"column:version"`
	Faculty                      *query.FacultiesJoinQuery `gorm:"foreignKey:FacultyID;references:ID"`
	FacultyID                    *uint                     `gorm:"column:faculty_id"`
	DepartmentName               *string                   `gorm:"column:department_name"`
	EducationYear                *string                   `gorm:"column:education_year"`
	CourseInfo                   *CourseInfoQueryEntity `gorm:"foreignKey:CourseInfoID;references:ID"`
	CourseInfoID              *int                      `gorm:"column:course_info_id"`
	Lecturer                     *CourseLecturer              `gorm:"foreignKey:CourseLecturerID;references:ID"`
	CourseLecturerID                *int                      `gorm:"column:course_lecturer_id"`
	Result                       *CourseResult                `gorm:"foreignKey:CourseResultID;references:ID"`
	CourseResultID                  *int                      `gorm:"column:course_result_id"`
	CourseTypeAndManagementID *int                      `gorm:"column:course_type_and_management_id"`
	Assessment                   *CourseAssessment            `gorm:"foreignKey:CourseAssessmentID;references:ID"`
	CourseAssessmentID              *int                      `gorm:"column:course_assessment_id"`
	CourseReferenceID               *int                      `gorm:"column:course_reference_id"`
	Status                       *string                   `gorm:"column:status"`
	CreatedAt                    *time.Time                `gorm:"column:created_at"`
	UpdatedAt                    *time.Time                `gorm:"column:updated_at"`
}

func (e *CourseQueryEntity) TableName() string {
	return "courses"
}
