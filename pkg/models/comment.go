package models

import (
	"time"

	"github.com/google/uuid"
)

type CommentFetchModel struct {
	ID           uint      `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	QFType       string    `gorm:"column:qf_type;type:varchar;size:255"`
	QFUID        uuid.UUID `gorm:"column:qf_uid;type:uuid;"`
	CategoryType string    `gorm:"column:category_type;type:varchar;size:255"`
	Attribute    string    `gorm:"column:attribute;type:varchar;size:255"`
	Commentator  string    `gorm:"column:commentator;type:varchar;size:255"`
	Comments     string    `gorm:"column:comments;type:varchar;size:255"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedBy    uuid.UUID `gorm:"column:updated_by;type:uuid;not null"`
}

type CommentCreateQuery struct {
	ID           uint      `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	QFType       string    `json:"qf_type" gorm:"column:qf_type;type:varchar;size:255"`
	QFUID        uuid.UUID `json:"qf_uid" gorm:"column:qf_uid;type:uuid;"`
	CategoryType string    `json:"category_type" gorm:"column:category_type;type:varchar;size:255"`
	Attribute    string    `json:"attribute" gorm:"column:attribute;type:varchar;size:255"`
	Commentator  string    `json:"commentator" gorm:"column:commentator;type:varchar;size:255"`
	Comments     string    `json:"comments" gorm:"column:comments;type:varchar;size:255"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedBy    uuid.UUID `json:"updated_by,omitempty" gorm:"column:updated_by;type:uuid;not null"`
}

type CommentUpdateModel struct {
	QFType       *string    `gorm:"column:qf_type;type:varchar;size:255"`
	QFUID        *uuid.UUID `gorm:"column:qf_uid;type:uuid;"`
	CategoryType *string    `gorm:"column:category_type;type:varchar;size:255"`
	Attribute    *string    `gorm:"column:attribute;type:varchar;size:255"`
	Commentator  *string    `gorm:"column:commentator;type:varchar;size:255"`
	Comments     *string    `gorm:"column:comments;type:varchar;size:255"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedBy    *uuid.UUID `gorm:"column:updated_by;type:uuid;not null"`
}

type CommentUpdateQuery struct {
	QFType       *string    `json:"qf_type" gorm:"column:qf_type;type:varchar;size:255"`
	QFUID        *uuid.UUID `json:"qf_uid" gorm:"column:qf_uid;type:uuid;"`
	CategoryType *string    `json:"category_type" gorm:"column:category_type;type:varchar;size:255"`
	Attribute    *string    `json:"attribute" gorm:"column:attribute;type:varchar;size:255"`
	Commentator  *string    `json:"commentator" gorm:"column:commentator;type:varchar;size:255"`
	Comments     *string    `json:"comments" gorm:"column:comments;type:varchar;size:255"`
	UpdatedBy    *uuid.UUID `json:"updated_by" gorm:"column:updated_by;type:uuid;not null"`
}
