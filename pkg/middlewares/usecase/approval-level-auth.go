package usecase

// import (
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// 	helpers "github.com/zercle/gofiber-helpers"
// 	"github.com/zercle/kku-qf-services/pkg/utils"
// 	"github.com/zercle/kku-qf-services/pkg/utils/middlewares"
// 	"gorm.io/gorm"
// )

// func (uc *middlewaresUsecase) ApprovalLevelAuth(levelToCheck uint) fiber.Handler {
// 	return func(c *fiber.Ctx) (err error) {
// 		userUID, err := middlewares.GetUserUIDFromClaims(c)
// 		if err != nil {
// 			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), http.StatusText(http.StatusUnauthorized))
// 		}

// 		userInfo, err := uc.userUC.GetUserByID(userUID)
// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				return helpers.NewError(http.StatusInternalServerError, helpers.WhereAmI(), http.StatusText(http.StatusInternalServerError))
// 			}
// 			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), http.StatusText(http.StatusUnauthorized))
// 		}

// 		systemPermission, err := uc.permissionUC.GetOneSystemPermission(*userInfo.SystemPermissionUID)
// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				return helpers.NewError(http.StatusInternalServerError, helpers.WhereAmI(), http.StatusText(http.StatusInternalServerError))
// 			}
// 			return helpers.NewError(http.StatusUnauthorized, helpers.WhereAmI(), http.StatusText(http.StatusUnauthorized))
// 		}

// 		if !systemPermission.CanApprove {
// 			return helpers.NewError(http.StatusForbidden, helpers.WhereAmI(), http.StatusText(http.StatusForbidden))
// 		}

// 		if utils.IsInRangeUINT(levelToCheck, userInfo.ProgramApprovalMinLevel, userInfo.ProgramApprovalMaxLevel) {
// 			return c.Next()
// 		} else {
// 			return helpers.NewError(http.StatusForbidden, helpers.WhereAmI(), http.StatusText(http.StatusForbidden))
// 		}
// 	}
// }
