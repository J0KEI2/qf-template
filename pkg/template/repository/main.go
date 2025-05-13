package templateRepository

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	"gorm.io/gorm"
)

type templateRepository struct {
	MainDbConn *gorm.DB
}

func NewTemplateRepository(mainDbConn *gorm.DB) domain.TemplateRepository {
	return &templateRepository{
		MainDbConn: mainDbConn,
	}
}

// func for migrate table and forget to call this func in router.go
func (repo *templateRepository) DbTemplateSVCMigrator() (err error) {
	// err = r.MainDbConn.AutoMigrate(&model.TemplateSchemaModel{})
	return
}
