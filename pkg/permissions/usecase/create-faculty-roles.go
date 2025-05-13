package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	commonQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) CreateFacultyRoles(request dto.CreateMapFacultyRolesRequestDto) (result *dto.MapFacultyRoles, err error) {

	queryMap := query.MapFacultiesRolesQueryEntity{
		UserID:    &request.UserID,
		RoleID:    &request.RoleID,
		FacultyID: &request.FacultyID,
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, createMapFacultyRoleStatement(u, &queryMap)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	queryRole := query.RoleQueryEntity{
		ID: &request.RoleID,
	}

	u.CommonRepository.GetFirst(&queryRole)

	queryFaculty := commonQuery.Faculty{
		ID: &request.FacultyID,
	}

	u.CommonRepository.GetFirst(&queryFaculty)

	result = &dto.MapFacultyRoles{
		ID:            queryMap.ID,
		UserID:        queryMap.UserID,
		RoleID:        queryMap.RoleID,
		FacultyID:     queryMap.FacultyID,
		FacultyNameTH: queryFaculty.FacultyNameTH,
		FacultyNameEN: queryFaculty.FacultyNameEN,
		RoleNameTH:    queryRole.RoleNameTH,
		RoleNameEN:    queryRole.RoleNameEN,
	}
	return result, nil
}

func createMapFacultyRoleStatement(u *PermissionUsecase, queryDb *query.MapFacultiesRolesQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Create(tx, queryDb)
	}
}
