package usecase

import (
	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

// TODO: create the new func don't forget to embed in domain use case interface
func (u programUsecase) GetMainProgram(userUID, id uuid.UUID) (result *dto.ProgramMainPagination, err error) {
	userStatement := query.UserQueryEntity{
		UID: &userUID,
	}

	err = u.CommonRepository.GetFirst(&userStatement, "SystemPermission")
	if err != nil {
		return nil, err
	}
	record, err := u.ProgramRepository.GetMainProgram(userStatement, id)
	if err != nil {
		return nil, err
	}

	Qf2MainDetail := dto.ProgramMainPagination{
		ProgramMainID:   record.ID,
		PermissionLevel: pointer.ToUint(uint(3)),
	}
	if detail := record.ProgramGeneralDetail; detail != nil {
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
	}

	return &Qf2MainDetail, nil
}

// func (u *userUsecase) GetUser(userID string) (user models.User, err error) {
// 	return u.userRepo.GetUser(userID)
// }
