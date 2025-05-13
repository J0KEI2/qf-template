package dto

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type ExportReportRequestDto struct {
	QfMainID  *string `query:"program_main_id"`
	SubplanID *string `query:"subplan_id"`
}

type CreateOrUpdateReportRequestDto struct {
	ProgramUID  *uuid.UUID            `query:"program_main_id"`
	ReportID    *uint                 `query:"report_id"`
	Name        *string               `query:"name"`
	Description *string               `query:"description"`
	File        *multipart.FileHeader `query:"file"`
}

type GetReportResponseDto struct {
	ProgramUID *uuid.UUID       `json:"program_main_id"`
	Items      []ReportResponse `json:"items"`
	*models.PaginationOptions
}

type ReportResponse struct {
	ReportID    *uint      `json:"report_id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	FileID      *uint      `json:"file_id"`
	FileName    *string    `json:"file_name"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
