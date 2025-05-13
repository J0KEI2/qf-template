package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type commonHandler struct {
	commonUseCase domain.CommonUseCase
}

func NewCommonHandler(commonRoute fiber.Router, commonUsecase domain.CommonUseCase) {
	handler := &commonHandler{
		commonUseCase: commonUsecase,
	}
	commonRoute.Get("/faculties", handler.GetFacultiesPagination())
	commonRoute.Get("/faculties/all", handler.GetAllFaculties())
	commonRoute.Get("/faculties/:id", handler.GetFacultyByID())
	commonRoute.Get("/references-option/all", handler.GetAllReferenceOption())

	// Download document
	commonRoute.Get("/program/documents/:file_id", handler.GetFileByID())
}
