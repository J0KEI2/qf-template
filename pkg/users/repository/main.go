package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/domain"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
	"gorm.io/gorm"
)

type userRepository struct {
	MainDbConn *gorm.DB
	// utils.CommonDatabaseRepository
	// utils.DatabaseCreateRepository
}

func NewUserRepository(mainDbConn *gorm.DB) domain.UserRepository {

	return &userRepository{
		MainDbConn: mainDbConn,
		// CommonDatabaseRepository: utils.CommonDatabaseRepository{DB: mainDbConn},
		// DatabaseCreateRepository: utils.DatabaseCreateRepository{
		// 	DBCreate: mainDbConn,
		// },
	}
}

func (r *userRepository) DbUserSVCMigrator() (err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	r.MainDbConn.Exec(enums.CreateUserStatusEnum())
	r.MainDbConn.Exec(enums.CreateUserTypesEnum())

	if err := r.MainDbConn.AutoMigrate(
		&migrateModels.Users{}); err != nil {
		return err
	}

	return
}
