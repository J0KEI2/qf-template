package delivery

import (
	"log"
	"net/http"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
)

func (h *authHandler) SwitchUserRole() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.OauthResponse{}

		reqBody := dto.SwitchRoleRequest{}
		if err = c.QueryParser(&reqBody); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			responseForm.Error = helpers.ServerError
			responseForm.ErrorDesc = err.Error()
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		userID, err := middlewares.GetUserUIDFromClaims(c)
		if err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			responseForm.Error = helpers.ServerError
			responseForm.ErrorDesc = err.Error()
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		possibleRoles, err := h.roleUsecase.GetPossibleRole(userID)
		if err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
			responseForm.Error = helpers.ServerError
			responseForm.ErrorDesc = "cannot switch to requested role"
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}

		var ok bool
		if reqBody.RoleID != 0 {
			for _, eachRole := range possibleRoles.Items {
				if *eachRole.ID == reqBody.RoleID {
					ok = true
					break
				}
			}
		} else {
			ok = true
		}

		if !ok {
			log.Print("SwitchUserRole: cannot switch to requested role")
			responseForm.Error = helpers.AccessDenied
			responseForm.ErrorDesc = "cannot switch to requested role"
			return c.Status(http.StatusForbidden).JSON(responseForm)
		}

		role, err := h.roleUsecase.GetRoleByID(reqBody.RoleID)
		if err != nil {
			log.Printf("SwitchUserRole: %v", err)
			responseForm.Error = helpers.ServerError
			responseForm.ErrorDesc = err.Error()
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		err = h.roleUsecase.UpdateUserCurrentRole(userID, reqBody.RoleID)
		if err != nil {
			log.Printf("SwitchUserRole: %v", err)
			responseForm.Error = helpers.ServerError
			responseForm.ErrorDesc = err.Error()
			return c.Status(http.StatusInternalServerError).JSON(responseForm)
		}

		// fmt.Printf("\n roleID: %+v \n userID: %+v \n\n", reqBody.RoleID, userID.String())

		if len(userID.String()) > 0 {
			jti, _ := helpers.UUIDv4()
			claims := dto.UserClaims{
				UserID: userID.String(),
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
				log.Printf("SwitchUserRole: %v", err)
				responseForm.Error = helpers.ServerError
				responseForm.ErrorDesc = err.Error()
				return c.Status(http.StatusInternalServerError).JSON(responseForm)
			}

			if fiberErr, ok := err.(*fiber.Error); ok {
				c.Status(fiberErr.Code)
			} else {
				c.Status(http.StatusUnauthorized)
			}
			return c.Status(http.StatusOK).JSON(responseForm)
		}

		responseForm.Error = helpers.InvalidRequest
		responseForm.ErrorDesc = "cannot get role id from token"
		return c.Status(http.StatusBadRequest).JSON(responseForm)
	}
}
