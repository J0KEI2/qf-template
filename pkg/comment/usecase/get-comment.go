package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u commentUC) GetComment(queryEntity query.CommentQueryEntity) (*entity.CommentFetchEntity, error) {
	queryDb := entity.CommentFetchEntity{
		ID: queryEntity.ID,
	}

	if err := u.commonRepo.GetFirst(&queryDb); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	result := entity.CommentFetchEntity{
		ID:           queryDb.ID,
		QFType:       queryDb.QFType,
		QFUID:        queryDb.QFUID,
		CategoryType: queryDb.CategoryType,
		Attribute:    queryDb.Attribute,
		Commentator:  queryDb.Commentator,
		Comments:     queryDb.Comments,
		UpdatedBy:    queryDb.UpdatedBy,
		CreatedAt:    queryDb.CreatedAt,
		UpdatedAt:    queryDb.UpdatedAt,
	}

	return &result, nil
}
