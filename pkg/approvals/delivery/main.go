package approvals

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/domain"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

type programApprovalHandler struct {
	ApprovalUseCase domain.ProgramApprovalUsecase
}

func NewProgramApprovalHandler(approvalRoute fiber.Router, approvalUseCase domain.ProgramApprovalUsecase, mdwUC domain.MiddlewaresUseCase) {
	handler := &programApprovalHandler{
		ApprovalUseCase: approvalUseCase,
	}

	pathPrefix := "/:program_uid"
	approvePrefix := pathPrefix + "/approve"
	rejectPrefix := pathPrefix + "/reject"

	programApprovalGroup := approvalRoute.Group("/programs")
	programApprovalGroup.Get(pathPrefix, handler.GetAllApproval())

	facultyGroup := programApprovalGroup.Group("/faculty")
	{
		facultyGroup.Get(pathPrefix, handler.GetFacultyApprovalByProgramUID())
		facultyGroup.Patch(approvePrefix, handler.ApproveFacultyApprovalByProgramUID())
		facultyGroup.Patch(rejectPrefix, handler.RejectFacultyApprovalByProgramUID())
	}
	committeeGroup := programApprovalGroup.Group("/committee")
	{
		committeeGroup.Get(pathPrefix, handler.GetCurriculumCommitteeApprovalByProgramUID())
		committeeGroup.Post(pathPrefix+"/select-committees", handler.SelectCurriculumCommittees())
		committeeGroup.Patch(pathPrefix+"/committee-approve", handler.ApproveCurriculumCommitteeResult())
		committeeGroup.Patch(pathPrefix+"/committee-reject", handler.RejectCurriculumCommitteeResult())
		committeeGroup.Patch(approvePrefix, handler.ApproveCurriculumCommitteeApprovalByProgramUID())
	}
	academicGroup := programApprovalGroup.Group("/academic")
	{
		academicGroup.Get(pathPrefix, handler.GetAcademicApprovalByProgramUID())
		academicGroup.Patch(approvePrefix, handler.ApproveAcademicApprovalByProgramUID())
		academicGroup.Patch(rejectPrefix, handler.RejectAcademicApprovalByProgramUID())
	}
	universityGroup := programApprovalGroup.Group("/university")
	{
		universityGroup.Get(pathPrefix, handler.GetUniversityApprovalByProgramUID())
		universityGroup.Patch(approvePrefix, handler.ApproveUniversityApprovalByProgramUID())
		universityGroup.Patch(rejectPrefix, handler.RejectUniversityApprovalByProgramUID())
	}

	// ./approval/checo... here
	checoGroup := programApprovalGroup.Group("/checo")
	{
		checoGroup.Get(pathPrefix, handler.GetCHECOByProgramUID())
		checoGroup.Post("/", handler.CreateCHECOStatus())
	}

}

func prepareApproveApprovalRequest(c *fiber.Ctx) (request dto.ApproveApprovalRequestDto, updatedByUUID uuid.UUID, err error) {
	var approvedDateParsedTime time.Time
	var submissionID int
	meetingNo := c.FormValue("meeting_no", "")
	updatedBy := c.FormValue("updated_by", "")
	committeeUserID := c.FormValue("committee_user_id", "")

	approvedDate := c.FormValue("approved_date", "")
	if approvedDate != "" {
		cleanedDateStr := strings.Split(approvedDate, " (")[0]
		approvedDateParsedTime, err = time.Parse(constant.FORM_DATA_TIME_FORMAT, cleanedDateStr)
		if err != nil {
			return dto.ApproveApprovalRequestDto{}, updatedByUUID, err
		}
	} else {
		approvedDateParsedTime = time.Now()
	}

	submissionIDstr := c.FormValue("submission_id", "")
	if submissionIDstr != "" {
		submissionID, err = strconv.Atoi(submissionIDstr)
		if err != nil {
			return dto.ApproveApprovalRequestDto{}, updatedByUUID, err
		}
	}

	// extract updatedBy from token if empty
	if updatedBy == "" {
		updatedByUUID, err = middlewares.GetUserUIDFromClaims(c)
		if err != nil {
			return dto.ApproveApprovalRequestDto{}, updatedByUUID, err
		}
		updatedBy = updatedByUUID.String()
	} else {
		updatedByUUID, err = uuid.Parse(updatedBy)
		if err != nil {
			return dto.ApproveApprovalRequestDto{}, updatedByUUID, err
		}
	}

	request = dto.ApproveApprovalRequestDto{
		Comment:         c.FormValue("comment", ""),
		MeetingNo:       &meetingNo,
		CommitteeUserID: committeeUserID,
		SubmissionID:    uint(submissionID),
		ApprovedDate:    &approvedDateParsedTime,
		UpdatedBy:       updatedBy,
	}

	return
}

func prepareRejectApprovalRequest(c *fiber.Ctx) (request dto.RejectApprovalRequestDto, updatedByUUID uuid.UUID, err error) {
	var submissionID int
	updatedBy := c.FormValue("updated_by", "")
	committeeUserID := c.FormValue("committee_user_id", "")

	submissionIDstr := c.FormValue("submission_id", "")
	if submissionIDstr != "" {
		submissionID, err = strconv.Atoi(submissionIDstr)
		if err != nil {
			return dto.RejectApprovalRequestDto{}, updatedByUUID, err
		}
	}

	// extract updatedBy from token if empty
	if updatedBy == "" {
		updatedByUUID, err = middlewares.GetUserUIDFromClaims(c)
		if err != nil {
			return dto.RejectApprovalRequestDto{}, updatedByUUID, err
		}
		updatedBy = updatedByUUID.String()
	} else {
		updatedByUUID, err = uuid.Parse(updatedBy)
		if err != nil {
			return dto.RejectApprovalRequestDto{}, updatedByUUID, err
		}
	}

	request = dto.RejectApprovalRequestDto{
		Comment:         c.FormValue("comment", ""),
		CommitteeUserID: committeeUserID,
		SubmissionID:    uint(submissionID),
		UpdatedBy:       updatedBy,
	}

	return
}
