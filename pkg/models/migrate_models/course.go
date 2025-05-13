package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	CourseID                  uuid.UUID               `gorm:"column:course_id;type:uuid;primaryKey" json:"course_id"`
	REGKkuKey                 string                  `gorm:"column:reg_kku_key;type:varchar;size:255"`
	CourseNumber              int                     `gorm:"column:course_number;type:integer;size:255" json:"course_number"`
	Version                   string                  `gorm:"column:version;type:varchar;size:255;not null" json:"version"`
	Faculty                   Faculty                 `gorm:"foreignKey:FacultyID;references:ID" json:"faculty"`
	FacultyID                 uint                    `gorm:"column:faculty_id;type:uint;not null" json:"faculty_id"`
	DepartmentName            string                  `gorm:"column:department_name;type:varchar;size:100;not null" json:"department_name"`
	EducationYear             string                  `gorm:"column:education_year;type:varchar;size:100;not null" json:"education_year"`
	CourseInfo                CourseInfo              `gorm:"foreignKey:CourseInfoID;references:ID" json:"course_info"`
	CourseInfoID              int                     `gorm:"column:course_info_id;type:integer;size:255" json:"course_info_id"`
	Lecturer                  CourseLecturer          `gorm:"foreignKey:CourseLecturerID;references:ID" json:"lecturer"`
	CourseLecturerID          int                     `gorm:"column:course_lecturer_id;type:integer;size:255" json:"course_lecturer_id"`
	Result                    CourseResult            `gorm:"foreignKey:CourseResultID;references:ID" json:"result"`
	CourseResultID            int                     `gorm:"column:course_result_id;type:integer;size:255" json:"course_result_id"`
	CourseTypeAndManagement   CourseTypeAndManagement `gorm:"foreignKey:CourseTypeAndManagementID;references:ID" json:"course_type_and_management"`
	CourseTypeAndManagementID int                     `gorm:"column:course_type_and_management_id;type:integer;size:255" json:"course_type_and_management_id"`
	Assessment                CourseAssessment        `gorm:"foreignKey:CourseAssessmentID;references:ID" json:"assessment"`
	CourseAssessmentID        int                     `gorm:"column:course_assessment_id;type:integer;size:255" json:"course_assessment_id"`
	CourseReferenceID         int                     `gorm:"column:course_reference_id;type:integer;size:255" json:"course_reference_id"`
	Status                    string                  `gorm:"column:status;type:varchar;size:255;default:'draft'" json:"status"`
	CreatedAt                 time.Time               `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt                 time.Time               `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt                 gorm.DeletedAt          `gorm:"column:deleted_at;type:timestamp"`
}

type CourseInfo struct {
	ID                  int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName        string         `gorm:"column:category_name;type:varchar;size:255;not null" json:"category_name"`
	CourseCode          string         `gorm:"column:course_code;type:varchar;size:255;not null" json:"course_code"`
	CourseNameTH        string         `gorm:"column:course_name_th;type:varchar;size:255;not null" json:"course_name_th"`
	CourseNameEN        string         `gorm:"column:course_name_en;type:varchar;size:255;not null" json:"course_name_en"`
	TotalCredit         uint           `gorm:"column:total_credit;type:int;not null;default:0" json:"total_credit"`
	Credit1             uint           `gorm:"column:credit_1;type:int;not null;default:0" json:"credit_1"`
	Credit2             uint           `gorm:"column:credit_2;type:int;not null;default:0" json:"credit_2"`
	Credit3             uint           `gorm:"column:credit_3;type:int;not null;default:0" json:"credit_3"`
	CourseType          CourseType     `gorm:"foreignKey:CourseTypeID;references:ID" json:"course_type"`
	CourseTypeID        int            `gorm:"column:course_type_id;type:integer;size:255" json:"course_type_id"`
	CourseConditionTH   string         `gorm:"column:course_condition_th;type:varchar;size:255" json:"course_condition_th"`
	CourseConditionEN   string         `gorm:"column:course_condition_en;type:varchar;size:255" json:"course_condition_en"`
	CourseDescriptionTH string         `gorm:"column:course_description_th;type:varchar;size:255" json:"course_description_th"`
	CourseDescriptionEN string         `gorm:"column:course_description_en;type:varchar;size:255" json:"course_description_en"`
	CourseObjective     string         `gorm:"column:course_objective;type:varchar;size:255" json:"course_objective"`
	Location            string         `gorm:"column:location;type:varchar;size:255;not null" json:"location"`
	CreatedAt           time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type CourseLecturer struct {
	ID                int                 `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName      string              `gorm:"column:category_name;type:varchar;size:255" json:"category_name"`
	CourseOwner       UserDetail          `gorm:"foreignKey:CourseOwnerID;references:UID" json:"course_owner"`
	CourseOwnerID     uuid.UUID           `gorm:"column:course_owner_id;type:uuid" json:"course_owner_id"`
	CourseMapLecturer []MapCourseLecturer `gorm:"foreignKey:CourseLecturerID;references:ID" json:"course_map_lecturer"`
	CreatedAt         time.Time           `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time           `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt         gorm.DeletedAt      `gorm:"column:deleted_at;type:timestamp"`
}

type MapCourseLecturer struct {
	ID               int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CourseLecturerID int            `gorm:"column:course_lecturer_id;type:integer;size:255;not null" json:"course_lecturer_id"`
	CourseLecturer   UserDetail     `gorm:"foreignKey:EmployeeID;references:UID" json:"course_lecturer"`
	EmployeeID       uuid.UUID      `gorm:"column:employee_id;type:uuid" json:"employee_id"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type CourseResult struct {
	ID                  int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName        string         `gorm:"column:category_name;type:varchar;size:255;not null" json:"category_name"`
	LearningOutcome     string         `gorm:"column:learning_outcome;type:varchar;size:255;not null" json:"learning_outcome"`
	LearningExpectation string         `gorm:"column:learning_expectation;type:varchar;size:255;not null" json:"learning_expectation"`
	CreatedAt           time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type CourseTypeAndManagement struct {
	ID                   int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName         string         `gorm:"column:category_name;type:varchar;size:255;not null" json:"category_name"`
	EducateFormatID      int            `gorm:"column:educate_format_id;type:integer;size:255;not null" json:"educate_format_id"`
	EducateFormatDetail  int            `gorm:"column:educate_format_detail;type:integer;size:255;not null" json:"educate_format_detail"`
	LearningFormatID     int            `gorm:"column:learning_format_id;type:integer;size:255;not null" json:"learning_format_id"`
	LearningFormatDetail int            `gorm:"column:learning_format_detail;type:integer;size:255;not null" json:"learning_format_detail"`
	CreatedAt            time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type CourseAssessment struct {
	ID                 int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	CategoryName       string         `gorm:"column:category_name;type:varchar;size:255;not null" json:"category_name"`
	LearningAssessment int            `gorm:"column:learning_assessment;type:integer;size:255;not null" json:"learning_assessment"`
	Grade              string         `gorm:"column:grade;type:varchar;size:255;not null" json:"grade"`
	GroupBased         string         `gorm:"column:group_based;type:varchar;size:255;not null" json:"group_based"`
	Other              string         `gorm:"column:other;type:varchar;size:255;not null" json:"other"`
	CreatedAt          time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type CoursePlan struct {
	ID               int            `gorm:"column:id;type:integer;primaryKey;index;not null" json:"id" sql:"AUTO_INCREMENT"`
	Qf3MainUid       uuid.UUID      `gorm:"column:course_main_uuid;type:uuid;not null" json:"course_main_uuid"`
	Week             string         `gorm:"column:week;type:varchar;size:255;not null" json:"week"`
	Title            string         `gorm:"column:title;type:varchar;size:255;not null" json:"title"`
	PlanDescription  string         `gorm:"column:other;type:varchar;size:255;not null" json:"other"`
	TheoryHour       int            `gorm:"column:theory_hour;type:integer;size:255;not null" json:"theory_hour"`
	OperationHour    int            `gorm:"column:operation_hour;type:integer;size:255;not null" json:"operation_hour"`
	SelfLearningHour int            `gorm:"column:self_learning_hour;type:integer;size:255;not null" json:"self_learning_hour"`
	LearningOutcome  string         `gorm:"column:learning_outcome;type:varchar;size:255;not null" json:"learning_outcome"`
	TeachingMedia    string         `gorm:"column:teaching_outcome;type:varchar;size:255;not null" json:"teaching_outcome"`
	LeaningActivity  string         `gorm:"column:learning_activity;type:varchar;size:255;not null" json:"learning_activity"`
	EvaluationMethod string         `gorm:"column:evaluation_method;type:varchar;size:255;not null" json:"evaluation_method"`
	Lecturer         string         `gorm:"column:lecturer;type:varchar;size:255;not null" json:"lecturer"`
	AssessmentScore  int            `gorm:"column:assessment_score;type:integer;size:255;not null" json:"assessment_score"`
	AssessmentNote   string         `gorm:"column:assessment_note;type:varchar;size:255;not null" json:"assessment_note"`
	TCourse          Course         `gorm:"foreignKey:Qf3MainUid"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

func (l *Course) BeforeCreate(tx *gorm.DB) (err error) {
	//create UUID
	l.CourseID = uuid.New()
	return
}
