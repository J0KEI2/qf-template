package query

import "github.com/google/uuid"

type Faculties struct {
	FacultyNameTH *string
	FacultyNameEN *string
	University    *string
	FacultyID     *string
	ID            *uuid.UUID
}

type FacultiesJoinQuery struct {
	FacultyNameTH *string
	FacultyNameEN *string
	University    *string
	FacultyID     *string
	ID            *uuid.UUID
}

func (s *FacultiesJoinQuery) TableName() string {
	return "faculties"
}

func (s FacultiesJoinQuery) String() string {
	return "Faculty"
}
