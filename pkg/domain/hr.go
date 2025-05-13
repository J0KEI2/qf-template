package domain

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

type HRUseCase interface {
	GetHRToken() (*string, error)
	GetEducationByEmail(email string) ([]dto.HREducationDetail, error)
	GetLecturerPagination(options models.PaginationOptions) (*dto.GetEmployeesPaginationResponseDto, error)
	GetLecturerBySSN(ssn string) (*dto.GetSingleEmployeeResponseDto, error) 
}

type HRRepository interface {
	GetHrToken() (*string, error)
	GetEducationByEmail(HRKey, email string) (*dto.GetEducationsResponseDto, error)
	GetLecturerPagination(HRKey, firstname, lastname string, options models.PaginationOptions) (*dto.GetEmployeesResponseDto, error)
	GetLecturerBySSN(HRKey, ssn string) (*dto.GetSingleEmployeeResponseDto, error)
}
