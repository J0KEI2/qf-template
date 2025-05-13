package domain

import (
	queryInterface "github.com/zercle/kku-qf-services/pkg/common/models"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	"gorm.io/gorm"
)

type QueryInterface interface {
	GetDBQuery(db *gorm.DB) *gorm.DB
}
type CommonUseCase interface {
	GetFacultiesPagination(options models.PaginationOptions) (result *dto.GetFacultyPaginationResponseDto, err error)
	GetAllFaculties() (result []dto.FacultyResponseDto, err error)
	GetFacultyByFacultyName(facultyName string) (result *dto.FacultyResponseDto, err error)
	GetFacultyByID(id uint) (result *dto.FacultyResponseDto, err error)
	GetAllReferenceOption() (result []dto.ReferenceOption, err error)
	GetFilePathByID(fileID uint) (filePath *string, err error)
}

type CommonRepository interface {
	Begin() (tx *gorm.DB, err error)
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
	Create(db *gorm.DB, value interface{}) (err error)
	Delete(db *gorm.DB, value interface{}) (err error)
	Update(tx *gorm.DB, query interface{}, updateValue interface{}) error
	GetList(queryTb queryInterface.StatementInterface, dest interface{}, options *models.PaginationOptions, joinTB ...string) (err error)
	GetListWithNilSearch(queryTb queryInterface.StatementInterface, dest interface{}, options *models.PaginationOptions, joinTB ...string) (err error)
	GetFirst(tb interface{}, joinTB ...string) (err error)
	DeleteMainQFWithWhereClause(db *gorm.DB, value interface{}, whereClause string, arg ...interface{}) (err error)
	DbCommonSVCMigrator() (err error)
	GetFacultiesPagination(paginationOptions *models.PaginationOptions) (record []query.Faculty, err error)
	GetFirstOrCreate(tb interface{}, joinTB ...string) (err error)
}
