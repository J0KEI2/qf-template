package repository

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	mgModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

type qf4Repository struct {
	MainDbConn *gorm.DB
}

func NewQF4Repository(mainDbConn *gorm.DB) domain.QF4Repository {
	return &qf4Repository{
		MainDbConn: mainDbConn,
	}
}

func (repo *qf4Repository) DbQF4SVCMigrator() (err error) {
	err = repo.MainDbConn.AutoMigrate(
		&mgModels.QF4{},
		&mgModels.QF4CourseInfo{},
		&mgModels.QF4Lecturer{},
		&mgModels.QF4Result{},
		&mgModels.QF4CourseTypeAndManagement{},
		&mgModels.QF4Assessment{},
		&mgModels.QF4CoursePlan{},
		&mgModels.MapQF4Lecturer{},
	)
	return
}
