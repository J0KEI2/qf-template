package query

import "github.com/google/uuid"

type RoleSystems struct {
	RoleNameTH      *string
	RoleNameEN      *string
	UID             *uuid.UUID
	ViewMyProfile   *bool
	UpdateMyProfile *bool
	ListAllLecture  *bool
	CreateLecture   *bool
	UpdateLecture   *bool
	DeleteLecture   *bool
	Status          *bool
}

func (s *RoleSystems) TableName() string {
	return "role_systems"
}

type RoleSystemJoinQuery struct {
	RoleNameTH      *string
	RoleNameEN      *string
	UID             *uuid.UUID
	ViewMyProfile   *bool
	UpdateMyProfile *bool
	ListAllLecture  *bool
	CreateLecture   *bool
	UpdateLecture   *bool
	DeleteLecture   *bool
	Status          *bool   
}

func (s *RoleSystemJoinQuery) TableName() string {
	return "role_systems"
}

func (s RoleSystemJoinQuery) String() string {
	return "SystemRole"
}

