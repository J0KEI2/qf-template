package repository

import (
	"fmt"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (r *commentRepository) GetComments(criteria entity.CommentFetchEntity) (*[]entity.CommentFetchEntity, error) {
	var err error
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return nil, err
	}
	comment := []entity.CommentFetchEntity{}
	dbTx := r.MainDbConn

	if criteria.ID != 0 {
		dbTx = dbTx.Where("id = ?", criteria.ID)
	}

	if criteria.QFType != "" {
		dbTx = dbTx.Where("qf_type = ?", criteria.QFType)
	}

	if criteria.QFUID.String() != "" {
		dbTx = dbTx.Where("qf_uid = ?", criteria.QFUID)
	}

	if criteria.CategoryType != "" {
		dbTx = dbTx.Where("category_type = ?", criteria.CategoryType)
	}

	if criteria.Resolve {
		if err = dbTx.Find(&comment).Error; err != nil {
			return nil, err
		}
	} else {
		if err = dbTx.Where("resolve = false").Find(&comment).Error; err != nil {
			return nil, err
		}
	}

	return &comment, nil
}
