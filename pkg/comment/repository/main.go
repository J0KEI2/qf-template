package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/domain"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

type commentRepository struct {
	MainDbConn *gorm.DB
}

func NewCommentRepository(mainDbConn *gorm.DB) domain.CommentRepository {
	return &commentRepository{
		MainDbConn: mainDbConn,
	}
}

func (r *commentRepository) DbCommentSVCMigrator() (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	if err := r.MainDbConn.AutoMigrate(
		&migrateModels.Comment{},
	); err != nil {

		return err
	}

	return
}
