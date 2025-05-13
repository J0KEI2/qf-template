package usecase

import (
	"sort"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u reportUsecase) GetReport(programID uuid.UUID, paginationOptions *models.PaginationOptions) (result dto.GetReportResponseDto, err error) {
	queryDb := query.ReportQueryEntity{
		QFMainID: &programID,
	}

	reports := []query.ReportQueryEntity{}

	err = u.CommonRepository.GetList(queryDb, &reports, paginationOptions)
	if err != nil {
		return
	}

	reportResult := make([]dto.ReportResponse, 0)
	for _, report := range reports {
		fileQuery := query.MapFilesSystemQueryEntity{
			ReportID: report.ID,
		}
		if err = u.CommonRepository.GetFirst(&fileQuery, "FileSystem"); err != nil {
			if err == gorm.ErrRecordNotFound {
				reportResult = append(reportResult, dto.ReportResponse{
					ReportID:    report.ID,
					Name:        report.Name,
					Description: report.Description,
					FileID:      nil,
					FileName:    nil,
					CreatedAt:   report.CreatedAt,
					UpdatedAt:   report.UpdatedAt,
				})

				continue
			} else {
				return dto.GetReportResponseDto{}, err
			}
		}

		reportResult = append(reportResult, dto.ReportResponse{
			ReportID:    report.ID,
			Name:        report.Name,
			Description: report.Description,
			FileID:      fileQuery.FileSystem.ID,
			FileName:    fileQuery.FileSystem.FileName,
			CreatedAt:   report.CreatedAt,
			UpdatedAt:   report.UpdatedAt,
		})
	}

	sort.SliceStable(reportResult, func(i, j int) bool {
		return *reportResult[i].ReportID < *reportResult[j].ReportID
	})

	result = dto.GetReportResponseDto{
		ProgramUID:        &programID,
		Items:             reportResult,
		PaginationOptions: paginationOptions,
	}

	return result, nil
}
