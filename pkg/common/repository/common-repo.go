package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"gorm.io/gorm"
)

func (repo commonRepository) Begin() (tx *gorm.DB, err error) {
	if repo.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return nil, err
	}
	return repo.MainDbConn.Begin(), nil
}

func (repo commonRepository) Commit(tx *gorm.DB) error {
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (repo commonRepository) Rollback(tx *gorm.DB) error {
	if err := tx.Rollback().Error; err != nil {
		return err
	}
	return nil
}
