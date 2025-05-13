package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type qf4Handler struct {
	qf4Usecase domain.QF4Usecase
}

func NewTemplateHandler(qf4Route fiber.Router, QF4Usecase domain.QF4Usecase) {

	handler := &qf4Handler{
		qf4Usecase: QF4Usecase,
	}
	qf4Route.Get("/main", handler.GetQF4Main())
	qf4Route.Get("/course-plan", handler.GetQF4CoursePlan())
	qf4Route.Get("/assessment", handler.GetQF4Assessment())
	qf4Route.Get("/result", handler.GetQF4Result())
	qf4Route.Get("/course-info", handler.GetQF4CourseInfo())
	qf4Route.Get("/lecturer", handler.GetQF4Lecturer())
	qf4Route.Post("/main/create", handler.CreateQF4Main())
	qf4Route.Post("/course-plan/create", handler.CreateQF4CoursePlan())
	qf4Route.Post("/assessment/create", handler.CreateQF4Assessment())
	qf4Route.Post("/result/create", handler.CreateQF4Result())
	qf4Route.Post("/course-info/create", handler.CreateQF4CourseInfo())
	qf4Route.Post("/lecturer/create", handler.CreateQF4Lecturer())
	qf4Route.Patch("/main/", handler.UpdateQF4Main())
	qf4Route.Patch("/course-plan/", handler.CreateQF4CoursePlan())
	qf4Route.Patch("/assessment/", handler.UpdateQF4Assessment())
	qf4Route.Patch("/result/", handler.UpdateQF4Result())
	qf4Route.Patch("/course-info/", handler.UpdateQF4CourseInfo())
	qf4Route.Patch("/lecturer/", handler.UpdateQF4Lecturer())
	qf4Route.Delete("/main/delete", handler.DeleteQF4Main())
}
