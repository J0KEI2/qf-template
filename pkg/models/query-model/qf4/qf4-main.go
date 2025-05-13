package query

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

type QF4Main struct {
	QF4ID           uuid.UUID                `gorm:"column:qf4_id"`
	CourseNumber    int                      `gorm:"column:course_number"`
	Version         string                   `gorm:"column:version"`
	Faculty         query.FacultiesJoinQuery `gorm:"foreignKey:FacultyID;references:ID"`
	FacultyID       uint                     `gorm:"column:faculty_id"`
	DepartmentName  string                   `gorm:"column:department_name"`
	EducationYear   string                   `gorm:"column:education_year"`
	CourseInfo      QF4CourseInfoQueryEntity `gorm:"foreignKey:QF4CourseInfoID;references:ID"`
	QF4CourseInfoID int                      `gorm:"column:qf4_course_info_id"`
	Lecturer        QF4Lecturer              `gorm:"foreignKey:QF4LecturerID;references:ID"`
	QF4LecturerID   int                      `gorm:"column:qf4_lecturer_id"`
	Result          QF4Result                `gorm:"foreignKey:QF4ResultID;references:ID"`
	QF4ResultID     int                      `gorm:"column:qf4_result_id"`
	QF4ManagementID int                      `gorm:"column:qf4_management_id"`
	Assessment      QF4Assessment            `gorm:"foreignKey:QF4AssessmentID;references:ID"`
	QF4AssessmentID int                      `gorm:"column:qf4_course_plan_and_assessment_id"`
	QF4ReferenceID  int                      `gorm:"column:qf4_reference_id"`
	Status          enums.UserStatus         `gorm:"column:status"`
	CreatedAt       time.Time                `gorm:"column:created_at"`
	UpdatedAt       time.Time                `gorm:"column:updated_at"`
}

type QF4MainQueryEntity struct {
	QF4ID                        *uuid.UUID                `gorm:"column:qf4_id"`
	CourseNumber                 *int                      `gorm:"column:course_number"`
	Version                      *string                   `gorm:"column:version"`
	Faculty                      *query.FacultiesJoinQuery `gorm:"foreignKey:FacultyID;references:ID"`
	FacultyID                    *uint                     `gorm:"column:faculty_id"`
	DepartmentName               *string                   `gorm:"column:department_name"`
	EducationYear                *string                   `gorm:"column:education_year"`
	CourseInfo                   *QF4CourseInfoJointQuery  `gorm:"foreignKey:QF4CourseInfoID;references:ID"`
	QF4CourseInfoID              *int                      `gorm:"column:qf4_course_info_id"`
	Lecturer                     *QF4Lecturer              `gorm:"foreignKey:QF4LecturerID;references:ID"`
	QF4LecturerID                *int                      `gorm:"column:qf4_lecturer_id"`
	Result                       *QF4ResultJointQuery      `gorm:"foreignKey:QF4ResultID;references:ID"`
	QF4ResultID                  *int                      `gorm:"column:qf4_result_id"`
	QF4CourseTypeAndManagementID *int                      `gorm:"column:qf4_course_type_and_management_id"`
	Assessment                   *QF4Assessment            `gorm:"foreignKey:QF4AssessmentID;references:ID"`
	QF4AssessmentID              *int                      `gorm:"column:qf4_assessment_id"`
	QF4ReferenceID               *int                      `gorm:"column:qf4_reference_id"`
	Status                       *enums.UserStatus         `gorm:"column:status"`
	CreatedAt                    *time.Time                `gorm:"column:created_at"`
	UpdatedAt                    *time.Time                `gorm:"column:updated_at"`
	DeletedAt                    *gorm.DeletedAt
}

func (s *QF4MainQueryEntity) TableName() string {
	return "qf4"
}
