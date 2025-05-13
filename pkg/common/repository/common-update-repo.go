package repository

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

func (repo commonRepository) Update(tx *gorm.DB,
	query interface{},
	updateValue interface{}) error {
	if query != nil {
		db := tx.Model(query)
		db.Where(query)
		if reflect.ValueOf(query).Kind() == reflect.Ptr {
			interfaceValue := reflect.ValueOf(updateValue).Elem()
			interfaceType := reflect.Indirect(reflect.ValueOf(updateValue)).Type()
			for index := 0; index < interfaceValue.NumField(); index++ {
				interfaceSubValue := interfaceValue.Field(index)
				if !interfaceSubValue.IsZero() {
					if interfaceSubElement := interfaceSubValue.Elem(); interfaceSubElement.IsZero() {
						gormTag := interfaceType.Field(index).Tag.Get("gorm")
						db.Update(ExtractColumnName(gormTag), nil)
						interfaceSubValue.Set(reflect.Zero(interfaceSubValue.Type()))
					}
				}
			}
		}
		updatedResult := db.Debug().Updates(updateValue)
		if err := updatedResult.Error; err != nil {
			return err
		}
		// if updatedResult.RowsAffected == 0 {
		// 	return errors.New("update query not affected")
		// }
		return nil
	} else {
		return errors.New("Update all using common repository is not allow")
	}
}
