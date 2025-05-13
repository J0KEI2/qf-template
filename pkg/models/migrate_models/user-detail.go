package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDetail struct {
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
	MiddlenameTh        *string   `gorm:"column:middlename_th;type:varchar;size:255" json:"middlenameTh"`
	MiddlenameEn        *string   `gorm:"column:middlename_en;type:varchar;size:255" json:"middlenameEn"`
	EducationBackGround *string   `gorm:"column:education_back_ground;type:varchar;size:511" json:"educationBackGround"`
	Position            *string   `gorm:"column:position;type:varchar" json:"position"`
	TitleTh             string    `gorm:"column:title_th;type:varchar;size:255;not null" json:"titleTh"`
	TitleEn             string    `gorm:"column:title_en;type:varchar;size:255;not null" json:"titleEn"`
	FirstnameTh         string    `gorm:"column:firstname_th;type:varchar;size:255;not null" json:"firstnameTh"`
	FirstnameEn         string    `gorm:"column:firstname_en;type:varchar;size:255;not null" json:"firstnameEn"`
	LastnameTh          string    `gorm:"column:lastname_th;type:varchar;size:255;not null" json:"lastnameTh"`
	LastnameEn          string    `gorm:"column:lastname_en;type:varchar;size:255;not null" json:"lastnameEn"`
	UID                 uuid.UUID `gorm:"column:uid;type:uuid;not null;primaryKey" json:"uid"`
	UserUID             uuid.UUID `gorm:"column:user_uid;type:uuid;unique" json:"userUid"`
	
}

func (l *UserDetail) BeforeCreate(tx *gorm.DB) (err error) {
	//create UUID
	l.UID = uuid.New()
	return
}
