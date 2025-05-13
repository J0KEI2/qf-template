package repository

import (
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/domain"
	"gorm.io/gorm"
)

type authRepository struct {
	MainDbConn *gorm.DB
	Fasthttp   *fasthttp.Client
}

func NewAuthRepository(mainDbConn *gorm.DB, fasthttp *fasthttp.Client) domain.AuthRepository {
	return &authRepository{
		MainDbConn: mainDbConn,
		Fasthttp:   fasthttp,
	}
}

// func for migrate table and forget to call this func in router.go
func (repo *authRepository) DbAuthSVCMigrator() (err error) {
	// err = r.MainDbConn.AutoMigrate(&model.AuthSchemaModel{})
	return
}
