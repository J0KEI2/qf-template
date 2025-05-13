package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type courseHandler struct {
	courseUsecase domain.CourseUsecase
}

func NewCourseHandler(courseRoute fiber.Router, courseUsecase domain.CourseUsecase) {

	handler := &courseHandler{
		courseUsecase: courseUsecase,
	}

	courseRoute.Get("/main", handler.GetCoursePagination())
	courseRoute.Get("/main/:uid", handler.GetCourseMain())
	courseRoute.Get("/course-plan", handler.GetCoursePlan())
	courseRoute.Get("/assessment", handler.GetCourseAssessment())
	courseRoute.Get("/result", handler.GetCourseResult())
	courseRoute.Get("/course-info", handler.GetCourseInfo())
	courseRoute.Get("/lecturer", handler.GetCourseLecturer())

	courseRoute.Post("/main/create", handler.CreateCourseMain())
	courseRoute.Post("/course-plan/create", handler.CreateCoursePlan())
	courseRoute.Post("/assessment/create", handler.CreateCourseAssessment())
	courseRoute.Post("/result/create", handler.CreateCourseResult())
	courseRoute.Post("/course-info/create", handler.CreateCourseInfo())
	courseRoute.Post("/lecturer/create", handler.CreateCourseLecturer())
	courseRoute.Patch("/main/", handler.UpdateCourseMain())
	courseRoute.Patch("/course-plan/", handler.CreateCoursePlan())
	courseRoute.Patch("/assessment/", handler.UpdateCourseAssessment())
	courseRoute.Patch("/result/", handler.UpdateCourseResult())
	courseRoute.Patch("/course-info/", handler.UpdateCourseInfo())
	courseRoute.Patch("/lecturer/", handler.UpdateCourseLecturer())
	courseRoute.Delete("/main/delete", handler.DeleteCourseMain())

}
