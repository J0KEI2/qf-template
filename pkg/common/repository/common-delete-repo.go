package repository

import (
	"fmt"

	"gorm.io/gorm"
)

func (repo commonRepository) Delete(db *gorm.DB,
	query interface{}) (
	err error) {

	db = db.Model(query).Where(query).Delete(query)
	if err = db.Error; err != nil {
		return err
	}

	return nil
}

func (repo commonRepository) DeleteMainQFWithWhereClause(db *gorm.DB,
	value interface{},
	whereClause string,
	arg ...interface{}) (
	err error) {
	db = db.Model(value).Where(fmt.Sprintf("%v = ?", whereClause), arg).Delete(value)
	if err = db.Error; err != nil {
		return err
	}

	return nil
}
