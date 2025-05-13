package middlewares

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func GetUserUIDFromClaims(c *fiber.Ctx) (userUID uuid.UUID, err error) {
	claims, ok := c.Locals("userClaims").(*dto.UserClaims)
	if !ok {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), "not ok")
		return userUID, errors.New("error on parsing jwt")
	}

	userUID, err = uuid.Parse(claims.UserID)
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return userUID, errors.New("error on parsing userUID")
	}
	return
}

func GetRoleIDFromClaims(c *fiber.Ctx) (roleID *uint, err error) {
	claims, ok := c.Locals("userClaims").(*dto.UserClaims)
	if !ok {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), "not ok")
		return nil, errors.New("error on parsing jwt")
	}

	return &claims.Role.ID, nil
}

func GetRoleInfoFromClaims(c *fiber.Ctx) (roleInfo *dto.RoleClaims, err error) {
	claims, ok := c.Locals("userClaims").(*dto.UserClaims)
	if !ok {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), "not ok")
		return nil, errors.New("error on parsing jwt")
	}

	return &claims.Role, nil
}

func GetProgramApprovalIDFromClaims(c *fiber.Ctx) (ProgramApprovalID *uint, err error) {
	claims, ok := c.Locals("userClaims").(*dto.UserClaims)
	if !ok {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), "not ok")
		return nil, errors.New("error on parsing jwt")
	}

	return claims.Role.ProgramApprovalLevel, nil
}

// func GetCourseApprovalIDFromClaims(c *fiber.Ctx) (CourseApprovalID *uint, err error) {
// 	claims, ok := c.Locals("userClaims").(*dto.UserClaims)
// 	if !ok {
// 		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), "not ok")
// 		return nil, errors.New("error on parsing jwt")
// 	}

// 	return claims.Role.CourseApprovalLevel, nil
// }
