package dto

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
)

type GetUserDetailPaginationResponseDto struct {
	Items   []UserDetailPagination   `json:"items"`
	Options models.PaginationOptions `json:"options"`
}

type UserDetailPagination struct {
	UID                 *uuid.UUID                 `json:"uid"`
	UserUID             *uuid.UUID                 `json:"user_uid"`
	MiddleNameTh        *string                    `json:"middle_name_th"`
	MiddleNameEn        *string                    `json:"middle_name_en"`
	EducationBackGround []EducationalBackgroundDto `json:"education_backgrounds"`
	Position            []LecturerPositionDto      `json:"positions"`
	TitleTh             *string                    `json:"title_th"`
	TitleEn             *string                    `json:"title_en"`
	FirstNameTh         *string                    `json:"first_name_th"`
	FirstNameEn         *string                    `json:"first_name_en"`
	LastNameTh          *string                    `json:"last_name_th"`
	LastNameEn          *string                    `json:"last_name_en"`
	Name                *string                    `json:"name"`
	NameEN              *string                    `json:"name_en"`
}
