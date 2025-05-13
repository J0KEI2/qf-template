package templateUsecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type templateUsecase struct {
	domain.TemplateRepository
}

func NewTemplateUsecase(repo domain.TemplateRepository) domain.TemplateUseCase {
	return &templateUsecase{
		repo,
	}
}
