package query

import "time"

type CourseInfoQueryEntity struct {
	ID                  *int                 `gorm:"column:id"`
	CategoryName        *string              `gorm:"column:category_name"`
	CourseCode          *string              `gorm:"column:course_code"`
	CourseNameTH        *string              `gorm:"column:course_name_th"`
	CourseNameEN        *string              `gorm:"column:course_name_en"`
	TotalCredit         *uint                `gorm:"column:total_credit;type:int;not null;default:0"`
	Credit1             *uint                `gorm:"column:credit_1;type:int;not null;default:0"`
	Credit2             *uint                `gorm:"column:credit_2;type:int;not null;default:0"`
	Credit3             *uint                `gorm:"column:credit_3;type:int;not null;default:0"`
	CourseType          *CourseTypeJoinQuery `gorm:"foreignKey:CourseTypeID;references:ID" json:"course_type"`
	CourseTypeID        *int                 `gorm:"column:course_type_id"`
	CourseConditionTH   *string              `gorm:"column:course_condition_th"`
	CourseConditionEN   *string              `gorm:"column:course_condition_en"`
	CourseDescriptionTH *string              `gorm:"column:course_description_th"`
	CourseDescriptionEN *string              `gorm:"column:course_description_en"`
	CourseObjective     *string              `gorm:"column:course_objective"`
	Location            *string              `gorm:"column:location"`
	CreatedAt           *time.Time           `gorm:"column:created_at"`
	UpdatedAt           *time.Time           `gorm:"column:updated_at"`
}

func (s *CourseInfoQueryEntity) TableName() string {
	return "course_infos"
}

func (s CourseInfoQueryEntity) String() string {
	return "CourseInfo"
}

type CourseInfoJointQuery struct {
	ID                  *int                 `gorm:"column:id"`
	CategoryName        *string              `gorm:"column:category_name"`
	CourseCode          *string              `gorm:"column:course_code"`
	CourseNameTH        *string              `gorm:"column:course_name_th"`
	CourseNameEN        *string              `gorm:"column:course_name_en"`
	NumberOfCredits     *string              `gorm:"column:number_of_credits"`
	CourseType          *CourseTypeJoinQuery `gorm:"foreignKey:CourseTypeID;references:ID" json:"course_type"`
	CourseTypeID        *int                 `gorm:"column:course_type_id"`
	CourseConditionTH   *string              `gorm:"column:course_condition_th"`
	CourseConditionEN   *string              `gorm:"column:course_condition_en"`
	CourseDescriptionTH *string              `gorm:"column:course_description_th"`
	CourseDescriptionEN *string              `gorm:"column:course_description_en"`
	CourseObjective     *string              `gorm:"column:course_objective"`
	Location            *string              `gorm:"column:location"`
	CreatedAt           *time.Time           `gorm:"column:created_at"`
	UpdatedAt           *time.Time           `gorm:"column:updated_at"`
}

func (s *CourseInfoJointQuery) TableName() string {
	return "course_infos"
}

func (s CourseInfoJointQuery) String() string {
	return "CourseInfo"
}

type CourseType struct {
	ID             int    `gorm:"column:id"`
	CourseTypeName string `gorm:"column:course_type_name"`
}

type CourseTypeQueryEntity struct {
	ID             *int    `gorm:"column:id"`
	CourseTypeName *string `gorm:"column:course_type_name"`
}

type CourseTypeJoinQuery struct {
	ID             *int    `gorm:"column:id"`
	CourseTypeName *string `gorm:"column:course_type_name"`
}

func (s *CourseTypeQueryEntity) TableName() string {
	return "course_types"
}

func (s CourseTypeQueryEntity) String() string {
	return "CourseType"
}
