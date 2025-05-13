package dto

import "github.com/zercle/kku-qf-services/pkg/models"

type CreateOrUpdateYearCourseRequestDto struct {
	Years []ProgramYearRequestDto `json:"years"`
}

type GetYearCourseResponseDto struct {
	Items []ProgramYearResponseDto `json:"items"`
	*models.PaginationOptions
}

type GetCourseDetailPaginationResponseDto struct {
	Items []ProgramCourseDetailResponseDto `json:"items"`
	*models.PaginationOptions
}
