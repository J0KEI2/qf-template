package dto

import "github.com/google/uuid"

type GetLecturerOwnerDto struct {
	Owner          []LecturerDto `json:"owners"`
	ThesisLecturer []LecturerDto `json:"thesis_lecturers"`
	Lecturer       []LecturerDto `json:"lecturers"`
}

type CreateOrUpdateLecturerOwnerDto struct {
	Owner          []LecturerDto `json:"owners"`
	ThesisLecturer []LecturerDto `json:"thesis_lecturers"`
	Lecturer       []LecturerDto `json:"lecturers"`
}

type LecturerDto struct {
	ID                   *uint                      `json:"id"`
	UserUID              *uuid.UUID                 `json:"user_uid"`
	Name                 *string                    `json:"name"`
	NameEN               *string                    `json:"name_en"`
	TitleTh              *string                    `json:"title_th"`
	FirstnameTh          *string                    `json:"firstname_th"`
	LastnameTh           *string                    `json:"lastname_th"`
	TitleEn              *string                    `json:"title_en"`
	FirstnameEn          *string                    `json:"firstname_en"`
	LastnameEn           *string                    `json:"lastname_en"`
	Email                *string                    `json:"email"`
	Positions            []LecturerPositionDto      `json:"positions"`
	EducationBackgrounds []EducationalBackgroundDto `json:"education_backgrounds"`
}

type LecturerPositionDto struct {
	Position *string `json:"position"`
	Year     *string `json:"year"`
}

type EducationalBackgroundDto struct {
	EducationLevel   *string `json:"education_level"`
	Qualification    *string `json:"qualification"`
	Department       *string `json:"department"`
	InstituteName    *string `json:"institute_name"`
	InstituteCountry *string `json:"institute_country"`
	GraduateYear     *string `json:"graduate_year"`
}
