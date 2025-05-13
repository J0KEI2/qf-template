package query

import (
	"time"

	"gorm.io/gorm"
)

type ProgramQualityAssuranceQueryEntity struct {
	ID               *uint      `gorm:"column:id;primaryKey;autoIncrement"`
	IsHescCheck      *bool      `gorm:"column:is_hesc_check;type:boolean"`
	HescDescription  *string    `gorm:"column:hesc_description;type:text"`
	IsAunQaCheck     *bool      `gorm:"column:is_aun_qa_check;type:boolean"`
	AunQaDescription *string    `gorm:"column:aun_qa_description;type:text"`
	IsAbetCheck      *bool      `gorm:"column:is_abet_check;type:boolean"`
	AbetDescription  *string    `gorm:"column:abet_description;type:text"`
	IsWfmeCheck      *bool      `gorm:"column:is_wfme_check;type:boolean"`
	WfmeDescription  *string    `gorm:"column:wfme_description;type:text"`
	IsAacsbCheck     *bool      `gorm:"column:is_aacsb_check;type:boolean"`
	AacsbDescription *string    `gorm:"column:aacsb_description;type:text"`
	CreatedAt        *time.Time `gorm:"column:created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt
}

func (q *ProgramQualityAssuranceQueryEntity) TableName() string {
	return "program_quality_assurance"
}
