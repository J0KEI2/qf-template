package delivery

import (
	"log"
	"net/http"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
)

func (h *authHandler) OauthCallback() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.OauthResponse{}

		// for debug
		// code := c.FormValue("code")
		// fmt.Println(">>> code: " + code)

		reqBody := models.OauthCallback{}
		if err = c.BodyParser(&reqBody); err != nil {
			responseForm.Error = helpers.Unavailable
			responseForm.ErrorDesc = err.Error()
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		if len(reqBody.Code) != 0 {
			accessToken, _, err := h.GetOauthToken(reqBody.Code)
			if err != nil {
				responseForm.Error = helpers.Unavailable
				responseForm.ErrorDesc = err.Error()
				return c.Status(http.StatusServiceUnavailable).JSON(responseForm)
			}

			userSSO, err := h.GetUserSSO(string(accessToken))
			if err != nil {
				log.Printf("OauthCallback: %v", err)
				responseForm.Error = helpers.AccessDenied
				responseForm.ErrorDesc = err.Error()
				return c.Status(http.StatusServiceUnavailable).JSON(responseForm)
			}

			log.Printf("%v", userSSO)

			// get user_id by ssn; if not exist, create new user
			faculty, err := h.commonUsecase.GetFacultyByFacultyName(userSSO.FacultyName)
			if err != nil {
				responseForm.Error = helpers.ServerError
				responseForm.ErrorDesc = err.Error()
				return c.Status(http.StatusInternalServerError).JSON(responseForm)
			}

			userInfo, err := h.userUsecase.GetOrCreateBySSN(models.UserCreateQuery{
				Email:       userSSO.Mail,
				SSN:         &userSSO.CitizenID,
				TitleTh:     userSSO.Title,
				FirstnameTh: userSSO.Firstname,
				LastnameTh:  userSSO.Lastname,
				TitleEn:     userSSO.TitleEng,
				FirstnameEn: userSSO.FirstnameEng,
				LastnameEn:  userSSO.LastnameEng,
				FacultyID:   *faculty.ID,
				Type:        enums.INTERNAL.ToString(),
			})
			if err != nil {
				log.Printf("OauthCallback: %v", err)
				responseForm.Error = helpers.ServerError
				responseForm.ErrorDesc = err.Error()
				return c.Status(http.StatusServiceUnavailable).JSON(responseForm)
			}

			role, err := h.roleUsecase.GetRoleByID(userInfo.CurrentRoleID)
			if err != nil {
				log.Printf("OauthCallback: %v", err)
				responseForm.Error = helpers.ServerError
				responseForm.ErrorDesc = err.Error()
				return c.Status(http.StatusServiceUnavailable).JSON(responseForm)
			}

			// TODO: add claims value here
			jti, _ := helpers.UUIDv4()
			claims := dto.UserClaims{
				UserID: userInfo.UID.String(),
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

			responseForm, err = h.GenUserAccessToken(claims)
			if err != nil {
				log.Printf("OauthCallback: %v", err)
				responseForm.Error = helpers.ServerError
				responseForm.ErrorDesc = err.Error()
				return c.Status(http.StatusServiceUnavailable).JSON(responseForm)
			}

			if fiberErr, ok := err.(*fiber.Error); ok {
				c.Status(fiberErr.Code)
			} else {
				c.Status(http.StatusUnauthorized)
			}
			return c.Status(http.StatusOK).JSON(responseForm)
		}

		responseForm.Error = helpers.InvalidRequest
		return c.Status(http.StatusBadRequest).JSON(responseForm)
	}
}
