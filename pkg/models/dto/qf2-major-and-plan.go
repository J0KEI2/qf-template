package dto

type ProgramMajorAndPLanGetResponseDto struct {
	ProgramMajor []ProgramMajorDto `json:"majors"`
}

type ProgramMajorAndPlanDeleteRequest struct {
	Majors   []uint `json:"majors"`
	Plans    []uint `json:"plans"`
	Subplans []uint `json:"subplans"`
}
