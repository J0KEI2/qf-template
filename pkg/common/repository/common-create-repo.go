package repository

import "gorm.io/gorm"

func (repo commonRepository) Create(db *gorm.DB,
	value interface{}) (
	err error) {

	db = db.Create(value)
	if err = db.Error; err != nil {
		return err
	}

	return nil
}
