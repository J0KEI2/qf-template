package usecase

import (
	"encoding/json"
	"strings"

	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u *userDetailUseCase) GetUserDetailPagination(options models.PaginationOptions) (result *dto.GetUserDetailPaginationResponseDto, err error) {
	userDetailRecords, err := u.userDetailRepo.GetUserDetailPagination(&options)

	if err != nil {
		return nil, err
	}
	items := make([]dto.UserDetailPagination, 0)

	for _, userDetail := range userDetailRecords {
		name := strings.Join([]string{*userDetail.TitleTh, *userDetail.FirstnameTh, *userDetail.LastnameTh}, " ")
		nameEN := strings.Join([]string{*userDetail.TitleEn, *userDetail.FirstnameEn, *userDetail.LastnameEn}, " ")
		position := []dto.LecturerPositionDto{}
		educationBackground := []dto.EducationalBackgroundDto{}
		if userDetail.Position != nil {
			json.Unmarshal([]byte(*userDetail.Position), &position)
		}
		if userDetail.EducationBackGround != nil {
			json.Unmarshal([]byte(*userDetail.EducationBackGround), &educationBackground)
		}
		userDetailData := dto.UserDetailPagination{
			UID:                 userDetail.UID,
			UserUID:             userDetail.UserUID,
			MiddleNameTh:        userDetail.MiddlenameTh,
			MiddleNameEn:        userDetail.MiddlenameEn,
			EducationBackGround: educationBackground,
			Position:            position,
			TitleTh:             userDetail.TitleTh,
			TitleEn:             userDetail.TitleEn,
			FirstNameTh:         userDetail.FirstnameTh,
			FirstNameEn:         userDetail.FirstnameEn,
			LastNameTh:          userDetail.LastnameTh,
			LastNameEn:          userDetail.LastnameEn,
			Name:                &name,
			NameEN:              &nameEN,
		}

		items = append(items, userDetailData)
	}

	result = &dto.GetUserDetailPaginationResponseDto{
		Items:   items,
		Options: options,
	}

	return
}
