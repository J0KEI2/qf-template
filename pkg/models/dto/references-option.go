package dto

import "github.com/zercle/kku-qf-services/pkg/models"

type GetReferencesOptionResponseDto struct {
	Items   []ReferenceOption        `json:"items"`
	Options models.PaginationOptions `json:"options"`
}

type ReferenceOption struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type ReferencesOptionResponseDto struct {
	Items []ReferenceOption `json:"items"`
}
