package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/migrate_models/enums"
)

type UserDetailFetchQueryEntity struct {
	MiddlenameEn        *string
	EducationBackGround *string
	MiddlenameTh        *string
	Position            *enums.LecturerPosition
	TitleTh             *string
	FirstnameEn         *string
	TitleEn             *string
	LastnameEn          *string
	LastnameTh          *string
	FirstnameTh         *string
	UID                 *uuid.UUID
}
type UserDetailFetchEntity struct {
	UpdatedAt           time.Time
	CreatedAt           time.Time
	MiddlenameEn        *string
	EducationBackGround *string
	MiddlenameTh        *string
	Position            *enums.LecturerPosition
	TitleTh             string
	FirstnameEn         string
	TitleEn             string
	LastnameEn          string
	LastnameTh          string
	FirstnameTh         string
	UID                 uuid.UUID
}

type EducationBackGround struct {
	Qualification    string `json:"qualification"`
	Department       string `json:"department"`
	InstituteName    string `json:"institute_name"`
	InstituteCountry string `json:"institute_country"`
	GraduateYear     int64  `json:"graduate_year"`
}
