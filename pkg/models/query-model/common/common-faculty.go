package query

type Faculty struct {
	FacultyNameTH *string `gorm:"column:faculty_name_th" json:"facultyNameTH,omitempty"`
	FacultyNameEN *string `gorm:"column:faculty_name_en" json:"facultyNameEN,omitempty"`
	University    *string `gorm:"column:university" json:"university,omitempty"`
	FacultyID     *string `gorm:"column:faculty_id" json:"facultyID,omitempty"`
	ID            *uint   `gorm:"column:id" json:"id"`
}

func (Faculty) TableName() string {
	return "common_faculties"
}
