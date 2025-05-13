package usecase

import (
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
	"gorm.io/gorm"
)

func (u roleUsecase) GetIfUserHasRoles(userID uuid.UUID) (ok bool, err error) {
	mapCoreRoleQuery := query.MapCoreRolesQueryEntity{
		UserID: &userID,
	}

	mapFacultyRoleQuery := query.MapFacultiesRolesQueryEntity{
		UserID: &userID,
	}

	mapProgramRoleQuery := query.MapProgramsRolesQueryEntity{
		UserID: &userID,
	}

	if err = u.CommonRepository.GetFirst(&mapProgramRoleQuery, "Role"); err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = u.CommonRepository.GetFirst(&mapFacultyRoleQuery, "Role"); err != nil {
				if err == gorm.ErrRecordNotFound {
					if err = u.CommonRepository.GetFirst(&mapCoreRoleQuery, "Role"); err != nil {
						log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
						return false, err
					}
					return true, nil
				}
				log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
				return false, err
			}
			return true, nil
		}
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return false, err
	}

	return
}
