package repository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	migrate_models "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (r *userRepository) GetUserByID(uid string) (*models.UserFetchModel, error) {
	if r.MainDbConn == nil {
		err := fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	uidConverted, err := uuid.Parse(uid)
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	var userRes models.UserFetchModel
	dbTx := r.MainDbConn.Model(&migrate_models.Users{})
	if err = dbTx.Where(migrate_models.Users{UID: uidConverted}).First(&userRes).Error; err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	return &userRes, nil
}
