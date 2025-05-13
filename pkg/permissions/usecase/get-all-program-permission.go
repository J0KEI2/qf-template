package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	userQuery "github.com/zercle/kku-qf-services/pkg/models/query-model"
	permissionQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u *PermissionUsecase) GetAllProgramPermissionByUser(uid uuid.UUID) (result []dto.PermissionProgramResponseDto, err error) {

	userStatement := userQuery.UserQueryEntity{
		UID: &uid,
	}

	err = u.CommonRepository.GetFirst(&userStatement)
	if err != nil {
		return nil, err
	}

	queryStatement := permissionQuery.PermissionProgramQueryEntity{
		UserUID: &uid,
	}

	queryResults := []permissionQuery.PermissionProgramQueryEntity{}

	err = u.CommonRepository.GetList(queryStatement, &queryResults, nil)
	if err != nil {
		return nil, err
	}

	result = make([]dto.PermissionProgramResponseDto, 0)

	for _, programPermission := range queryResults {
		program := programQuery.ProgramMainQueryEntity{
			ID: programPermission.ProgramUID,
		}

		err := u.CommonRepository.GetFirst(&program, "ProgramGeneralDetail.Faculty")
		if err != nil {
			continue
		}

		result = append(result, dto.PermissionProgramResponseDto{
			ID:            programPermission.ID,
			UserUID:       programPermission.UserUID,
			ProgramUID:    programPermission.ProgramUID,
			Accessibility: programPermission.Accessibility,
			FacultyName:   program.ProgramGeneralDetail.Faculty.FacultyNameTH,
			ProgramName:   program.ProgramGeneralDetail.ProgramNameTH,
			ProgramType:   program.ProgramGeneralDetail.ProgramType,
			BranchName:    program.ProgramGeneralDetail.BranchNameTH,
			Email:         userStatement.Email,
			Name:          userStatement.NameTH,
		})
	}

	return result, nil
}
func (u *PermissionUsecase) GetAllProgramPermissionByProgram(uid uuid.UUID) (result []dto.PermissionProgramResponseDto, err error) {

	program := programQuery.ProgramMainQueryEntity{
		ID: &uid,
	}

	err = u.CommonRepository.GetFirst(&program, "ProgramGeneralDetail.Faculty")
	if err != nil {
		return nil, err
	}

	queryStatement := permissionQuery.PermissionProgramQueryEntity{
		ProgramUID: &uid,
	}

	queryResults := []permissionQuery.PermissionProgramQueryEntity{}

	err = u.CommonRepository.GetList(queryStatement, &queryResults, nil)
	if err != nil {
		return nil, err
	}

	result = make([]dto.PermissionProgramResponseDto, 0)

	for _, programPermission := range queryResults {
		userStatement := userQuery.UserQueryEntity{
			UID: programPermission.UserUID,
		}

		err = u.CommonRepository.GetFirst(&userStatement)
		if err != nil {
			continue
		}

		result = append(result, dto.PermissionProgramResponseDto{
			ID:            programPermission.ID,
			UserUID:       programPermission.UserUID,
			ProgramUID:    programPermission.ProgramUID,
			Accessibility: programPermission.Accessibility,
			FacultyName:   program.ProgramGeneralDetail.Faculty.FacultyNameTH,
			ProgramName:   program.ProgramGeneralDetail.ProgramNameTH,
			ProgramType:   program.ProgramGeneralDetail.ProgramType,
			BranchName:    program.ProgramGeneralDetail.BranchNameTH,
			Email:         userStatement.Email,
			Name:          userStatement.NameTH,
		})
	}

	return result, nil
}
