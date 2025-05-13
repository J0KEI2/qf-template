package usecase

import "github.com/zercle/kku-qf-services/pkg/domain"

type qf4Usecase struct {
	qf4Repo          domain.QF4Repository
	CommonRepository domain.CommonRepository
}

func NewQF4Usecase(repo domain.QF4Repository, common domain.CommonRepository) domain.QF4Usecase {
	return &qf4Usecase{
		qf4Repo:          repo,
		CommonRepository: common,
	}
}
