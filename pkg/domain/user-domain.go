package domain

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	migrate_models "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

type UserUsecase interface {
	GetUserByID(uid uuid.UUID) (*models.UserFetchModel, error)
	GetUserByIDWithRoleSystem(uid string) (*models.UserFetchWithSystemRoleModel, error)
	// GetOrCreateBySSN(userCreateQuery models.UserCreateQuery) (*entity.UserFetchEntity, error)
	GetOrCreateBySSN(userCreateQuery models.UserCreateQuery) (*entity.UserFetchEntity, error)
	GetUsers(userQueryEntity query.UserQueryEntity, options *models.PaginationOptions) (userFetchEntity *entity.UserListFetchEntity, err error)
	CreateUser(createUserDto models.UserCreateQuery) (userResult *entity.UserFetchEntity, err error)
	GetUserList(entity entity.UserFetchQueryEntity, option *models.PaginationOptions) (*entity.UserFetchListEntity, error)
	EditUser(userID string, user models.PatchUserRequest) (err error)
	DeleteUser(userID string) (err error)
	UpdateRoleUser(userID string, patchRequest *models.UpdateUserRoleRequest) error
}

type UserRepository interface {
	GetUsers(criteria *models.UserFetchQuery) (users []models.UserFetchModel, err error)
	GetUser(criteria *models.UserFetchQuery) (user *entity.UserFetchEntity, err error)
	GetUserByID(uid string) (*models.UserFetchModel, error)
	GetUserByIDWithRoleSystem(uid string) (*models.UserFetchWithSystemRoleModel, error)
	GetOrCreateBySSN(userCreateQuery models.UserCreateQuery) (*entity.UserFetchEntity, error)
	CreateUser(createUserDto models.UserCreateQuery) (userResult *entity.UserFetchEntity, err error)
	GetUserList(criteria *models.UserFetchWithRelationQueryModel, options *models.PaginationOptions) (users []models.UserFetchWithRelationModel, err error)
	CountUserList(criteria *models.UserFetchWithRelationQueryModel) (count *int64, err error)
	EditUser(userID string, user migrate_models.Users) (err error)
	DeleteUser(userID string) (err error)
	UpdateUser(uid string, user migrate_models.Users) (err error)
	UpdateRoleUser(uid string, patchRequest models.UpdateUserRoleRequest) (err error)
	DeleteMapCourseByUID(uid string) (err error)
	DeleteMapProgramByUID(uid string) (err error)
	CreateMapCourse(newData migrate_models.MapUserCourse) (err error)
	CreateMapProgram(newData migrate_models.MapUserProgram) (err error)
	DbUserSVCMigrator() (err error)
}
