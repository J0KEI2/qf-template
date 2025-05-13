package repository

import (
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/domain"
	model "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

type userDetailRepository struct {
	MainDbConn *gorm.DB
	FastHttp   *fasthttp.Client
}

func NewUserDetailRepository(mainDbConn *gorm.DB, fastHttp *fasthttp.Client) domain.UserDetailRepository {
	return &userDetailRepository{
		MainDbConn: mainDbConn,
		FastHttp:   fastHttp,
	}
}

// func for migrate table and forget to call this func in router.go
func (repo *userDetailRepository) DbUserDetailSVCMigrator() (err error) {
	// repo.MainDbConn.Exec(enums.CreateUserDetailPositionEnum())

	err = repo.MainDbConn.AutoMigrate(&model.UserDetail{})
	return
}
