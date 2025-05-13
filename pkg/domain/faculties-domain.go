package domain

import "github.com/zercle/kku-qf-services/pkg/models/dto"

type FacultiesUsecase interface {
	GetFaculties() (*dto.FacultiesResponseDto, error)
}

type FacultiesRepository interface {
	DbFacultiesSVCMigrator() (err error)
}
