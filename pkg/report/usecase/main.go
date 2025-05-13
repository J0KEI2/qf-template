package usecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type reportUsecase struct {
	domain.ReportRepository
	domain.CommonRepository
}

func NewReportUsecase(repo domain.ReportRepository, common domain.CommonRepository) domain.ReportUsecase {
	return &reportUsecase{
		ReportRepository: repo,
		CommonRepository: common,
	}
}
