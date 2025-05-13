package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type userDetailHandler struct {
	userDetailUseCase domain.UserDetailUseCase
	hrUseCase         domain.HRUseCase
}

func NewUserDetailHandler(userDetailRoute fiber.Router, userDetailCase domain.UserDetailUseCase, hrUsecase domain.HRUseCase) {

	handler := &userDetailHandler{
		userDetailUseCase: userDetailCase,
		hrUseCase:         hrUsecase,
	}

	userDetailRoute.Get("/searchByName", handler.SearchLecturerByName())
	userDetailRoute.Get("/cron/update-userDetails", handler.CronUpdateLecturers())
	userDetailRoute.Get("/user-detail/", handler.GetUserDetailPagination())
}
