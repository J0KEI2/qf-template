package domain

import (
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

type CommentUseCase interface {
	GetComments(queryEntity query.CommentQueryEntity) (*[]entity.CommentFetchEntity, error)
	GetComment(queryEntity query.CommentQueryEntity) (*entity.CommentFetchEntity, error)
	CreateComment(createQuery query.CommentCreateEntity) (*entity.CommentCreateResultEntity, error)
	UpdateComment(id uint, createQuery query.CommentUpdateEntity) (*entity.CommentUpdateResultEntity, error)
	ResolveComment(id uint, resolveQuery bool) (*entity.CommentUpdateResultEntity, error)
}

type CommentRepository interface {
	DbCommentSVCMigrator() (err error)
	GetComments(criteria entity.CommentFetchEntity) (*[]entity.CommentFetchEntity, error)

	// GetComment(criteria entity.CommentFetchQueryEntity) (*entity.CommentFetchEntity, error)
	// CreateComment(dbTx *gorm.DB,
	//
	//	createCommentEntity entity.CommentCreateEntity) (
	//	*entity.CommentFetchEntity,
	//	error)
	//
	// UpdateComment(dbTx *gorm.DB,
	//
	//	updateQueryEntity entity.CommentUpdateQueryEntity,
	//	updateCommentEntity entity.CommentUpdateEntity) (*entity.CommentUpdateReusltEntity,
	//	error)
}
