package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u commentUC) ResolveComment(id uint, resolveQuery bool) (*entity.CommentUpdateResultEntity, error) {
	queryDb := query.CommentUpdateEntity{
		ID:      id,
		Resolve: resolveQuery,
	}

	if err := helper.ExecuteTransaction(u.commonRepo, updateTransaction(&u, id, &queryDb)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	if err := u.commonRepo.GetFirst(&queryDb); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	result := entity.CommentUpdateResultEntity{
		ID:           id,
		QFType:       &queryDb.QFType,
		QFUID:        &queryDb.QFUID,
		CategoryType: &queryDb.CategoryType,
		Attribute:    &queryDb.Attribute,
		Commentator:  &queryDb.Commentator,
		Comments:     &queryDb.Comments,
		UpdatedBy:    &queryDb.UpdatedBy,
		UpdatedAt:    &queryDb.UpdatedAt,
	}

	return &result, nil
}
