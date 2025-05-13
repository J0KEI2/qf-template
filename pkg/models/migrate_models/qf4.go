package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
	"gorm.io/gorm"
)

type QF4 struct {
	QF4ID                        uuid.UUID                  `gorm:"column:qf4_id;type:uuid;primaryKey" json:"qf4_id"`
	CourseNumber                 int                        `gorm:"column:course_number;type:integer;size:255;" json:"course_number"`
	Version                      string                     `gorm:"column:version;type:varchar;size:100;not null" json:"version"`
	Faculty                      Faculty                    `gorm:"foreignKey:FacultyID;references:ID" json:"faculty"`
	FacultyID                    uint                       `gorm:"column:faculty_id;type:uint;not null" json:"faculty_id"`
	DepartmentName               string                     `gorm:"column:department_name;type:varchar;size:100;not null" json:"department_name"`
	EducationYear                string                     `gorm:"column:education_year;type:varchar;size:100;not null" json:"education_year"`
	CourseInfo                   QF4CourseInfo              `gorm:"foreignKey:QF4CourseInfoID;references:ID" json:"course_info"`
	QF4CourseInfoID              int                        `gorm:"column:qf4_course_info_id;type:integer;size:255;" json:"qf4_course_info_id"`
	Lecturer                     QF4Lecturer                `gorm:"foreignKey:QF4LecturerID;references:ID" json:"lecturer"`
	QF4LecturerID                int                        `gorm:"column:qf4_lecturer_id;type:integer;size:255;" json:"qf4_lecturer_id"`
	Result                       QF4Result                  `gorm:"foreignKey:QF4ResultID;references:ID" json:"result"`
	QF4ResultID                  int                        `gorm:"column:qf4_result_id;type:integer;size:255;" json:"qf4_result_id"`
	CourseTypeAndManagement      QF4CourseTypeAndManagement `gorm:"foreignKey:QF4CourseTypeAndManagementID;references:ID" json:"course_type_and_management"`
	QF4CourseTypeAndManagementID int                        `gorm:"column:qf4_course_type_and_management_id;type:integer;size:255;" json:"qf4_course_type_and_management_id"`
	Assessment                   QF4Assessment              `gorm:"foreignKey:QF4AssessmentID;references:ID" json:"assessment"`
	QF4AssessmentID              int                        `gorm:"column:qf4_assessment_id;type:integer;size:255;" json:"qf4_assessment_id"`
	Reference                    QF4Reference               `gorm:"foreignKey:QF4ReferenceID;references:ID" json:"reference"`
	QF4ReferenceID               int                        `gorm:"column:qf4_reference_id;type:integer;size:255;" json:"qf4_reference_id"`
	Status                       enums.UserStatus           `gorm:"column:status;type:user_status;not null" json:"status"`
	CreatedAt                    time.Time                  `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt                    time.Time                  `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt                    gorm.DeletedAt             `gorm:"column:deleted_at;type:timestamp"`
}

type QF4CourseInfo struct {
	ID                  int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName        string         `gorm:"column:category_name;type:varchar;size:100;not null" json:"category_name"`
	CourseCode          string         `gorm:"column:course_code;type:varchar;size:255;not null" json:"course_code"`
	CourseNameTH        string         `gorm:"column:course_name_th;type:varchar;size:100;not null" json:"course_name_th"`
	CourseNameEN        string         `gorm:"column:course_name_en;type:varchar;size:100;not null" json:"course_name_en"`
	NumberOfCredits     string         `gorm:"column:number_of_credits;type:varchar;size:13;not null" json:"number_of_credits"`
	CourseType          CourseType     `gorm:"foreignKey:CourseTypeID;references:ID" json:"course_type"`
	CourseTypeID        int            `gorm:"column:course_type_id;type:integer;size:255;not null" json:"course_type_id"`
	CourseConditionTH   string         `gorm:"column:course_condition_th;type:varchar;size:255" json:"course_condition_th"`
	CourseConditionEN   string         `gorm:"column:course_condition_en;type:varchar;size:255" json:"course_condition_en"`
	CourseDescriptionTH string         `gorm:"column:course_description_th;type:varchar;size:255" json:"course_description_th"`
	CourseDescriptionEN string         `gorm:"column:course_description_en;type:varchar;size:255" json:"course_description_en"`
	CourseObjective     string         `gorm:"column:course_objective;type:varchar;size:255" json:"course_objective"`
	StudentActivity     string         `gorm:"column:student_activity;type:varchar;size:255" json:"student_activity"`
	FacilitatorTask     string         `gorm:"column:facilitator_task;type:varchar;size:255" json:"facilitator_task"`
	ConsultantTask      string         `gorm:"column:consultant_task;type:varchar;size:255" json:"consultant_task"`
	StudentGuideline    string         `gorm:"column:student_guideline;type:varchar;size:255" json:"student_guideline"`
	Location            string         `gorm:"column:location;type:varchar;size:255;not null" json:"location"`
	StudentSupport      string         `gorm:"column:student_support;type:varchar;size:255;not null" json:"student_support"`
	CreatedAt           time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type QF4Lecturer struct {
	ID                int              `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName      string           `gorm:"column:category_name;type:varchar;size:100" json:"category_name"`
	CourseOwner       UserDetail       `gorm:"foreignKey:CourseOwnerID;references:UID" json:"course_owner"`
	CourseOwnerID     uuid.UUID        `gorm:"column:course_owner_id;type:uuid" json:"course_owner_id"`
	CourseMapLecturer []MapQF4Lecturer `gorm:"foreignKey:QF4LecturerID;references:ID" json:"course_map_lecturer"`
	CreatedAt         time.Time        `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time        `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt         gorm.DeletedAt   `gorm:"column:deleted_at;type:timestamp"`
}

type MapQF4Lecturer struct {
	ID               int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	QF4LecturerID    int            `gorm:"column:qf4_lecturer_id;type:integer;size:255;not null" json:"qf4_lecturer_id"`
	CourseLecturer   UserDetail     `gorm:"foreignKey:CourseLecturerID;references:UID" json:"course_lecturer"`
	CourseLecturerID uuid.UUID      `gorm:"column:course_lecturer_id;type:uuid" json:"course_lecturer_id"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type QF4Result struct {
	ID                  int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName        string         `gorm:"column:category_name;type:varchar;size:100;not null" json:"category_name"`
	LearningOutcome     string         `gorm:"column:learning_outcome;type:varchar;size:255;not null" json:"learning_outcome"`
	LearningExpectation string         `gorm:"column:learning_expectation;type:varchar;size:255;not null" json:"learning_expectation"`
	CreatedAt           time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type QF4CourseTypeAndManagement struct {
	ID                   int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName         string         `gorm:"column:category_name;type:varchar;size:100;not null" json:"category_name"`
	EducateFormatID      int            `gorm:"column:educate_format_id;type:integer;size:255;not null" json:"educate_format_id"`
	EducateFormatDetail  int            `gorm:"column:educate_format_detail;type:integer;size:100;not null" json:"educate_format_detail"`
	LearningFormatID     int            `gorm:"column:learning_format_id;type:integer;size:255;not null" json:"learning_format_id"`
	LearningFormatDetail int            `gorm:"column:learning_format_detail;type:integer;size:100;not null" json:"learning_format_detail"`
	CreatedAt            time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type QF4Assessment struct {
	ID                 int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName       string         `gorm:"column:category_name;type:varchar;size:100;not null" json:"category_name"`
	LearningAssessment int            `gorm:"column:learning_assessment;type:integer;size:255;not null" json:"learning_assessment"`
	Grade              string         `gorm:"column:grade;type:varchar;size:255;not null" json:"grade"`
	GroupBased         string         `gorm:"column:group_based;type:varchar;size:255;not null" json:"group_based"`
	Other              string         `gorm:"column:other;type:varchar;size:255;not null" json:"other"`
	CreatedAt          time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type QF4CoursePlan struct {
	ID               int `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	uuid.UUID        `gorm:"column:course_main_uuid;type:varchar;size:255;not null" json:"course_main_uuid"`
	Week             string    `gorm:"column:week;type:varchar;size:100;not null" json:"week"`
	Title            string    `gorm:"column:title;type:varchar;size:100;not null" json:"title"`
	PlanDescription  string    `gorm:"column:other;type:varchar;size:100;not null" json:"other"`
	LearningActivity string    `gorm:"column:learning_activity;type:varchar;size:255;not null" json:"learning_activity"`
	EvaluationMethod string    `gorm:"column:evaluation_method;type:varchar;size:255;not null" json:"evaluation_method"`
	Score            int       `gorm:"column:score;type:integer;size:255;not null" json:"score"`
	AssessmentScore  int       `gorm:"column:assessment_score;type:integer;size:255;not null" json:"assessment_score"`
	AssessmentNote   string    `gorm:"column:assessment_note;type:varchar;size:255;not null" json:"assessment_note"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type QF4Reference struct {
	ID                  int       `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	TextbookAndDocument string    `gorm:"column:textbook_and_document;type:varchar;size:255;not null" json:"textbook_and_document"`
	OtherDocument       string    `gorm:"column:other_document;type:varchar;size:255;not null" json:"other_document"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

func (l *QF4) BeforeCreate(tx *gorm.DB) (err error) {
	//create UUID
	l.QF4ID = uuid.New()
	return
}
