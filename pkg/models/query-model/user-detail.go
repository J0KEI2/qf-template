package query

import (
	"time"

	"github.com/google/uuid"
)

type UserDetail struct {
	UpdatedAt           *time.Time 
	CreatedAt           *time.Time 
	MiddlenameEn        *string    
	EducationBackGround *string    
	MiddlenameTh        *string    
	Position            *string    
	TitleTh             *string    
	FirstnameEn         *string    
	TitleEn             *string    
	LastnameEn          *string    
	LastnameTh          *string    
	FirstnameTh         *string    
	UID                 *uuid.UUID 
	UserUID             *uuid.UUID 
}

type UserDetailQueryEntity struct {
	UpdatedAt           *time.Time 
	CreatedAt           *time.Time 
	MiddlenameEn        *string    
	EducationBackGround *string    
	MiddlenameTh        *string    
	Position            *string    
	TitleTh             *string    
	FirstnameEn         *string    
	TitleEn             *string    
	LastnameEn          *string    
	LastnameTh          *string    
	FirstnameTh         *string    
	UID                 *uuid.UUID 
	UserUID             *uuid.UUID 
}

type UserDetailJoinQuery struct {
	UpdatedAt           *time.Time 
	CreatedAt           *time.Time 
	MiddlenameEn        *string    
	EducationBackGround *string    
	MiddlenameTh        *string    
	Position            *string    
	TitleTh             *string    
	FirstnameEn         *string    
	TitleEn             *string    
	LastnameEn          *string    
	LastnameTh          *string    
	FirstnameTh         *string    
	UID                 *uuid.UUID 
	UserUID             *uuid.UUID 
}

func (s *UserDetailQueryEntity) TableName() string {
	return "user_details"
}

func (s UserDetailJoinQuery) String() string {
	return "UserDetail"
}

