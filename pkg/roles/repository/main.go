package repository

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	"gorm.io/gorm"
)

type roleRepository struct {
	MainDbConn *gorm.DB
}

func NewRoleRepository(mainDbConn *gorm.DB) domain.RoleRepository {
	return &roleRepository{
		MainDbConn: mainDbConn,
	}
}

// func for migrate table and forget to call this func in router.go
func (repo *roleRepository) DbRoleSVCMigrator() (err error) {
	// err = r.MainDbConn.AutoMigrate(&model.TemplateSchemaModel{})
	return
}
