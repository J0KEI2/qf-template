package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u commentUC) GetComments(queryEntity query.CommentQueryEntity) (*[]entity.CommentFetchEntity, error) {
	queryDb := entity.CommentFetchEntity{
		QFType:       queryEntity.QFType,
		QFUID:        queryEntity.QFUID,
		CategoryType: queryEntity.CategoryType,
		Resolve:      queryEntity.Resolve,
	}

	return u.commentRepo.GetComments(queryDb)
}
