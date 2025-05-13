package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type commentHandler struct {
	CommentUseCase domain.CommentUseCase
}

func NewcommentHandler(router fiber.Router, commentUsecase domain.CommentUseCase, mdwUC domain.MiddlewaresUseCase) {

	handler := &commentHandler{
		CommentUseCase: commentUsecase,
	}

	router.Get("/:qf_type/:qf_uid/:category_type", mdwUC.ProgramApprovalAuth(1), handler.GetComments())
	router.Get("/:id", mdwUC.ProgramApprovalAuth(1), handler.GetComment())

	router.Post("/", mdwUC.ProgramApprovalAuth(2), handler.CreateComment())

	router.Patch("/:id", mdwUC.ProgramApprovalAuth(2), handler.UpdateComment())
	router.Patch("/:id/resolve", mdwUC.ProgramApprovalAuth(2), handler.ResolveComment())
}
