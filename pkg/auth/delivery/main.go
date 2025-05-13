package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github.com/zercle/kku-qf-services/pkg/domain"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type authHandler struct {
	authUseCase       domain.AuthUseCase
	userUsecase       domain.UserUsecase
	hrUsecase         domain.HRUseCase
	commonUsecase     domain.CommonUseCase
	permissionUsecase domain.PermissionUseCase
	roleUsecase       domain.RoleUseCase
	fasthttpClient    *fasthttp.Client
	jwtResources      *models.JwtResources
}

func NewAuthHandler(authRoute fiber.Router, authUseCase domain.AuthUseCase, userUsecase domain.UserUsecase, permissionUsecase domain.PermissionUseCase, hrUsecase domain.HRUseCase, commonUsecase domain.CommonUseCase, mdwUC domain.MiddlewaresUseCase, roleUsecase domain.RoleUseCase, fasthttpClient *fasthttp.Client, jwtResources *models.JwtResources) {
	handler := &authHandler{
		authUseCase:       authUseCase,
		userUsecase:       userUsecase,
		hrUsecase:         hrUsecase,
		commonUsecase:     commonUsecase,
		permissionUsecase: permissionUsecase,
		roleUsecase:       roleUsecase,
		fasthttpClient:    fasthttpClient,
		jwtResources:      jwtResources,
	}

	authRoute.Get("/login", handler.Login())
	authRoute.Post("/thaid-login", handler.ThaidLogin())
	authRoute.All("/logout", handler.Logout())
	authRoute.All("/callback", handler.OauthCallback())
	authRoute.Get("/", handler.CurrentToken())
	authRoute.Patch("/switch-role", mdwUC.ExtractUserJwt(), handler.SwitchUserRole())

	authRoute.Get("/test-uam/1", mdwUC.ExtractUserJwt(), mdwUC.SystemAuth(1), handler.TestAuth())    // get user
	authRoute.Post("/test-uam/2", mdwUC.ExtractUserJwt(), mdwUC.SystemAuth(2), handler.TestAuth())   // create user
	authRoute.Patch("/test-uam/2", mdwUC.ExtractUserJwt(), mdwUC.SystemAuth(2), handler.TestAuth())  // update user
	authRoute.Delete("/test-uam/3", mdwUC.ExtractUserJwt(), mdwUC.SystemAuth(3), handler.TestAuth()) // delete user

	authRoute.Get("/test-read-approval", mdwUC.ExtractUserJwt(), mdwUC.ProgramApprovalAuth(1), handler.TestAuth())      // can see approval, comment
	authRoute.Get("/test-comment", mdwUC.ExtractUserJwt(), mdwUC.ProgramApprovalAuth(2), handler.TestAuth())            // can see, edit comment
	authRoute.Patch("/test-approval", mdwUC.ExtractUserJwt(), mdwUC.ProgramApprovalAuth(3), handler.TestAuth())         // can update approval
	authRoute.Delete("/test-delete-approval", mdwUC.ExtractUserJwt(), mdwUC.ProgramApprovalAuth(4), handler.TestAuth()) // can delete approval? NOTE: can revert or smth TBC

	authRoute.Get("/test-read-program", mdwUC.ExtractUserJwt(), mdwUC.ProgramActionAuth(1), handler.TestAuth())           // can see program detail
	authRoute.Post("/test-create-update-program", mdwUC.ExtractUserJwt(), mdwUC.ProgramActionAuth(2), handler.TestAuth()) // can create, update program
	authRoute.Delete("/test-delete-program", mdwUC.ExtractUserJwt(), mdwUC.ProgramActionAuth(3), handler.TestAuth())      // can delete program

}
