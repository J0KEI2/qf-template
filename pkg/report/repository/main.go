package repository

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

type reportRepository struct {
	MainDbConn *gorm.DB
}

func NewReportRepository(mainDbConn *gorm.DB) domain.ReportRepository {
	return &reportRepository{
		MainDbConn: mainDbConn,
	}
}

// func for migrate table and forget to call this func in router.go
func (repo *reportRepository) DbReportSVCMigrator() (err error) {
	err = repo.MainDbConn.AutoMigrate(
		&migrateModels.Report{},
	)
	return
}
