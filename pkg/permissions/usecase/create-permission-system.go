package usecase

import (
	"log"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"github.com/zercle/kku-qf-services/pkg/utils"
	"gorm.io/gorm"
)

func (u *PermissionUsecase) CreateNewSystemPermission(request dto.CreatePermissionSystemDto) (result *dto.PermissionSystemResponseDto, err error) {

	systemPermissionUID := uuid.New()
	pageAccessibility := utils.StringifyStringArray(request.PageAccessibility)
	programAccessibility := utils.StringifyStringArray(request.ProgramAccessibility)
	courseAccessibility := utils.StringifyStringArray(request.CourseAccessibility)

	queryDb := query.PermissionSystemQueryEntity{
		UID:                       &systemPermissionUID,
		RoleNameTH:                &request.RoleNameTH,
		RoleNameEN:                &request.RoleNameEN,
		PageAccessibility:         &pageAccessibility,
		ProgramAccessibility:      &programAccessibility,
		CourseAccessibility:       &courseAccessibility,
		ProgramAccessibilityLevel: &request.ProgramAccessibilityLevel,
		CourseAccessibilityLevel:  &request.CourseAccessibilityLevel,
		UAMControl:                request.UAMControl,
		CanComment:                request.CanComment,
		CanApproved:               request.CanApproved,
	}

	if err = helper.ExecuteTransaction(u.CommonRepository, createPermissionStatement(u, &queryDb)); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	result = &dto.PermissionSystemResponseDto{
		UID:                       queryDb.UID.String(),
		RoleNameTH:                queryDb.RoleNameTH,
		RoleNameEN:                queryDb.RoleNameEN,
		PageAccessibility:         utils.ParseStringArray(pointer.GetString(queryDb.PageAccessibility)),
		ProgramAccessibility:      utils.ParseStringArray(pointer.GetString(queryDb.ProgramAccessibility)),
		CourseAccessibility:       utils.ParseStringArray(pointer.GetString(queryDb.CourseAccessibility)),
		ProgramAccessibilityLevel: queryDb.ProgramAccessibilityLevel,
		CourseAccessibilityLevel:  queryDb.CourseAccessibilityLevel,
		UAMControl:                queryDb.UAMControl,
		CanComment:                queryDb.CanComment,
		CanApprove:                queryDb.CanApproved,
	}

	return result, nil
}

func createPermissionStatement(usecase *PermissionUsecase, queryDb *query.PermissionSystemQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return usecase.CommonRepository.Create(tx, queryDb)
	}
}
