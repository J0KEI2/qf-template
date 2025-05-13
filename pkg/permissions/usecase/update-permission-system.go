package usecase

import (
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"github.com/zercle/kku-qf-services/pkg/utils"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) UpdateSystemPermission(uid uuid.UUID, update dto.UpdatePermissionSystemDto) (result *dto.PermissionSystemResponseDto, err error) {

	queryDB := query.PermissionSystemQueryEntity{
		UID: &uid,
	}

	updateBody := query.PermissionSystemQueryEntity{
		RoleNameTH:                update.RoleNameTH,
		RoleNameEN:                update.RoleNameEN,
		ProgramAccessibilityLevel: update.ProgramAccessibilityLevel,
		CourseAccessibilityLevel:  update.CourseAccessibilityLevel,
		UAMControl:                update.UAMControl,
		CanComment:                update.CanComment,
		CanApproved:               update.CanApproved,
	}

	if update.PageAccessibility != nil {
		*updateBody.PageAccessibility = utils.StringifyStringArray(update.PageAccessibility)
	}
	if update.ProgramAccessibility != nil {
		*updateBody.ProgramAccessibility = utils.StringifyStringArray(update.ProgramAccessibility)
	}
	if update.CourseAccessibility != nil {
		*updateBody.CourseAccessibility = utils.StringifyStringArray(update.CourseAccessibility)
	}

	helper.ExecuteTransaction(u.CommonRepository, updatePermissionStatement(u, queryDB, &updateBody))

	result, err = u.GetOneSystemPermission(uid)

	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	return result, nil
}

func updatePermissionStatement(usecase *PermissionUsecase, queryDb query.PermissionSystemQueryEntity, updateBody *query.PermissionSystemQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return usecase.CommonRepository.Update(tx, queryDb, updateBody)
	}
}
