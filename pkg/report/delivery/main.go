package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type reportHandler struct {
	reportUseCase domain.ReportUsecase
}

func NewReportHandler(pdfRoute fiber.Router, pdfUsecase domain.ReportUsecase) {
	handler := &reportHandler{
		reportUseCase: pdfUsecase,
	}
	pdfRoute.Get("/export", handler.ExportReport())
	pdfRoute.Get("/report/:uuid", handler.GetReport())

	pdfRoute.Post("/upload/report", handler.CreateOrUpdateReport())

	pdfRoute.Delete("/report/:report_id", handler.DeleteReport())

}
