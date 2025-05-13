package repository

import (
	"fmt"

	"github.com/AlekSi/pointer"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
)

func (r commonRepository) GetFacultiesPagination(paginationOptions *models.PaginationOptions) (record []query.Faculty, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}
	dbTx := r.MainDbConn

	db := dbTx.Model(query.Faculty{})

	if paginationOptions.Search != nil {
		search := "%" + pointer.GetString(paginationOptions.Search) + "%"
		db.Where("(faculty_name_th LIKE ? OR faculty_name_en LIKE ?)", search, search)
	}
	
	db.Order("id ASC")
	db.Count(paginationOptions.Total)
	db.Scopes(Paginate(paginationOptions))

	record = make([]query.Faculty, 0)
	if err = db.Find(&record).Error; err != nil {
		return nil, err
	}
	return
}
