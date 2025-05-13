package domain

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

type UserDetailUseCase interface {
	SearchLecturerByName(accessToken string, criteria *models.LecturerFetchWithNameRequestModel) (*models.LecturerFetchWithNameResponseModel, error)
	CronUpdateLecturers(accessToken string) error
	GetUserDetailPagination(options models.PaginationOptions) (result *dto.GetUserDetailPaginationResponseDto, err error)
}

type UserDetailRepository interface {
	GetAllEmployees(accessToken string, criteria *models.LecturerFetchWithNameQueryModel) (*models.LecturerFetchWithNameResponseModel, error)
	GetAllHrKeyLecturers() (users []models.CronUpdateLecturer, err error)
	UpdateLecturer(uid string, user migrateModels.UserDetail) (err error)
	GetSingleEmployee(hrKey, accessToken string) ([]byte, error)
	DbUserDetailSVCMigrator() (err error)
	GetUserDetailPagination(paginationOptions *models.PaginationOptions) (record []query.UserDetailQueryEntity, err error)
}
