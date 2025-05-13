package usecase

import (
	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"github.com/zercle/kku-qf-services/pkg/utils"
)

func (u *PermissionUsecase) GetOneSystemPermission(uid uuid.UUID) (result *dto.PermissionSystemResponseDto, err error) {

	queryStatement := query.PermissionSystemQueryEntity{
		UID: &uid,
	}

	u.CommonRepository.GetFirst(&queryStatement)

	result = &dto.PermissionSystemResponseDto{
		UID:                       queryStatement.UID.String(),
		RoleNameTH:                queryStatement.RoleNameTH,
		RoleNameEN:                queryStatement.RoleNameEN,
		PageAccessibility:         utils.ParseStringArray(pointer.GetString(queryStatement.PageAccessibility)),
		ProgramAccessibility:      utils.ParseStringArray(pointer.GetString(queryStatement.ProgramAccessibility)),
		CourseAccessibility:       utils.ParseStringArray(pointer.GetString(queryStatement.CourseAccessibility)),
		ProgramAccessibilityLevel: queryStatement.ProgramAccessibilityLevel,
		CourseAccessibilityLevel:  queryStatement.CourseAccessibilityLevel,
		UAMControl:                queryStatement.UAMControl,
		CanComment:                queryStatement.CanComment,
		CanApprove:                queryStatement.CanApproved,
	}

	return result, nil
}

func (u *PermissionUsecase) GetAllSystemPermission() (result []dto.PermissionSystemResponseDto, err error) {

	queryStatement := query.PermissionSystemQueryEntity{}

	queryResults := []query.PermissionSystemQueryEntity{}

	u.CommonRepository.GetList(queryStatement, &queryResults, nil)

	for _, permissionSystem := range queryResults {
		result = append(result, dto.PermissionSystemResponseDto{
			UID:                       permissionSystem.UID.String(),
			RoleNameTH:                permissionSystem.RoleNameTH,
			RoleNameEN:                permissionSystem.RoleNameEN,
			PageAccessibility:         utils.ParseStringArray(pointer.GetString(permissionSystem.PageAccessibility)),
			ProgramAccessibility:      utils.ParseStringArray(pointer.GetString(permissionSystem.ProgramAccessibility)),
			CourseAccessibility:       utils.ParseStringArray(pointer.GetString(permissionSystem.CourseAccessibility)),
			ProgramAccessibilityLevel: permissionSystem.ProgramAccessibilityLevel,
			CourseAccessibilityLevel:  permissionSystem.CourseAccessibilityLevel,
			UAMControl:                permissionSystem.UAMControl,
			CanComment:                permissionSystem.CanComment,
			CanApprove:                permissionSystem.CanApproved,
		})
	}

	return result, nil
}
