package entity

import (
	"time"

	"github.com/google/uuid"
)

type CommentFetchQueryEntity struct {
	ID           uint      `json:"id"`
	QFType       string    `json:"qf_type"`
	QFUID        uuid.UUID `json:"qf_uid"`
	CategoryType string    `json:"category_type"`
	Attribute    string    `json:"attribute"`
}

type CommentFetchEntity struct {
	ID           uint      `json:"id"`
	QFType       string    `json:"qf_type"`
	QFUID        uuid.UUID `json:"qf_uid"`
	CategoryType string    `json:"category_type"`
	Attribute    string    `json:"attribute"`
	Commentator  string    `json:"commentator"`
	Comments     string    `json:"comments"`
	Resolve      bool      `json:"resolve"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
}

func (c *CommentFetchEntity) TableName() string {
	return "comments"
}

type CommentUpdateReusltEntity struct {
	ID           *uint      `json:"id"`
	QFType       *string    `json:"qf_type"`
	QFUID        *uuid.UUID `json:"qf_uid"`
	CategoryType *string    `json:"category_type"`
	Attribute    *string    `json:"attribute"`
	Commentator  *string    `json:"commentator"`
	Comments     *string    `json:"comments"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	UpdatedBy    *uuid.UUID `json:"updated_by"`
}

type CommentCreateEntity struct {
	QFType       string    `json:"qf_type"`
	QFUID        uuid.UUID `json:"qf_uid"`
	CategoryType string    `json:"category_type"`
	Attribute    string    `json:"attribute"`
	Commentator  string    `json:"commentator"`
	Comments     string    `json:"comments"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
}

type CommentCreateResultEntity struct {
	ID           uint      `json:"id"`
	QFType       string    `json:"qf_type"`
	QFUID        uuid.UUID `json:"qf_uid"`
	CategoryType string    `json:"category_type"`
	Attribute    string    `json:"attribute"`
	Commentator  string    `json:"commentator"`
	Comments     string    `json:"comments"`
	Resolve      bool      `json:"resolve"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	UpdatedBy    uuid.UUID `json:"updated_by,omitempty"`
}

type CommentUpdateQueryEntity struct {
	ID uint `json:"id"`
}

type CommentUpdateEntity struct {
	QFType       *string    `json:"qf_type"`
	QFUID        *uuid.UUID `json:"qf_uid"`
	CategoryType *string    `json:"category_type"`
	Attribute    *string    `json:"attribute"`
	Commentator  *string    `json:"commentator"`
	Comments     *string    `json:"comments"`
	Resolve      bool       `json:"resolve"`

	UpdatedBy *uuid.UUID `json:"updated_by"`
}

type CommentUpdateResultEntity struct {
	ID           uint       `json:"id"`
	QFType       *string    `json:"qf_type"`
	QFUID        *uuid.UUID `json:"qf_uid"`
	CategoryType *string    `json:"category_type"`
	Attribute    *string    `json:"attribute"`
	Commentator  *string    `json:"commentator"`
	Comments     *string    `json:"comments"`
	Resolve      bool       `json:"resolve"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	UpdatedBy    *uuid.UUID `json:"updated_by"`
}
