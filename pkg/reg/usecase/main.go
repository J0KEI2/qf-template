package usecase

import (
	"time"

	"github.com/zercle/kku-qf-services/pkg/domain"
)

type regUseCase struct {
	repo            domain.RegRepository
	RegToken        *string
	RegTokenExpired time.Time
}

func NewRegUseCase(repo domain.RegRepository) domain.RegUseCase {
	return &regUseCase{
		repo:            repo,
		RegToken:        nil,
		RegTokenExpired: time.Time{},
	}
}
