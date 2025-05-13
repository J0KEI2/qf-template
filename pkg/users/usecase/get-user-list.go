package usecase

import (
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
)

func (usecase userUsecase) GetUserList(entity entity.UserFetchQueryEntity, options *models.PaginationOptions) (result *entity.UserFetchListEntity, err error) {

	offset := options.GetOffset()

	fetchUsers, err := usecase.userRepo.GetUserList(&models.UserFetchWithRelationQueryModel{
		UID:           entity.UID,
		Name:          entity.Name,
		Email:         entity.Email,
		FacultyNameTH: entity.FacultyNameTH,
		RoleNameEN:    entity.RoleSystemNameEN,
		Status:        entity.Status,
		OrderBy:       entity.OrderBy,
		Direction:     entity.Direction,
		Limit:         entity.Limit,
		Offset:        &offset,
	}, options)

	if err != nil {
		return nil, err
	}

	totalRecord, err := usecase.userRepo.CountUserList(&models.UserFetchWithRelationQueryModel{
		UID:           entity.UID,
		Name:          entity.Name,
		Email:         entity.Email,
		FacultyNameTH: entity.FacultyNameTH,
		RoleNameEN:    entity.RoleSystemNameEN,
		Status:        entity.Status,
		OrderBy:       entity.OrderBy,
		Direction:     entity.Direction,
		Limit:         entity.Limit,
		Offset:        &offset,
	})

	if err != nil {
		return nil, err
	}

	options.SetTotal(pointer.GetInt64(totalRecord))
	options.CalTotalPage()

	return prepareUserFetchListEntity(prepareUserFetchEntity(usecase, fetchUsers), options), nil
}

func prepareUserFetchListEntity(entities []entity.UserFetchEntity, options *models.PaginationOptions) *entity.UserFetchListEntity {
	return &entity.UserFetchListEntity{
		Items:   entities,
		Options: options,
	}
}

func prepareUserFetchEntity(usecase userUsecase, fetchUsers []models.UserFetchWithRelationModel) []entity.UserFetchEntity {
	results := []entity.UserFetchEntity{}

	for _, fetchUser := range fetchUsers {
		nameTHList := append(strings.SplitN(fetchUser.NameTH, " ", 3), []string{"", ""}...)
		nameENList := append(strings.SplitN(fetchUser.NameEN, " ", 3), []string{"", ""}...)
		results = append(results, entity.UserFetchEntity{
			UID:                 fetchUser.UID,
			Email:               fetchUser.Email,
			SSN:                 fetchUser.SSN,
			HRKey:               fetchUser.SSN,
			SystemPermissionUID: fetchUser.SystemPermissionUID,
			FacultyID:           fetchUser.FacultyID,
			FacultyNameEN:       fetchUser.FacultyNameEN,
			FacultyNameTH:       fetchUser.FacultyNameTH,
			Type:                fetchUser.Type,
			Status:              fetchUser.Status,
			CreatedAt:           fetchUser.CreatedAt,
			UpdatedAt:           fetchUser.UpdatedAt,
			LastAccessAt:        fetchUser.LastAccessAt,
			TitleTh:             nameTHList[0],
			FirstnameTh:         nameTHList[1],
			LastnameTH:          nameTHList[2],
			TitleEn:             nameENList[0],
			FirstnameEn:         nameENList[1],
			LastnameEn:          nameENList[2],
			NameTH:              fetchUser.NameTH,
			NameEN:              fetchUser.NameEN,
			RoleNameTH:          fetchUser.RoleNameTH,
			RoleNameEN:          fetchUser.RoleNameEN,
		})
	}
	return results
}
