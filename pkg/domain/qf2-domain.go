package domain

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	courseQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	rapQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

type ProgramUsecase interface {
	GetMainProgram(userUID, id uuid.UUID) (result *dto.ProgramMainPagination, err error)
	GetStructure(structureId uint) (result *dto.ProgramStructureResponseDto, err error)
	GetStructureBySubPlan(subPlanID uint, paginationOptions *models.PaginationOptions) (result *dto.GetCourseStructureResponseDto, err error)
	GetCourseByStructureId(structureId uint, paginationOptions *models.PaginationOptions) (result *dto.GetCourseByStructureResponseDto, err error)
	GetStructureAndCourse(subPlanID uint, paginationOptions *models.PaginationOptions) (result *dto.GetCourseStructureResponseDto, err error)
	GetYearAndCourse(planDetailID uint, paginationOptions *models.PaginationOptions) (result *dto.GetYearCourseResponseDto, err error)
	GetPlo(ProgramSubPlanId uint) (result dto.ProgramPLOGetResponseDto, err error)
	GetPloMapWithKsec(ProgramSubPlanId uint, ploID uint) (ksec *dto.KsecResponseDto, err error)
	GetMapPloWithLearningSolution(ploID uint, paginationOptions *models.PaginationOptions) (result dto.LearningSolutionResponseDto, err error)
	GetMapPloWithLearningEvaluation(ploID uint, paginationOptions *models.PaginationOptions) (learningEvaluationList dto.LearningEvaluationResponseDto, err error)
	GetLecturerOwner(ProgramID uuid.UUID) (result *dto.GetLecturerOwnerDto, err error)
	GetQualityAssurance(ProgramID uuid.UUID) (result *dto.ProgramQualityAssurance, err error)
	GetSystemMechanic(ProgramID uuid.UUID) (result *dto.ProgramSystemAndMechanicDto, err error)
	GetReference(ProgramID uuid.UUID, paginationOptions *models.PaginationOptions) (result dto.GetReferenceResponseDto, err error)
	CreateStructure(request dto.ProgramStructureRequestDto, subPlanID uint) (result *dto.ProgramStructureResponseDto, err error)
	CreateCourse(request dto.ProgramCourseDetailRequestDto, subPlanID uint, structureId uint) (result *dto.ProgramCourseDetailResponseDto, err error)
	CreateYearAndSemester(request dto.ProgramYearAndSemesterRequestDto, subPlanID uint) (result *dto.ProgramYearAndSemesterResponseDto, err error)
	CreateOrUpdateCourseStructure(request dto.CreateOrUpdateCourseStructureRequestDto, planDetailID int) (err error)
	CreateOrUpdateYearCourse(years []dto.ProgramYearRequestDto, planDetailID int) (err error)
	CreateOrUpdatePlo(plo dto.CreateOrUpdatePLORequestDto, subPlanID uint) (err error)
	CreateOrUpdateLecturerOwner(lecturerOwner dto.CreateOrUpdateLecturerOwnerDto, programMainUID uuid.UUID) (err error)
	CreateOrUpdateQualityAssurance(qa dto.CreateOrUpdateQualityAssuranceDto, programMainUID uuid.UUID) (err error)
	CreateOrUpdateSystemMechanic(systemAndMechanic dto.CreateOrUpdateSystemAndMechanicDto, programMainUID uuid.UUID) (err error)
	CreateOrUpdateReference(reference dto.ProgramReferenceDto, programMainUID uuid.UUID, file *multipart.FileHeader, c *fiber.Ctx, userUID uuid.UUID) (err error)
	CreateOrUpdateGeneralDetail(generalDetail dto.ProgramGeneralDetailRequestDto) (err error)
	GetGeneralDetail(ProgramID uuid.UUID) (result *dto.ProgramGeneralDetailGetResponseDto, err error)
	GetCoursePaginationForStructure(paginationOptions models.PaginationOptions) (result *dto.GetCoursePaginationForStructureResponseDto, err error)
	GetCourseDetailPagination(paginationOptions models.PaginationOptions, planID int) (result *dto.GetCourseDetailPaginationResponseDto, err error)
	CreateOrUpdatePolicyAndStrategic(policyAndStrategic dto.ProgramPolicyAndStrategicRequestDto) (err error)
	GetPolicyAndStrategic(ProgramID uuid.UUID) (result *dto.ProgramPolicyAndStrategicGetResponseDto, err error)
	CreateOrUpdateCompetency(competency dto.ProgramCompetencyRequestDto) (err error)
	GetCompetency(ProgramID uuid.UUID, paginationOptions *models.PaginationOptions) (result *dto.ProgramCompetencyResponseDto, err error)
	DeleteCompetency(competencyID uint) (err error)
	GetMajorAndPlan(ProgramID uuid.UUID) (result *dto.ProgramMajorAndPLanGetResponseDto, err error)
	CreateOrUpdateYLODetail(yloDetail dto.CreateOrUpdateYLODetailRequestDto) (err error)
	GetYLODetail(programSubPlanID uint) (result dto.ProgramYLODetailGetResponseDto, err error)
	GetMainProgramPagination(userUID uuid.UUID, roleID *uint, options models.PaginationOptions, param dto.GetMainProgramPaginationQueryParam) (result *dto.GetMainProgramPaginationResponseDto, err error)
	CreateOrUpdatePlanAndEvaluate(planAndEvaluate dto.ProgramPlanAndEvaluateRequestDto) (err error)
	GetPlanAndEvaluate(programSubPlanID uint) (result *dto.ProgramPlanAndEvaluateResponseDto, err error)
	CreateOrUpdateLearningEvaluation(LearningEvaluation []dto.CreateOrUpdateLearningEvaluationRequestDto, learningSolutionID uint) (err error)
	CreateOrUpdatLearningSolution(learningSolutions []dto.CreateOrUpdateLearningSolutionRequestDto, learningSolutionID uint) (err error)
	CreateOrUpdateMapPloWithKsec(ksecRequest dto.KsecRequestDto, ploID uint) (err error)
	CreateYearAndSemesterByEducationYear(subPlanId, academicYear uint) error
	GetCurriculumMapping(paginationOptions models.PaginationOptions, subPlanID uint) (result *dto.ProgramCurMapRespGetResponseDto, err error)
	GetCurriculumMappingKsa(paginationOptions models.PaginationOptions, subPlanID uint) (result *dto.ProgramCurMapKsaGetResponseDto, err error)
	CreateOrUpdateCurriculumMappingResp(paginationOptions *models.PaginationOptions, request dto.CreateOrUpdateCurMapRespRequestDto) (err error)
	CreateOrUpdateCurriculumMappingKsa(paginationOptions *models.PaginationOptions, request dto.CreateOrUpdateCurMapKsaRequestDto) (err error)
	GetKSADetail(subPlanID uint) (result *dto.ProgramKsaDetailGetResponseDto, err error)
	CreateOrUpdateKsaDetail(ksaDetail dto.CreateOrUpdateKsaDetailRequestDto) (err error)
	CreateMainProgram(userUID uuid.UUID, request dto.ProgramMainRequestDto) (result *dto.ProgramMainPagination, err error)
	DuplicateProgram(userUID uuid.UUID, request dto.ProgramDuplicateRequestDto) (newProgramId *uuid.UUID, err error)
	UpdateCourse(request dto.ProgramCourseDetailRequestDto, courseId uint) (result *dto.ProgramCourseDetailResponseDto, err error)
	UpdateStructure(request dto.ProgramStructureRequestDto, structureId uint) (result *dto.ProgramStructureResponseDto, err error)
	UpdateYearAndSemester(request dto.ProgramYearAndSemesterRequestDto, yearAndSemesterId uint) (result *dto.ProgramYearAndSemesterResponseDto, err error)
	UpdateLearningSolution(request dto.LearningSolution, learningSolutionId uint) (result *dto.LearningSolution, err error)
	UpdateMapPloKsec(request dto.ProgramMapPloWithKsec, ksecId uint) (result *dto.ProgramMapPloWithKsec, err error)
	UpdateKsec(request dto.KsecDetail, ksecId uint) (result *dto.KsecDetail, err error)
	UpdatePlo(request dto.ProgramPLODetailDto, ploId uint) (result *dto.ProgramPLODetailDto, err error)
	UpdateLearningEvaluation(request dto.LearningEvaluation, learningEvaluationId uint) (result *dto.LearningEvaluation, err error)
	DeleteMain(id uuid.UUID) (err error)
	DeleteStructure(id uint) (err error)
	DeleteYearAndSemester(id uint) (err error)
	DeleteYearCourse(id uint) (err error)
	DeletePlo(id uint) (err error)
	DeleteMapPloKsec(id uint) (err error)
	DeleteLearningSolution(id uint) (err error)
	DeleteLearningEvaluation(id uint) (err error)
	DeleteKsec(id uint) (err error)
	DeleteCourse(id uint) (err error)
	DeleteMajorPlanSubPlan(request dto.ProgramMajorAndPlanDeleteRequest) (err error)
	CheckOpenablePage(ProgramID uuid.UUID) (result *dto.GetOpenablePageResponseDto, err error)
	DeleteLecturerOwner(id uint) (err error)
	DeleteLecturer(id uint) (err error)
	DeleteLecturerThesis(id uint) (err error)
	UploadGeneralDetailMouDocumentsFile(c *fiber.Ctx, programUID uuid.UUID, generalDetailId *uint, files []dto.MouFileDto, userUID uuid.UUID) ([]uint, error)
	UploadReferenceDocumentsFile(c *fiber.Ctx, programUID uuid.UUID, files []*multipart.FileHeader, userUID uuid.UUID) error
	DeleteReferences(id uint) (err error)
	CreateOrUpdateReferenceOption(reference dto.ProgramReferenceDto) (err error)
	DeleteMapFileSystem(id uint) (err error)
	GetGeneralDetailMou(generalDetailID uint) (result dto.GetMouFileResponseDto, err error)
}

type ProgramRepository interface {
	DbProgramSVCMigrator() (err error)
	GetMainProgram(user query.UserQueryEntity, id uuid.UUID) (record *programQuery.ProgramMainQueryEntity, err error)
	CreateOrUpdateCourseStructure(tx *gorm.DB, structures []dto.ProgramStructureRequestDto, planDetailID *uint, parentID *uint) (err error)
	CreateOrUpdatePLODetail(tx *gorm.DB, ploDetails []dto.ProgramPLODetailDto, ploId *uint, parentID *uint) (err error)
	CreateOrUpdateMajor(tx *gorm.DB, majors []dto.ProgramMajorDto, generalDetailId *uint, parentID *uint) (err error)
	CreateOrUpdatePlanDetail(tx *gorm.DB, plans []dto.ProgramPlanDetailDto, majorId *uint, parentID *uint) (err error)
	GetCourseInfoPagnation(paginationOptions *models.PaginationOptions) (record []courseQuery.CourseMainQueryEntity, err error)
	GetCourseDetailPaginationForYear(paginationOptions *models.PaginationOptions, planID uint) (record []programQuery.ProgramCourseDetailQueryEntity, err error)
	// CreateOrUpdateYLODetail(tx *gorm.DB, yloDetails []dto.ProgramYLODetailDto, programSubPlanID *uint, parentID *uint) (err error)
	GetProgramMainPagination(systemPermission *query.UserQueryEntity, role *rapQuery.RoleQueryEntity, paginationOptions *models.PaginationOptions, param dto.GetMainProgramPaginationQueryParam) (record []programQuery.ProgramMainQueryEntity, err error)
	CreateOrUpdateSubPlan(tx *gorm.DB, subPlans []dto.ProgramSubPlanDto, planDetailID *uint) (err error)
	GetOrCreateEmployeeByEmail(employee dto.LecturerDto) (EmployeeUID *uuid.UUID, err error)
	CreateOrUpdateKsecDetail(tx *gorm.DB, ksec dto.KsecRequestDto, subPlanID *uint) (err error)
	CreateOrUpdateYLOWithPLO(tx *gorm.DB, ploDetails []dto.ProgramYLOPLODto, yearAndSemesterID uint) error
	CreateOrUpdateYLOWithKsec(tx *gorm.DB, ksecDetail []dto.ProgramYLOPLODto, yearAndSemesterID uint) error
	CreateOrUpdateCurMapResp(tx *gorm.DB, paginationOptions *models.PaginationOptions, reqData dto.CreateOrUpdateCurMapRespRequestDto) (err error)
	CreateOrUpdateCurMapKsa(tx *gorm.DB, paginationOptions *models.PaginationOptions, reqData dto.CreateOrUpdateCurMapKsaRequestDto) (err error)
	CreatePLODetail(tx *gorm.DB, ploDetails []dto.ProgramPLODetailDto, ploId *uint, parentID *uint, ksecDetails []dto.MapKsecList) (err error)
	CreateKsecDetail(tx *gorm.DB, ksec dto.KsecRequestDto, subPlanID *uint) (err error)
	CreateOrUpdateMapFile(tx *gorm.DB, fileId *uint, id *uint) (err error)
	UpdateRefNilOption(tx *gorm.DB, id *uint) (err error)
}
