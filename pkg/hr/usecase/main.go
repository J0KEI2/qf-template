package usecase

import (
	"time"

	"github.com/zercle/kku-qf-services/pkg/domain"
)

type hrUseCase struct {
	repo domain.HRRepository
	HRToken *string
	HRTokenExpired time.Time
}

func NewHRUseCase(repo domain.HRRepository) domain.HRUseCase {
	return &hrUseCase{
		repo: repo,
		HRToken: nil,
		HRTokenExpired: time.Time{},
	}
}
