package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	// migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

// TODO: create the new func don't forget to embed in domain use case interface
func (u *userUsecase) UpdateRoleUser(userID string, patchRequest *models.UpdateUserRoleRequest) error {
	err := u.userRepo.UpdateRoleUser(userID, *patchRequest)
	if err != nil {
		return err
	}

	newUserData := migrateModels.Users{
		SystemPermissionUID: patchRequest.SystemPermissionUID,
	}
	err = u.userRepo.UpdateUser(userID, newUserData)
	if err != nil {
		return err
	}
	// Delete all mapping for a user
	u.userRepo.DeleteMapCourseByUID(userID)

	for _, course := range patchRequest.RoleCourseUID {
		newMapUserCourseData := migrateModels.MapUserCourse{
			UserUID:   patchRequest.UserUID,
			CourseUID: course,
			RoleUID:   patchRequest.UserRoleUID,
		}
		err = u.userRepo.CreateMapCourse(newMapUserCourseData) // Create New
		if err != nil {
			return err
		}
	}
	// Delete all mapping for a user
	u.userRepo.DeleteMapProgramByUID(userID)

	for _, program := range patchRequest.RoleProgramUID {
		newMapUserProgramData := migrateModels.MapUserProgram{
			UserUID:    patchRequest.UserUID,
			ProgramUID: program,
			RoleUID:    patchRequest.UserRoleUID,
		}

		err = u.userRepo.CreateMapProgram(newMapUserProgramData) // Create New
		if err != nil {
			return err
		}
	}

	return nil
}
