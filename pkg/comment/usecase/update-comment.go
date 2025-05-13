package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u commentUC) UpdateComment(id uint, updateQuery query.CommentUpdateEntity) (*entity.CommentUpdateResultEntity, error) {
	queryDb := query.CommentUpdateEntity{
		QFType:       updateQuery.QFType,
		QFUID:        updateQuery.QFUID,
		CategoryType: updateQuery.CategoryType,
		Attribute:    updateQuery.Attribute,
		Commentator:  updateQuery.Commentator,
		Comments:     updateQuery.Comments,
		UpdatedBy:    updateQuery.UpdatedBy,
	}

	if err := helper.ExecuteTransaction(u.commonRepo, updateTransaction(&u, id, &queryDb)); err != nil {
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

func updateTransaction(useCase *commentUC, id uint, update *query.CommentUpdateEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		query := query.CommentUpdateEntity{
			ID: id,
		}
		return useCase.commonRepo.Update(tx, query, &update)
	}
}
