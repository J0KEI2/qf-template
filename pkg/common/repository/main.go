package repository

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	model "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

type commonRepository struct {
	MainDbConn *gorm.DB
}

func NewCommonRepository(mainDbConn *gorm.DB) domain.CommonRepository {
	return &commonRepository{
		MainDbConn: mainDbConn,
	}
}

// func for migrate table and forget to call this func in router.go
func (repo *commonRepository) DbCommonSVCMigrator() (err error) {

	err = repo.MainDbConn.AutoMigrate(
		&model.Faculty{},
		&model.EmployeeDetails{},
		&model.ReferenceOption{},
	)
	return
}
