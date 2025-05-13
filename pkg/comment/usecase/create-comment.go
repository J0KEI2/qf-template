package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u commentUC) CreateComment(createQuery query.CommentCreateEntity) (*entity.CommentCreateResultEntity, error) {

	queryDb := query.CommentCreateEntity{
		QFType:       createQuery.QFType,
		QFUID:        createQuery.QFUID,
		CategoryType: createQuery.CategoryType,
		Attribute:    createQuery.Attribute,
		Commentator:  createQuery.Commentator,
		Comments:     createQuery.Comments,
		UpdatedBy:    createQuery.UpdatedBy,
	}

	if err := helper.ExecuteTransaction(u.commonRepo, prepareStatement(&u, &queryDb)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	result := entity.CommentCreateResultEntity{
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

func prepareStatement(useCase *commentUC, queryDb *query.CommentCreateEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return useCase.commonRepo.Create(tx, queryDb)
	}
}
