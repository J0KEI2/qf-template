package query

import "time"

type QF4CourseInfo struct {
	ID                  int       `gorm:"column:id;type:integer;primaryKey"`
	CategoryName        string    `gorm:"column:category_name"`
	CourseCode          string    `gorm:"column:course_code"`
	CourseNameTH        string    `gorm:"column:course_name_th"`
	CourseNameEN        string    `gorm:"column:course_name_en"`
	NumberOfCredits     string    `gorm:"column:number_of_credits"`
	CourseTypeID        int       `gorm:"column:course_type_id"`
	CourseConditionTH   string    `gorm:"column:course_condition_th"`
	CourseConditionEN   string    `gorm:"column:course_condition_en"`
	CourseDescriptionTH string    `gorm:"column:course_description_th"`
	CourseDescriptionEN string    `gorm:"column:course_description_en"`
	CourseObjective     string    `gorm:"column:course_objective"`
	StudentActivity     string    `gorm:"column:student_activity"`
	FacilitatorTask     string    `gorm:"column:facilitator_task"`
	ConsultantTask      string    `gorm:"column:consultant_task"`
	StudentGuideline    string    `gorm:"column:student_guideline"`
	Location            string    `gorm:"column:location"`
	StudentSupport      string    `gorm:"column:student_support"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
}

type QF4CourseInfoQueryEntity struct {
	ID                  *int                    `gorm:"column:id"`
	CategoryName        *string                 `gorm:"column:category_name"`
	CourseCode          *string                 `gorm:"column:course_code"`
	CourseNameTH        *string                 `gorm:"column:course_name_th"`
	CourseNameEN        *string                 `gorm:"column:course_name_en"`
	NumberOfCredits     *string                 `gorm:"column:number_of_credits"`
	CourseType          *QF4CourseTypeJoinQuery `gorm:"foreignKey:CourseTypeID;references:ID" json:"course_type"`
	CourseTypeID        *int                    `gorm:"column:course_type_id"`
	CourseConditionTH   *string                 `gorm:"column:course_condition_th"`
	CourseConditionEN   *string                 `gorm:"column:course_condition_en"`
	CourseDescriptionTH *string                 `gorm:"column:course_description_th"`
	CourseDescriptionEN *string                 `gorm:"column:course_description_en"`
	CourseObjective     *string                 `gorm:"column:course_objective"`
	StudentActivity     *string                 `gorm:"column:student_activity"`
	FacilitatorTask     *string                 `gorm:"column:facilitator_task"`
	ConsultantTask      *string                 `gorm:"column:consultant_task"`
	StudentGuideline    *string                 `gorm:"column:student_guideline"`
	Location            *string                 `gorm:"column:location"`
	StudentSupport      *string                 `gorm:"column:student_support"`
	CreatedAt           *time.Time              `gorm:"column:created_at"`
	UpdatedAt           *time.Time              `gorm:"column:updated_at"`
}

func (s *QF4CourseInfoQueryEntity) TableName() string {
	return "qf4_course_infos"
}

func (s QF4CourseInfoQueryEntity) String() string {
	return "CourseInfo"
}

type QF4CourseInfoJointQuery struct {
	ID                  *int                    `gorm:"column:id"`
	CategoryName        *string                 `gorm:"column:category_name"`
	CourseCode          *string                 `gorm:"column:course_code"`
	CourseNameTH        *string                 `gorm:"column:course_name_th"`
	CourseNameEN        *string                 `gorm:"column:course_name_en"`
	NumberOfCredits     *string                 `gorm:"column:number_of_credits"`
	CourseType          *QF4CourseTypeJoinQuery `gorm:"foreignKey:CourseTypeID;references:ID" json:"course_type"`
	CourseTypeID        *int                    `gorm:"column:course_type_id"`
	CourseConditionTH   *string                 `gorm:"column:course_condition_th"`
	CourseConditionEN   *string                 `gorm:"column:course_condition_en"`
	CourseDescriptionTH *string                 `gorm:"column:course_description_th"`
	CourseDescriptionEN *string                 `gorm:"column:course_description_en"`
	CourseObjective     *string                 `gorm:"column:course_objective"`
	StudentActivity     *string                 `gorm:"column:student_activity"`
	FacilitatorTask     *string                 `gorm:"column:facilitator_task"`
	ConsultantTask      *string                 `gorm:"column:consultant_task"`
	StudentGuideline    *string                 `gorm:"column:student_guideline"`
	Location            *string                 `gorm:"column:location"`
	StudentSupport      *string                 `gorm:"column:student_support"`
	CreatedAt           *time.Time              `gorm:"column:created_at"`
	UpdatedAt           *time.Time              `gorm:"column:updated_at"`
}

func (s *QF4CourseInfoJointQuery) TableName() string {
	return "qf4_course_infos"
}

func (s QF4CourseInfoJointQuery) String() string {
	return "CourseInfo"
}

type QF4CourseType struct {
	ID             int    `gorm:"column:id"`
	CourseTypeName string `gorm:"column:course_type_name"`
}

type QF4CourseTypeQueryEntity struct {
	ID             *int    `gorm:"column:id"`
	CourseTypeName *string `gorm:"column:course_type_name"`
}

type QF4CourseTypeJoinQuery struct {
	ID             *int    `gorm:"column:id"`
	CourseTypeName *string `gorm:"column:course_type_name"`
}

func (s *QF4CourseTypeQueryEntity) TableName() string {
	return "course_types"
}

func (s QF4CourseTypeQueryEntity) String() string {
	return "CourseType"
}
