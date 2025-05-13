package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type PermissionHandler struct {
	permissionUsecase domain.PermissionUseCase
}

func NewPermissionHandler(permissionRoute fiber.Router, permissionUsercase domain.PermissionUseCase) {

	permissionHandler := &PermissionHandler{
		permissionUsecase: permissionUsercase,
	}

	permissionRoute.Get("core/", permissionHandler.GetUserCoreRoles())
	permissionRoute.Get("core/:uid", permissionHandler.GetUserCoreRolesByUid())
	permissionRoute.Post("core/", permissionHandler.CreateUserCoreRoles())
	permissionRoute.Delete("core/:id", permissionHandler.DeleteUserCoreRoles())

	permissionRoute.Get("faculty/", permissionHandler.GetUserFacultyRoles())
	permissionRoute.Get("faculty/:uid", permissionHandler.GetUserFacultyRolesByUid())
	permissionRoute.Post("faculty/", permissionHandler.CreateUserFacultyRoles())
	permissionRoute.Delete("faculty/:id", permissionHandler.DeleteUserFacultyRoles())

	permissionRoute.Get("user/program", permissionHandler.GetUserProgramRoles())
	permissionRoute.Get("user/program/:uid", permissionHandler.GetUserProgramRolesByUid())
	permissionRoute.Post("user/program", permissionHandler.CreateUserProgramRoles())
	permissionRoute.Delete("user/program/:id", permissionHandler.DeleteUserProgramRoles())

	permissionRoute.Get("const/permissions", permissionHandler.GetPermissionConst())
	permissionRoute.Get("const/pages", permissionHandler.GetPageConst())

	permissionRoute.Get("system/", permissionHandler.GetAllPermissionSystem())
	permissionRoute.Get("system/:uid", permissionHandler.GetOnePermissionSystem())
	permissionRoute.Post("system/", permissionHandler.CreateNewPermissionSystem())
	permissionRoute.Patch("system/:uid", permissionHandler.UpdatePermissionSystem())
	permissionRoute.Delete("system/:uid", permissionHandler.DeletePermissionSystem())

	permissionRoute.Get("program/:uid", permissionHandler.GetProgramRoleByProgramUID())
	permissionRoute.Get("program/user/", permissionHandler.GetUserPermissionProgram())
	permissionRoute.Get("program/user/:uid", permissionHandler.GetPermissionProgramByUser())
	permissionRoute.Post("program/", permissionHandler.CreateOrUpdatePermissionProgram())
	permissionRoute.Delete("program/", permissionHandler.DeletePermissionProgram())
}
