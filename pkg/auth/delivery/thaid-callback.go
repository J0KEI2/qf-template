package delivery

import (
	"net/http"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
)

func (h authHandler) ThaidLogin() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}

		request := entity.ThaidLoginRequestDto{}
		if err := c.QueryParser(&request); err != nil {
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		thaidData, thaidError := h.authUseCase.ExtractThaidData(request.Code)

		if thaidError != nil {
			responseForm.Errors = append(responseForm.Errors, *thaidError)
			return c.Status(thaidError.Code).JSON(responseForm)
		}

		if thaidData.Pid == "" {
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    http.StatusBadRequest,
				Title:   http.StatusText(http.StatusBadRequest),
				Message: "Pid from thaid not found.",
			})
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		hrData, err := h.hrUsecase.GetLecturerBySSN(thaidData.Pid)
		email := thaidData.Pid
		userType := enums.EXTERNAL
		facultyID := uint(0)
		if err == nil {
			email = hrData.Data.Email
			userType = enums.INTERNAL
			faculty, err := h.commonUsecase.GetFacultyByFacultyName(hrData.Data.Faculty)
			if err != nil {
				responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
					Code:    http.StatusBadRequest,
					Title:   http.StatusText(http.StatusBadRequest),
					Message: err.Error(),
				})
				return c.Status(http.StatusBadRequest).JSON(responseForm)
			}
			facultyID = *faculty.ID
		}

		user, err := h.userUsecase.GetOrCreateBySSN(models.UserCreateQuery{
			Email:       email,
			SSN:         &thaidData.Pid,
			TitleTh:     thaidData.Title,
			FirstnameTh: thaidData.GivenName,
			LastnameTh:  thaidData.FamilyName,
			TitleEn:     thaidData.TitleEn,
			FirstnameEn: thaidData.GivenNameEn,
			LastnameEn:  thaidData.FamilyNameEn,
			FacultyID:   facultyID,
			Type:        userType.ToString(),
		})
		if err != nil {
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    fiber.StatusInternalServerError,
				Title:   "Get user data in database",
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		role, err := h.roleUsecase.GetRoleByID(user.CurrentRoleID)
		if err != nil {
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    fiber.StatusInternalServerError,
				Title:   "Get user data in database",
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		jti, _ := helpers.UUIDv4()
		claims := dto.UserClaims{
			UserID: user.UID.String(),
			Role: dto.RoleClaims{
				ID:                   pointer.GetUint(role.ID),
				RoleNameTH:           role.RoleNameTH,
				RoleNameEN:           role.RoleNameEN,
				ProgramRoleType:      role.ProgramRoleType,
				CourseRoleType:       role.CourseRoleType,
				ProgramApprovalLevel: role.ProgramApprovalLevel,
				CourseApprovalLevel:  role.CourseApprovalLevel,
				ProgramActionLevel:   role.ProgramActionLevel,
				CourseActionLevel:    role.CourseActionLevel,
			},
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        jti,
				Issuer:    c.BaseURL(),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			},
		}

		oauthResponse, err := h.GenUserAccessToken(claims)

		if err != nil {
			responseForm.Errors = append(responseForm.Errors, helpers.ResponseError{
				Code:    fiber.StatusInternalServerError,
				Title:   "Error on generate jwt token",
				Message: err.Error(),
			})
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}

		responseForm.Success = true
		responseForm.Result = oauthResponse
		return c.JSON(responseForm)
	}
}
