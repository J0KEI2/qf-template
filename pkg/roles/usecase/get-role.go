package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u roleUsecase) GetRoleByID(roleID uint) (result *dto.GetRoleResponseDto, err error) {
	roleQuery := query.RoleQueryEntity{
		ID: &roleID,
	}

	if err = u.CommonRepository.GetFirst(&roleQuery); err != nil {
		return nil, err
	}

	result = &dto.GetRoleResponseDto{
		ID:                   roleQuery.ID,
		RoleNameTH:           roleQuery.RoleNameTH,
		RoleNameEN:           roleQuery.RoleNameEN,
		ProgramRoleType:      roleQuery.ProgramRoleType,
		CourseRoleType:       roleQuery.CourseRoleType,
		SystemLevel:          roleQuery.SystemLevel,
		ProgramApprovalLevel: roleQuery.ProgramApprovalLevel,
		CourseApprovalLevel:  roleQuery.CourseApprovalLevel,
		ProgramActionLevel:   roleQuery.ProgramActionLevel,
		CourseActionLevel:    roleQuery.CourseActionLevel,
		ProgramAccessLevel:   roleQuery.ProgramAccessLevel,
		CourseAccessLevel:    roleQuery.CourseAccessLevel,
	}

	return
}
