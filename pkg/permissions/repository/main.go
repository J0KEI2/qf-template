package repository

import (
	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/permission"
	"github.com/zercle/kku-qf-services/pkg/domain"
	model "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/permission"
	"github.com/zercle/kku-qf-services/pkg/utils"
	"gorm.io/gorm"
)

type permissionRepository struct {
	MainDbConn *gorm.DB
}

func NewPermissionRepository(mainDbConn *gorm.DB) domain.PermissionRepository {
	return &permissionRepository{
		MainDbConn: mainDbConn,
	}
}

func (repo *permissionRepository) DbPermissionSVCMigrator() (err error) {
	if err = repo.MainDbConn.AutoMigrate(&model.PermissionSystem{},
		&model.PermissionCourse{},
		&model.PermissionProgram{},
		// Will delete went create new CRUD
		&model.Role{},
		&model.MapCoreRoles{},
		&model.MapFacultiesRoles{},
		&model.MapProgramsRoles{},
		&model.ActionLog{},
	); err != nil {
		return err
	}

	if err = initialPermissionData(repo); err != nil {
		return
	}

	return
}

func initialPermissionData(repo *permissionRepository) (err error) {
	permission := utils.StringifyStringArray([]string{constant.PERMISSION_NONE})
	pagePermission := utils.StringifyStringArray([]string{constant.PAGE_HOME})
	anonymousRole := &query.PermissionSystemQueryEntity{
		UID:                       &uuid.UUID{},
		RoleNameTH:                pointer.ToString("ยังไม่ระบุตำแหน่ง"),
		RoleNameEN:                pointer.ToString("anonymous"),
		PageAccessibility:         &pagePermission,
		ProgramAccessibility:      &permission,
		CourseAccessibility:       &permission,
		ProgramAccessibilityLevel: pointer.ToUint(0),
		CourseAccessibilityLevel:  pointer.ToUint(0),
		UAMControl:                false,
		CanComment:                false,
		CanApproved:               false,
	}
	if err = repo.MainDbConn.FirstOrCreate(&anonymousRole).Error; err != nil {
		return err
	}
	return nil
}
