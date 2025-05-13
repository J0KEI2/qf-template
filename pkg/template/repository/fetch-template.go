package templateRepository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
)

// TODO: create the new func don't forget to embed in domain repository interface
func (repo templateRepository) FetchTemplate(criteria interface{}) (responseData interface{}, err error) {
	if repo.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	// dbTx := repo.MainDbConn.Model(&models.User{})

	// if len(criteria.ID) != 0 {
	// 	dbTx = dbTx.Where(models.User{ID: criteria.ID})
	// } else {
	// 	if len(criteria.FullName) != 0 {
	// 		dbTx = dbTx.Where("title LIKE ?", "%"+criteria.FullName+"%")
	// 	}
	// }

	// err = dbTx.Find(&responseData).Error

	return
}
