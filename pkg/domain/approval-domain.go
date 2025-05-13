package domain

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

type ProgramApprovalUsecase interface {
	GetAllApproval(programUID *uuid.UUID) (*[]dto.GetAllApprovalDto, error)

	GetFacultyApprovalByProgramUID(programUID *uuid.UUID, userProgramLevel uint) (result *dto.GetApprovalFacultyDto, err error)
	ApproveFacultyApprovalByProgramUID(programUID *uuid.UUID, request dto.ApproveApprovalRequestDto) (result *dto.GetApprovalFacultyDto, err error)
	RejectFacultyApprovalByProgramUID(programUID *uuid.UUID, request dto.RejectApprovalRequestDto) (result *dto.GetApprovalFacultyDto, err error)

	GetCurriculumCommitteeApprovalByProgramUID(programUID, userUID *uuid.UUID, userProgramLevel uint) (result *dto.GetApprovalCurriculumCommitteeDto, err error)
	GetCurriculumCommitteesApprovalByProgramUID(programUID *uuid.UUID, userProgramLevel uint) (result *dto.GetApprovalCurriculumCommitteeDto, err error)
	SelectCurriculumCommittees(programUID *uuid.UUID, committeeUserIDs []string) (result *dto.GetApprovalCurriculumCommitteeDto, err error)
	ApproveCurriculumCommitteeResult(programUID *uuid.UUID, request dto.ApproveApprovalCommitteeResultRequestDto) (result *dto.GetApprovalCurriculumCommitteeDto, err error)
	RejectCurriculumCommitteeResult(programUID *uuid.UUID, request dto.RejectApprovalCommitteeResultRequestDto) (result *dto.GetApprovalCurriculumCommitteeDto, err error)
	// CommitteeReject(programUID *uuid.UUID) (result *dto.GetApprovalCurriculumCommitteeDto, err error)
	ApproveCurriculumCommitteeApprovalByProgramUID(programUID *uuid.UUID, request dto.ApproveApprovalRequestDto, userProgramLevel uint) (result *dto.GetApprovalCurriculumCommitteeDto, err error)

	GetAcademicApprovalByProgramUID(programUID *uuid.UUID, userProgramLevel uint) (result *dto.GetApprovalAcademicDto, err error)
	ApproveAcademicApprovalByProgramUID(programUID *uuid.UUID, request dto.ApproveApprovalRequestDto) (result *dto.GetApprovalAcademicDto, err error)
	RejectAcademicApprovalByProgramUID(programUID *uuid.UUID, request dto.RejectApprovalRequestDto) (result *dto.GetApprovalAcademicDto, err error)

	GetUniversityApprovalByProgramUID(programUID *uuid.UUID, userProgramLevel uint) (result *dto.GetApprovalUniversityDto, err error)
	ApproveUniversityApprovalByProgramUID(programUID *uuid.UUID, request dto.ApproveApprovalRequestDto) (result *dto.GetApprovalUniversityDto, err error)
	RejectUniversityApprovalByProgramUID(programUID *uuid.UUID, request dto.RejectApprovalRequestDto) (result *dto.GetApprovalUniversityDto, err error)

	UploadApprovalDocumentFile(c *fiber.Ctx, programUID uuid.UUID, files []*multipart.FileHeader, approvalAttribute string, userUID uuid.UUID, approvalID *uint) error

	// + Create 4 level of approval when new program created
	// CreateProgramApproval(programUID *uuid.UUID) (err error)

	CreateCHECOStatus(request dto.CreateCHECOStatusRequestDto, uid uuid.UUID) (checoId *uint, err error)
	GetCHECOByProgramUID(programUID *uuid.UUID) (result *dto.GetCHECOByUIDResponseDto, err error)
	UploadChecoDocumentFile(c *fiber.Ctx, programUID uuid.UUID, files []*multipart.FileHeader, userUID uuid.UUID, checoID *uint) error
}

type ProgramApprovalRepository interface {
	DbApprovalVCMigrator() (err error)
	GetFacultyApprovalByProgramUID(programUID uuid.UUID) (facultyApproval *dto.GetApprovalFacultyDto, err error)
	GetCurriculumCommitteeApprovalByProgramUID(programUID, userUID uuid.UUID) (*migrateModels.ProgramApproval, []dto.GetApprovalCommitteesDto, error)
	GetCurriculumCommitteesApprovalByProgramUID(programUID uuid.UUID) (*migrateModels.ProgramApproval, []dto.GetApprovalCommitteesDto, error)
	GetAcademicApprovalByProgramUID(programUID uuid.UUID) (academicApproval *dto.GetApprovalAcademicDto, err error)
	GetUniversityApprovalByProgramUID(programUID uuid.UUID) (universityApproval *dto.GetApprovalUniversityDto, err error)

	SelectCurriculumCommittees(programUID *uuid.UUID, committeeUserIDs []string) (err error)
	CommitteesSelected(approvalID uint) (err error)

	ApproveApproval(programUID uuid.UUID, approvalStatusLevel uint, request dto.ApproveApprovalRequestDto) (err error)
	ApproveCurriculumCommitteeApproval(programUID uuid.UUID, approvalStatusLevel uint, request dto.ApproveApprovalRequestDto) (err error)
	ApproveCurriculumCommitteeResult(programUID *uuid.UUID, request dto.ApproveApprovalCommitteeResultRequestDto) (err error)
	RejectCurriculumCommitteeResult(programUID *uuid.UUID, request dto.RejectApprovalCommitteeResultRequestDto) (err error)
	RejectApproval(programUID uuid.UUID, approvalStatusLevel uint, request dto.RejectApprovalRequestDto) (err error)

	// + Create 4 level of approval when new program created
	// CreateProgramApproval(programUID *uuid.UUID)

	GetCurrentApprovalProgress(programUID uuid.UUID) (currProgress *uint, err error)
	GetCurrentCHECOProgress(programUID uuid.UUID) (currProgress *uint, err error)
}
