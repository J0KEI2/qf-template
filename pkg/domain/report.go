package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"gorm.io/gorm"
)

type ReportUsecase interface {
	ExportReport(QFMainID, qfSubplanId *string) (string, error)
	CreateOrUpdateReport(c *fiber.Ctx, request dto.CreateOrUpdateReportRequestDto, userUID uuid.UUID) error
	GetReport(ProgramID uuid.UUID, paginationOptions *models.PaginationOptions) (result dto.GetReportResponseDto, err error)
	DeleteReport(id uint) (err error)
}

type ReportRepository interface {
	ExportReport(QFMainID, qfSubplanId *string) (string, error)
	CreateOrUpdateReportMapFile(tx *gorm.DB, fileId, id *uint) (err error)
	DbReportSVCMigrator() (err error)
}
