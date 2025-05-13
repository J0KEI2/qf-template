package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
)

type CourseFetchModel struct {
	CourseID                        uuid.UUID        `gorm:"column:course_id;type:uuid;primaryKey" json:"course_id"`
	CourseNumber                 int              `gorm:"column:course_number;type:integer;size:255;not null" json:"course_number"`
	Version                      string           `gorm:"column:version;type:varchar;size:255;not null" json:"version"`
	FacultyID                    uint             `gorm:"column:faculty_id;type:uint;not null" json:"faculty_id"`
	DepartmentName               string           `gorm:"column:department_name;type:varchar;size:100;not null" json:"department_name"`
	EducationYear                string           `gorm:"column:education_year;type:varchar;size:100;not null" json:"education_year"`
	CourseInfoID              int              `gorm:"column:course_info_id;type:integer;size:255;not null" json:"course_info_id"`
	CourseLecturerID                int              `gorm:"column:course_lecturer_id;type:integer;size:255;not null" json:"course_lecturer_id"`
	CourseResultID                  int              `gorm:"column:course_result_id;type:integer;size:255;not null" json:"course_result_id"`
	CourseTypeAndManagementID int              `gorm:"column:course_type_and_management_id;type:integer;size:255;not null" json:"course_type_and_management_id"`
	CourseAssessmentID              int              `gorm:"column:course_assessment_id;type:integer;size:255;not null" json:"course_assessment_id"`
	CourseReferenceID               int              `gorm:"column:course_reference_id;type:integer;size:255;not null" json:"course_reference_id"`
	Status                       enums.UserStatus `gorm:"column:status;type:user_status;not null" json:"status"`
	CreatedAt                    time.Time        `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;index" json:"created_at"`
	UpdatedAt                    time.Time        `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
