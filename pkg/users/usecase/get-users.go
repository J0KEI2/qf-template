package usecase

import (
	"strings"

	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u userUsecase) GetUsers(userQueryEntity query.UserQueryEntity, options *models.PaginationOptions) (userFetchEntity *entity.UserListFetchEntity, err error) {
	results, err := u.userRepo.GetUserList(nil, options)

	if err != nil {
		return nil, err
	}

	userList := make([]entity.UserFetchEntity, len(results))
	
	for i, res := range results {
		nameTHList := append(strings.SplitN(res.NameTH, " ", 3), []string{"", ""}...)
		nameENList := append(strings.SplitN(res.NameEN, " ", 3), []string{"", ""}...)
		userList[i] = entity.UserFetchEntity{
			TitleTh:             nameTHList[0],
			FirstnameTh:         nameTHList[1],
			LastnameTH:          nameTHList[2],
			TitleEn:             nameENList[0],
			FirstnameEn:         nameENList[1],
			LastnameEn:          nameENList[2],
			Email:               res.Email,
			Status:              res.Status,
			SSN:                 res.SSN,
			FacultyID:           res.FacultyID,
			SystemPermissionUID: res.SystemPermissionUID,
			UID:                 res.UID,
		}
	}

	userFetchEntity = &entity.UserListFetchEntity{
		Pagination: options,
		Data:       userList,
	}

	return userFetchEntity, nil
}
