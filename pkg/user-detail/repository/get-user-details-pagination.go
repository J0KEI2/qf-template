package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/common/repository"
	"github.com/zercle/kku-qf-services/pkg/models"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (r *userDetailRepository) GetUserDetailPagination(paginationOptions *models.PaginationOptions) (record []query.UserDetailQueryEntity, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}
	dbTx := r.MainDbConn

	db := dbTx.Model(query.UserDetailQueryEntity{})
	if paginationOptions.Search != nil {
		search := "%" + *paginationOptions.Search + "%"
		db.Where("(firstname_th LIKE ? OR firstname_en LIKE ? OR lastname_th LIKE ? OR lastname_en LIKE ?)", search, search, search, search)
	}

	db.Count(paginationOptions.Total)
	db.Scopes(repository.Paginate(paginationOptions))

	record = make([]query.UserDetailQueryEntity, 0)
	if err = db.Find(&record).Error; err != nil {
		return nil, err
	}
	return
}
