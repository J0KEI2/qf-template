package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type LeturerDetailDto struct {
	UID         uuid.UUID                `json:"uid"`
	TitleTh     string                   `json:"title_th"`
	TitleEn     string                   `json:"title_en"`
	FirstnameTh string                   `json:"firstname_th"`
	FirstnameEn string                   `json:"firstname_en"`
	LastnameTh  string                   `json:"lastname_th"`
	LastnameEn  string                   `json:"lastname_en"`
	Position    string                   `json:"position"`
	Educations  []map[string]interface{} `json:"educations"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}

type GetEmployeesResponseDto struct {
	Status string                  `json:"status"`
	Data   HREmployeePaginationDto `json:"data"`
}

type HREmployeePaginationDto struct {
	TotalItems  int           `json:"totalItems"`
	Items       []HREmployees `json:"items"`
	TotalPages  int           `json:"totalPages"`
	CurrentPage int           `json:"currentPage"`
}

type HREmployees struct {
	TitleTh     string `json:"title"`
	FirstnameTh string `json:"firstname"`
	LastnameTh  string `json:"lastname"`
	TitleEn     string `json:"titleEng"`
	FirstnameEn string `json:"firstnameEng"`
	LastnameEn  string `json:"lastnameEng"`
	Email       string `json:"email"`
	Position    string `json:"position"`
	Faculty     string `json:"faculty"`
}

type GetEmployeesPaginationResponseDto struct {
	Items   []HREmployeesResponseDto `json:"items"`
	Options models.PaginationOptions `json:"options"`
}

type HREmployeesResponseDto struct {
	TitleTh     string                `json:"title_th"`
	FirstnameTh string                `json:"firstname_th"`
	LastnameTh  string                `json:"lastname_th"`
	Name        string                `json:"name"`
	TitleEn     string                `json:"title_en"`
	FirstnameEn string                `json:"firstname_en"`
	LastnameEn  string                `json:"lastname_en"`
	NameEn      string                `json:"name_en"`
	Email       string                `json:"email"`
	Position    []LecturerPositionDto `json:"positions"`
	Faculty     string                `json:"faculty"`
}

type GetSingleEmployeeResponseDto struct {
	Status string      `json:"status"`
	Data   HREmployees `json:"data"`
}
