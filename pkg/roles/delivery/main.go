package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type roleHandler struct {
	roleUseCase domain.RoleUseCase
}

func NewRoleHandler(templateRoute fiber.Router, roleUseCase domain.RoleUseCase) {

	handler := &roleHandler{
		roleUseCase: roleUseCase,
	}

	templateRoute.Get("/possible-role", handler.GetPossibleRole())
	templateRoute.Get("/:roleID", handler.GetRoleByID())
	templateRoute.Get("/setting/cores/user-roles", handler.GetSettingUserCoreRole())
	templateRoute.Get("/setting/faculties/user-roles", handler.GetSettingUserFacultyRole())

	templateRoute.Post("/", handler.CreateOrUpdateRole())
	templateRoute.Post("/setting/cores", handler.CreateSettingUserCoreRole())        // Map core role
	templateRoute.Post("/setting/faculties", handler.CreateSettingUserFacultyRole()) // Map faculty role

	templateRoute.Delete("/:roleID", handler.DeleteRoleByID())
	templateRoute.Delete("/setting/cores/:userID/:roleID", handler.DeleteSettingUserCoreRole())
	templateRoute.Delete("/setting/faculties/:userID/:roleID/:facultyID", handler.DeleteSettingUserFacultyRole())
}
