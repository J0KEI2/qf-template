package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type userHandler struct {
	UserUsecase domain.UserUsecase
	HrUsecase   domain.HRUseCase
}

// NewPostHandler will initialize the post resource endpoint
func NewUserHandler(router fiber.Router, userUsecase domain.UserUsecase, hrUsecase domain.HRUseCase, mdwUC domain.MiddlewaresUseCase) {

	handler := &userHandler{
		UserUsecase: userUsecase,
		HrUsecase:   hrUsecase,
	}

	router.Get("/", handler.GetProfile())
	router.Get("/all", mdwUC.SystemAuth(1), handler.GetUserList())
	router.Get("/:id", mdwUC.SystemAuth(1), handler.GetUserByID())

	router.Post("/", mdwUC.SystemAuth(2), handler.CreateUser())

	router.Patch("/:uid", mdwUC.SystemAuth(2), handler.EditUser())
	router.Patch("/:id/role", mdwUC.SystemAuth(2), handler.UpdateRoleUser())

	router.Delete("/:id", mdwUC.SystemAuth(3), handler.DeleteUserById())
}
