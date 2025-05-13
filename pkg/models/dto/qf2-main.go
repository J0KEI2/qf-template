package dto

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type GetMainProgramPaginationQueryParam struct {
	ProgramName *string `query:"program_name"`
	FacultyID   *uint   `query:"faculty_id"`
	FacultyName *string `query:"faculty_name"`
	ProgramId   *string `query:"program_id"`
	ProgramCode *string `query:"program_code"`
}

type GetMainProgramPaginationResponseDto struct {
	Items []ProgramMainPagination `json:"items"`
	models.PaginationOptions
}

type ProgramMainPagination struct {
	ProgramMainID           *uuid.UUID `json:"id"`
	FacultyNameTH           *string    `json:"faculty_name_th"`
	FacultyNameEN           *string    `json:"faculty_name_en"`
	ProgramNameTH           *string    `json:"program_name_th"`
	ProgramNameEN           *string    `json:"program_name_en"`
	ProgramCode             *string    `json:"program_code"`
	BranchNameTH            *string    `json:"branch_name_th"`
	BranchNameEN            *string    `json:"branch_name_en"`
	ProgramType             *string    `json:"program_type"`
	ProgramTypeID           *int       `json:"program_type_id"`
	ProgramYear             *uint      `json:"program_year"`
	ProgramYearID           *int       `json:"program_year_id"`
	PermissionLevel         *uint      `json:"permission_level"`
	CurrentApprovalProgress *uint      `json:"current_approval_progress"`
	CurrentCHECOProgress    *uint      `json:"current_checo_progress"`
}

type ProgramMainRequestDto struct {
	FacultyID     uint   `json:"faculty_id"`
	ProgramNameTH string `json:"program_name_th"`
	ProgramNameEN string `json:"program_name_en"`
	BranchNameTH  string `json:"branch_name_th"`
	BranchNameEN  string `json:"branch_name_en"`
	ProgramTypeID int    `json:"program_type_id"`
	ProgramType   string `json:"program_type"`
	ProgramYear   uint   `json:"program_year"`
	ProgramYearID int    `json:"program_year_id"`
}
