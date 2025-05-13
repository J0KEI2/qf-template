package usecase

import (
	"log"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
	rapQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u programUsecase) GetMainProgramPagination(userUID uuid.UUID, roleID *uint, options models.PaginationOptions, param dto.GetMainProgramPaginationQueryParam) (result *dto.GetMainProgramPaginationResponseDto, err error) {

	userStatement := query.UserQueryEntity{
		UID: &userUID,
	}

	err = u.CommonRepository.GetFirst(&userStatement)

	if err != nil {
		return nil, err
	}

	roleStatement := rapQuery.RoleQueryEntity{
		ID: roleID,
	}

	err = u.CommonRepository.GetFirst(&roleStatement)

	if err != nil {
		return nil, err
	}

	log.Printf("user : %v \n role : %v", userUID, *roleID)
	switch pointer.GetUint(roleStatement.ProgramAccessLevel) {
	case 2:
		permissionStatement := rapQuery.MapFacultiesRolesQueryEntity{UserID: &userUID, RoleID: roleID}
		permissionResult := make([]rapQuery.MapFacultiesRolesQueryEntity, 0)
		err = u.CommonRepository.GetList(&permissionStatement, &permissionResult, nil)
		roleStatement.MapFacultiesRoles = permissionResult
	case 1:
		permissionStatement := rapQuery.MapProgramsRolesQueryEntity{UserID: &userUID, RoleID: roleID}
		permissionResult := make([]rapQuery.MapProgramsRolesQueryEntity, 0)
		err = u.CommonRepository.GetList(&permissionStatement, &permissionResult, nil)
		roleStatement.MapProgramRoles = permissionResult
	}

	programMainRecords, err := u.ProgramRepository.GetProgramMainPagination(&userStatement, &roleStatement, &options, param)

	if err != nil {
		return nil, err
	}
	items := make([]dto.ProgramMainPagination, 0)

	for _, programMain := range programMainRecords {
		main := programMain
		Qf2MainDetail := dto.ProgramMainPagination{
			ProgramMainID: main.ID,
		}

		// find approval progress
		currLevel, err := u.approval.GetCurrentApprovalProgress(*main.ID)
		if err != nil {
			return nil, err
		}

		// find checo progress
		currCheco, err := u.approval.GetCurrentCHECOProgress(*main.ID)
		if err != nil {
			return nil, err
		}

		Qf2MainDetail.PermissionLevel = roleStatement.ProgramActionLevel
		if detail := main.ProgramGeneralDetail; detail != nil {
			Qf2MainDetail.BranchNameEN = detail.BranchNameEN
			Qf2MainDetail.BranchNameTH = detail.BranchNameTH
			Qf2MainDetail.ProgramNameEN = detail.ProgramNameEN
			Qf2MainDetail.ProgramNameTH = detail.ProgramNameTH
			Qf2MainDetail.ProgramCode = detail.ProgramCode
			Qf2MainDetail.ProgramType = detail.ProgramType
			Qf2MainDetail.ProgramYear = detail.ProgramYear
			if faculty := detail.Faculty; faculty != nil {
				Qf2MainDetail.FacultyNameEN = faculty.FacultyNameEN
				Qf2MainDetail.FacultyNameTH = faculty.FacultyNameTH
			}
			Qf2MainDetail.CurrentApprovalProgress = currLevel
			Qf2MainDetail.CurrentCHECOProgress = currCheco
		}
		items = append(items, Qf2MainDetail)
	}

	result = &dto.GetMainProgramPaginationResponseDto{
		Items:             items,
		PaginationOptions: options,
	}

	return
}
