package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (u *PermissionUsecase) GetFacultyRolesByUid(uid uuid.UUID, options models.PaginationOptions) (result *dto.GetMapFacultyRolesResponseDto, err error) {
	mapFacultyStatement := query.MapFacultiesRolesQueryEntity{
		UserID: &uid,
	}

	mapFacultyQuery := make([]query.MapFacultiesRolesQueryEntity, 0)

	err = u.CommonRepository.GetList(mapFacultyStatement, &mapFacultyQuery, &options, "Role", "Faculties")

	if err != nil {
		return nil, err
	}

	mapFacultyResult := make([]dto.MapFacultyRoles, 0)

	for _, mapFaculty := range mapFacultyQuery {
		mapFacultyResult = append(mapFacultyResult, dto.MapFacultyRoles{
			ID:            mapFaculty.ID,
			UserID:        mapFaculty.UserID,
			RoleID:        mapFaculty.RoleID,
			RoleNameTH:    mapFaculty.Role.RoleNameTH,
			RoleNameEN:    mapFaculty.Role.RoleNameEN,
			FacultyNameTH: mapFaculty.Faculty.FacultyNameTH,
			FacultyNameEN: mapFaculty.Faculty.FacultyNameEN,
			FacultyID:     mapFaculty.FacultyID,
		})
	}

	result = &dto.GetMapFacultyRolesResponseDto{
		Items:             mapFacultyResult,
		PaginationOptions: &options,
	}
	return result, nil
}
