package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u roleUsecase) CreateOrUpdateRole(request dto.CreateOrUpdateRoleRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateRoleTransaction(request))
}

func (u roleUsecase) CreateOrUpdateRoleTransaction(request dto.CreateOrUpdateRoleRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		roleQuery := query.RoleQueryEntity{
			ID: request.ID,
		}

		updateBody := query.RoleQueryEntity{
			RoleNameTH:           request.RoleNameTH,
			RoleNameEN:           request.RoleNameEN,
			ProgramRoleType:      request.ProgramRoleType,
			CourseRoleType:       request.CourseRoleType,
			SystemLevel:          request.SystemLevel,
			ProgramApprovalLevel: request.ProgramApprovalLevel,
			CourseApprovalLevel:  request.CourseApprovalLevel,
			ProgramActionLevel:   request.ProgramActionLevel,
			CourseActionLevel:    request.CourseActionLevel,
		}

		if roleQuery.ID != nil {
			if err = u.CommonRepository.Update(tx, &roleQuery, &updateBody); err != nil {
				return err
			}
		} else {
			if err = u.CommonRepository.Create(tx, &updateBody); err != nil {
				return err
			}
		}

		return nil
	}
}
