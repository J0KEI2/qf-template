package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/domain"
)

type programHandler struct {
	programUseCase domain.ProgramUsecase
}

func NewProgramHandler(programRoute fiber.Router, mdwUC domain.MiddlewaresUseCase, programUseCase domain.ProgramUsecase, hrUseCase domain.HRUseCase) {

	handler := &programHandler{
		programUseCase: programUseCase,
	}

	// Get
	programRoute.Get("/main/", mdwUC.ProgramActionAuth(1), handler.GetProgramPagination())
	programRoute.Get("/main/:id", mdwUC.ProgramActionAuth(1), handler.GetProgram())
	programRoute.Get("/structure-course/:subPlanID", mdwUC.ProgramActionAuth(1), handler.GetCourseStructure())
	programRoute.Get("/year-course/:subPlanID", mdwUC.ProgramActionAuth(1), handler.GetYearCourse())
	programRoute.Get("/plo/:subPlanID", mdwUC.ProgramActionAuth(1), handler.GetPloKsec())
	programRoute.Get("/plo/:subPlanID/map-ksec/:ploID", mdwUC.ProgramActionAuth(1), handler.GetMapPloKsec())
	programRoute.Get("/plo/:subPlanID/map-solution/:ploID", mdwUC.ProgramActionAuth(1), handler.GetMapPloLearningSolution())
	programRoute.Get("/plo/:subPlanID/map-evaluation/:ploID", mdwUC.ProgramActionAuth(1), handler.GetMapPloLearningEvaluation())
	programRoute.Get("/lecturer/:uuid", mdwUC.ProgramActionAuth(1), handler.GetLecturerOwner())
	programRoute.Get("/quality-assurance/:uuid", mdwUC.ProgramActionAuth(1), handler.GetQualityAssurance())
	programRoute.Get("/system-mechanic/:uuid", mdwUC.ProgramActionAuth(1), handler.GetSystemMechanic())
	programRoute.Get("/reference/:uuid", mdwUC.ProgramActionAuth(1), handler.GetReference())
	programRoute.Get("/clo/:planDetailID", mdwUC.ProgramActionAuth(1), handler.GetYearCourse())
	programRoute.Get("/general-detail/:uuid", mdwUC.ProgramActionAuth(1), handler.GetGeneralDetail())
	programRoute.Get("/course/", mdwUC.ProgramActionAuth(1), handler.GetCourseForStructure())
	programRoute.Get("/course/plan/:subPlanID", mdwUC.ProgramActionAuth(1), handler.GetCourseDetails())
	programRoute.Get("/policy-and-strategic/:uuid", mdwUC.ProgramActionAuth(1), handler.GetPolicyAndStrategic())
	programRoute.Get("/competency/:uuid", mdwUC.ProgramActionAuth(1), handler.GetCompetency())
	programRoute.Get("/major/:uuid", mdwUC.ProgramActionAuth(1), handler.GetMajorAndPlan())
	programRoute.Get("/ylo-detail/:subPlanID", mdwUC.ProgramActionAuth(1), handler.GetYLODetail())
	programRoute.Get("/plan-and-evaluate/:subPlanID", mdwUC.ProgramActionAuth(1), handler.GetPlanAndEvaluate())
	programRoute.Get("/curriculum-mapping/:subPlanID/type/:curMapType", mdwUC.ProgramActionAuth(1), handler.GetCurriculumMapping())
	programRoute.Get("/ksa-detail/:subPlanID", mdwUC.ProgramActionAuth(1), handler.GetKSADetail())
	programRoute.Get(":subPlanID/structure/", mdwUC.ProgramActionAuth(1), handler.GetStructureBySubPlan())
	programRoute.Get(":subPlanID/structure/:structureId", mdwUC.ProgramActionAuth(1), handler.GetStructure())
	programRoute.Get(":subPlanID/structure/:structureId/courses", mdwUC.ProgramActionAuth(1), handler.GetCourseByStructureId())
	programRoute.Get("/check-openable-page/:uuid", mdwUC.ProgramActionAuth(1), handler.CheckOpenablePage())
	programRoute.Get("/download/documents/:uuid/:fileName", handler.DownloadDocumentFile())
	programRoute.Get("/upload/general-detail/documents/mou/:general_detail_id", handler.GetGeneralDetailMou())

	// Post
	programRoute.Post("/main", mdwUC.ProgramActionAuth(2), handler.CreateProgramMain())
	programRoute.Post("/general-detail", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateGeneralDetail())
	programRoute.Post("/structure-course/:subPlanID", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateCourseStructure())
	programRoute.Post("/year-course/:subPlanID", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateYearCourse())
	programRoute.Post("/plo/:subPlanID", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdatePlo())
	programRoute.Post("/plo/:subPlanID/map-ksec/:ploID", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateMapPloWithKsec())
	programRoute.Post("/plo/:subPlanID/map-solution/:ploID", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateLearningSolution())
	programRoute.Post("/plo/:subPlanID/map-evaluation/:ploID", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateLearningEvaluation())
	programRoute.Post("/lecturer/:uuid", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateLecturerOwner())
	programRoute.Post("/quality-assurance/:uuid", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateQualityAssurance())
	programRoute.Post("/system-mechanic/:uuid", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateSystemMechanic())
	programRoute.Post("/reference/:uuid", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateReference())
	programRoute.Post("/policy-and-strategic", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdatePolicyAndStrategic())
	programRoute.Post("/competency", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateCompetency())
	programRoute.Post("/ylo-detail", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateYLODetail())
	programRoute.Post("/plan-and-evaluate", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdatePlanAndEvaluate())
	programRoute.Post("/curriculum-mapping/type/:curMapType", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateCurriculumMapping())
	programRoute.Post("/ksa-detail", mdwUC.ProgramActionAuth(2), handler.CreateOrUpdateKsaDetail())
	programRoute.Post("/duplicate", mdwUC.ProgramActionAuth(2), handler.DuplicateProgram())
	programRoute.Post(":subPlanID/structure/", mdwUC.ProgramActionAuth(2), handler.CreateStructure())
	programRoute.Post(":subPlanID/structure/:structureId/courses", mdwUC.ProgramActionAuth(2), handler.CreateCourse())
	programRoute.Post(":subPlanID/year-and-semester/", mdwUC.ProgramActionAuth(2), handler.CreateStructure())
	programRoute.Post(":subPlanID/year-and-semester/:yearAndSemesterId", mdwUC.ProgramActionAuth(2), handler.CreateStructure())
	programRoute.Post("/upload/general-detail/documents/mou", handler.UploadGeneralDetailMouDocumentsFile())
	programRoute.Post("/upload/reference/documents", handler.UploadReferenceDocumentsFile())

	// Patch
	programRoute.Patch("/course/:id", mdwUC.ProgramActionAuth(2), handler.UpdateCourse())
	programRoute.Patch("/structure/:id", mdwUC.ProgramActionAuth(2), handler.UpdateCourse())
	programRoute.Patch("/year-and-semester/:id", mdwUC.ProgramActionAuth(2), handler.UpdateCourse())
	programRoute.Patch("/plo/:id", mdwUC.ProgramActionAuth(2), handler.UpdatePlo())
	programRoute.Patch("/ksec/:id", mdwUC.ProgramActionAuth(2), handler.UpdateKsec())
	programRoute.Patch("/map-plo-ksec/:id", mdwUC.ProgramActionAuth(2), handler.UpdateMapPloKsec())
	programRoute.Patch("/learning-evaluation/:id", mdwUC.ProgramActionAuth(2), handler.UpdateLearningEvaluation())
	programRoute.Patch("/learning-solution/:id", mdwUC.ProgramActionAuth(2), handler.UpdateLearningSolution())

	// Delete
	programRoute.Delete("/main/:id", mdwUC.ProgramActionAuth(3), handler.DeleteMain())
	programRoute.Delete("/major-and-plan/", mdwUC.ProgramActionAuth(3), handler.DeleteMajorAndPlan())
	programRoute.Delete("/competency/:id", mdwUC.ProgramActionAuth(3), handler.DeleteCompetency())
	programRoute.Delete("/structure/:id", mdwUC.ProgramActionAuth(3), handler.DeleteStructure())
	programRoute.Delete("/course/:id", mdwUC.ProgramActionAuth(3), handler.DeleteCourse())
	programRoute.Delete("/year-and-semester/:id", mdwUC.ProgramActionAuth(3), handler.DeleteYearAndSemester())
	programRoute.Delete("/year-course/:id", mdwUC.ProgramActionAuth(3), handler.DeleteYearCourse())
	programRoute.Delete("/plo/:id", mdwUC.ProgramActionAuth(3), handler.DeletePlo())
	programRoute.Delete("/ksec/:id", mdwUC.ProgramActionAuth(3), handler.DeleteKsec())
	programRoute.Delete("/map-plo-ksec/:id", mdwUC.ProgramActionAuth(3), handler.DeleteMapPloKsec())
	programRoute.Delete("/learning-evaluation/:id", mdwUC.ProgramActionAuth(3), handler.DeleteLearningEvaluation())
	programRoute.Delete("/learning-solution/:id", mdwUC.ProgramActionAuth(3), handler.DeleteLearningSolution())
	programRoute.Delete("/lecturer-owner/:id", mdwUC.ProgramActionAuth(3), handler.DeleteLecturerOwner())
	programRoute.Delete("/lecturer-thesis/:id", mdwUC.ProgramActionAuth(3), handler.DeleteLecturerThesis())
	programRoute.Delete("/lecturer/:id", mdwUC.ProgramActionAuth(3), handler.DeleteLecturer())
	programRoute.Delete("/reference/:id", mdwUC.ProgramActionAuth(3), handler.DeleteReference())
	programRoute.Delete("/map-file-system/:id", mdwUC.ProgramActionAuth(3), handler.DeleteMapFileSystem())
}
