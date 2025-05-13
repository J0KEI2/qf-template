package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) CreateProgramRoles(request dto.CreateMapProgramRolesRequestDto) (result *dto.MapProgramRoles, err error) {

	queryMap := query.MapProgramsRolesQueryEntity{
		UserID:    &request.UserID,
		RoleID:    &request.RoleID,
		ProgramID: &request.ProgramID,
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, createMapProgramRoleStatement(u, &queryMap)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	queryRole := query.RoleQueryEntity{
		ID: &request.RoleID,
	}

	u.CommonRepository.GetFirst(&queryRole)

	queryProgram := programQuery.ProgramMainQueryEntity{
		ID: &request.ProgramID,
	}

	u.CommonRepository.GetFirst(&queryProgram, "ProgramGeneralDetail")

	result = &dto.MapProgramRoles{
		ID:            queryMap.ID,
		UserID:        queryMap.UserID,
		RoleID:        queryMap.RoleID,
		ProgramID:     queryMap.ProgramID,
		ProgramNameTH: queryProgram.ProgramGeneralDetail.ProgramNameTH,
		ProgramNameEN: queryProgram.ProgramGeneralDetail.ProgramNameEN,
		RoleNameTH:    queryRole.RoleNameTH,
		RoleNameEN:    queryRole.RoleNameEN,
	}
	return result, nil
}

func createMapProgramRoleStatement(u *PermissionUsecase, queryDb *query.MapProgramsRolesQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Create(tx, queryDb)
	}
}
