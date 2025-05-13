package query

import (
	"time"

	"github.com/google/uuid"
)

type CommentQueryEntity struct {
	ID    uint      `json:"id"`
	QFUID uuid.UUID `json:"qf_uid"`

	QFType       string `json:"qf_type"`
	CategoryType string `json:"category_type"`
	Attribute    string `json:"attribute"`
	Resolve      bool   `json:"resolve"`
}

func (c *CommentQueryEntity) TableName() string {
	return "comments"
}

type CommentCreateEntity struct {
	ID           uint      `json:"id"`
	QFType       string    `json:"qf_type"`
	QFUID        uuid.UUID `json:"qf_uid"`
	CategoryType string    `json:"category_type"`
	Attribute    string    `json:"attribute"`
	Commentator  string    `json:"commentator"`
	Comments     string    `json:"comments"`
	Resolve      bool      `json:"resolve"`
	UpdatedBy    uuid.UUID `json:"updated_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}

func (c *CommentCreateEntity) TableName() string {
	return "comments"
}

type CommentUpdateEntity struct {
	ID           uint      `json:"id"`
	QFType       string    `json:"qf_type"`
	QFUID        uuid.UUID `json:"qf_uid"`
	CategoryType string    `json:"category_type"`
	Attribute    string    `json:"attribute"`
	Commentator  string    `json:"commentator"`
	Comments     string    `json:"comments"`
	Resolve      bool      `json:"resolve"`
	UpdatedBy    uuid.UUID `json:"updated_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}

func (c *CommentUpdateEntity) TableName() string {
	return "comments"
}
