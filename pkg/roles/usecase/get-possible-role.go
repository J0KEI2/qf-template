package usecase

import (
	"sort"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u roleUsecase) GetPossibleRole(userID uuid.UUID) (result *dto.GetPossibleRolesResponseDto, err error) {
	mapCoreRoleQuery := query.MapCoreRolesQueryEntity{
		UserID: &userID,
	}

	mapFacultyRoleQuery := query.MapFacultiesRolesQueryEntity{
		UserID: &userID,
	}

	mapProgramRoleQuery := query.MapProgramsRolesQueryEntity{
		UserID: &userID,
	}

	mapCoreRoleList := []query.MapCoreRolesQueryEntity{}
	mapFacultyRoleList := []query.MapFacultiesRolesQueryEntity{}
	mapProgramRoleList := []query.MapProgramsRolesQueryEntity{}

	if err = u.CommonRepository.GetList(&mapCoreRoleQuery, &mapCoreRoleList, nil, "Role"); err != nil {
		return nil, err
	}

	if err = u.CommonRepository.GetList(&mapFacultyRoleQuery, &mapFacultyRoleList, nil, "Role"); err != nil {
		return nil, err
	}

	if err = u.CommonRepository.GetList(&mapProgramRoleQuery, &mapProgramRoleList, nil, "Role"); err != nil {
		return nil, err
	}

	roleResponse := []dto.GetRoleResponseDto{}
	roleMap := make(map[uint]*dto.GetRoleResponseDto)
	for _, input := range mapCoreRoleList {
		if _, exists := roleMap[*input.RoleID]; exists && input.Role != nil {
			continue
		} else {
			roleMap[*input.RoleID] = &dto.GetRoleResponseDto{
				ID:                   input.RoleID,
				RoleNameTH:           input.Role.RoleNameTH,
				RoleNameEN:           input.Role.RoleNameEN,
				ProgramRoleType:      input.Role.ProgramRoleType,
				CourseRoleType:       input.Role.CourseRoleType,
				SystemLevel:          input.Role.SystemLevel,
				ProgramApprovalLevel: input.Role.ProgramApprovalLevel,
				CourseApprovalLevel:  input.Role.CourseApprovalLevel,
				ProgramActionLevel:   input.Role.ProgramActionLevel,
				CourseActionLevel:    input.Role.CourseActionLevel,
				ProgramAccessLevel:   input.Role.ProgramAccessLevel,
				CourseAccessLevel:    input.Role.CourseAccessLevel,
			}
		}
	}

	for _, input := range mapFacultyRoleList {
		if _, exists := roleMap[*input.RoleID]; exists {
			continue
		} else {
			roleMap[*input.RoleID] = &dto.GetRoleResponseDto{
				ID:                   input.RoleID,
				RoleNameTH:           input.Role.RoleNameTH,
				RoleNameEN:           input.Role.RoleNameEN,
				ProgramRoleType:      input.Role.ProgramRoleType,
				CourseRoleType:       input.Role.CourseRoleType,
				ProgramApprovalLevel: input.Role.ProgramApprovalLevel,
				CourseApprovalLevel:  input.Role.CourseApprovalLevel,
				ProgramActionLevel:   input.Role.ProgramActionLevel,
				CourseActionLevel:    input.Role.CourseActionLevel,
			}
		}
	}

	for _, input := range mapProgramRoleList {
		if _, exists := roleMap[*input.RoleID]; exists {
			continue
		} else {
			roleMap[*input.RoleID] = &dto.GetRoleResponseDto{
				ID:                   input.RoleID,
				RoleNameTH:           input.Role.RoleNameTH,
				RoleNameEN:           input.Role.RoleNameEN,
				ProgramRoleType:      input.Role.ProgramRoleType,
				CourseRoleType:       input.Role.CourseRoleType,
				ProgramApprovalLevel: input.Role.ProgramApprovalLevel,
				CourseApprovalLevel:  input.Role.CourseApprovalLevel,
				ProgramActionLevel:   input.Role.ProgramActionLevel,
				CourseActionLevel:    input.Role.CourseActionLevel,
			}
		}
	}

	for _, role := range roleMap {
		roleResponse = append(roleResponse, *role)
	}

	sort.Slice(roleResponse, func(i, j int) bool {
		firstID := pointer.GetUint(roleResponse[i].ID)
		secondID := pointer.GetUint(roleResponse[j].ID)
		return firstID < secondID
	})

	result = &dto.GetPossibleRolesResponseDto{
		Items: roleResponse,
	}

	return
}
