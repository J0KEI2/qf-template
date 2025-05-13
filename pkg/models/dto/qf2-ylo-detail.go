package dto

import (
	"time"

	"gorm.io/gorm"
)

type ProgramYLODetailGetResponseDto struct {
	ProgramSubPlanID uint                  `json:"program_sub_plan_id"`
	YLODetails       []ProgramYLODetailDto `json:"ylo_details"`
}

type CreateOrUpdateYLODetailRequestDto struct {
	ProgramSubPlanID uint                  `json:"program_sub_plan_id"`
	YLODetails       []ProgramYLODetailDto `json:"ylo_details"`
}

type ProgramYLODetailDto struct {
	ID                       *uint             `json:"id"`
	ProgramYearAndSemesterID *uint             `json:"program_year_and_semester_id"`
	Year                     *string           `json:"year"`
	Knowledge                *string           `json:"knowledge"`
	Skill                    *string           `json:"skill"`
	Ethic                    *string           `json:"ethic"`
	Character                *string           `json:"character"`
	YLOData                  ProgramYLODataDto `json:"ylo_data"`
	CreatedAt                *time.Time        `json:"created_at"`
	UpdatedAt                *time.Time        `json:"updated_at"`
	DeletedAt                *gorm.DeletedAt   `json:"deleted_at,omitempty"`
}

type ProgramYLODataDto struct {
	ID         *uint              `json:"id"`
	PLOFormat  *string            `json:"plo_format"`
	PLODetails []ProgramYLOPLODto `json:"plo_details"`
	CreatedAt  *time.Time         `json:"created_at"`
	UpdatedAt  *time.Time         `json:"updated_at"`
	DeletedAt  *gorm.DeletedAt    `json:"deleted_at,omitempty"`
}

type ProgramYLOPLODto struct {
	ID        *uint              `json:"id"`
	MapYLOID  *uint              `json:"map_ylo_id"`
	Order     *uint              `json:"order"`
	ParentID  *uint              `json:"parent_id"`
	Children  []ProgramYLOPLODto `json:"children"`
	Ksec      []YLOKsecDetail    `json:"ksec"`
	PLOPrefix *string            `json:"plo_prefix"`
	PLODetail *string            `json:"plo_detail"`
	Remark    *string            `json:"remark"`
	IsChecked *bool              `json:"is_checked"`
	CreatedAt *time.Time         `json:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at"`
	DeletedAt *gorm.DeletedAt    `json:"deleted_at,omitempty"`
}

type YLOKsecDetail struct {
	ID        *uint      `json:"id"`
	MapYLOID  *uint      `json:"map_ylo_id"`
	MapPLOID  *uint      `json:"map_plo_id"`
	Type      *string    `json:"type"`
	Order     *uint      `json:"order"`
	Detail    *string    `json:"detail"`
	IsChecked *bool      `json:"is_checked"`
	Remark    *string    `json:"remark"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
