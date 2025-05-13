package dto

import (
	"github.com/zercle/kku-qf-services/pkg/models"
)

type GetFacultyPaginationResponseDto struct {
	Items   []FacultyResponseDto     `json:"items"`
	Options models.PaginationOptions `json:"options"`
}

type FacultyResponseDto struct {
	ID            *uint   `json:"id"`
	FacultyNameEN *string `json:"faculty_name_en"`
	FacultyNameTH *string `json:"faculty_name_th"`
	University    *string `json:"university"`
}

type FacultiesResponseDto struct {
	Items []FacultyResponseDto `json:"items"`
}
