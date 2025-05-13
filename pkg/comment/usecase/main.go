package usecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type commentUC struct {
	commentRepo domain.CommentRepository
	commonRepo  domain.CommonRepository
}

func NewCommentUseCase(commentRepo domain.CommentRepository, commonRepo domain.CommonRepository) domain.CommentUseCase {
	return &commentUC{
		commentRepo: commentRepo,
		commonRepo:  commonRepo,
	}
}
