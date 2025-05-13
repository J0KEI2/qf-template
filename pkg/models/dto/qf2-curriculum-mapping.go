package dto

type ProgramCurMapRespGetResponseDto struct {
	ProgramSubPlanID          uint              `json:"program_sub_plan_id"`
	ProgramFundamentalDetails ProgramCurMapResp `json:"program_fundamental_details"`
	ProgramCompulsoryDetails  ProgramCurMapResp `json:"program_compulsory_details"`
	ProgramEnrichmentDetails  ProgramCurMapResp `json:"program_enrichment_details"`
}

type CreateOrUpdateCurMapRespRequestDto struct {
	ProgramSubPlanID          uint              `json:"program_sub_plan_id"`
	ProgramFundamentalDetails ProgramCurMapResp `json:"program_fundamental_details"`
	ProgramCompulsoryDetails  ProgramCurMapResp `json:"program_compulsory_details"`
	ProgramEnrichmentDetails  ProgramCurMapResp `json:"program_enrichment_details"`
}

type ProgramCurMapResp struct {
	ProgramHeaderNameEN *string                     `json:"program_header_name_en,omitempty"`
	ProgramHeaderNameTH *string                     `json:"program_header_name_th,omitempty"`
	ProgramCourseType   *int                        `json:"program_course_type,omitempty"`
	Items               []ProgramCurMapRespResponse `json:"items,omitempty"`
}

type ProgramCurMapRespResponse struct {
	ID                       uint                       `json:"id"`
	ProgramCourseID          *uint                      `json:"program_course_id,omitempty"`
	CourseNameTH             *string                    `json:"course_name_th,omitempty"`
	CourseNameEN             *string                    `json:"course_name_en,omitempty"`
	ProgramCurMapRespDetails []ProgramCurMapRespDetails `json:"program_curmap_resp_details,omitempty"`
}

type ProgramCurMapRespDetails struct {
	CurMapID         *uint   `json:"cur_map_id"`
	PLOID            *uint   `json:"plo_id"`
	ProgramPloPrefix *string `json:"plo_prefix,omitempty"`
	Status           *int    `json:"status"`
}

type ProgramCurMapDetail struct {
	ID                  *uint   `json:"id"`
	ProgramCourseID     *uint   `json:"program_course_id"`
	ProgramCourseType   *int    `json:"program_course_type"`
	CourseNameTH        *string `json:"course_name_th"`
	CourseNameEN        *string `json:"course_name_en"`
	ProgramHeaderNameTH *string `json:"program_header_name_th"`
	ProgramHeaderNameEN *string `json:"program_header_name_en"`
	PLOID               *uint   `json:"plo_id"`
	ProgramPloPrefix    *string `json:"plo_prefix"`
	Status              *int    `json:"status"`
	Order               *uint   `json:"order"`
}

type ProgramCurMapKsaGetResponseDto struct {
	ProgramSubPlanID          uint             `json:"program_sub_plan_id"`
	ProgramFundamentalDetails ProgramCurMapKsa `json:"program_fundamental_details"`
	ProgramCompulsoryDetails  ProgramCurMapKsa `json:"program_compulsory_details"`
	ProgramEnrichmentDetails  ProgramCurMapKsa `json:"program_enrichment_details"`
}

type CreateOrUpdateCurMapKsaRequestDto struct {
	ProgramSubPlanID          uint             `json:"program_sub_plan_id"`
	ProgramFundamentalDetails ProgramCurMapKsa `json:"program_fundamental_details"`
	ProgramCompulsoryDetails  ProgramCurMapKsa `json:"program_compulsory_details"`
	ProgramEnrichmentDetails  ProgramCurMapKsa `json:"program_enrichment_details"`
}

type ProgramCurMapKsa struct {
	ProgramHeaderNameEN *string                    `json:"program_header_name_en,omitempty"`
	ProgramHeaderNameTH *string                    `json:"program_header_name_th,omitempty"`
	ProgramCourseType   *int                       `json:"program_course_type,omitempty"`
	Items               []ProgramCurMapKsaResponse `json:"items,omitempty"`
}

type ProgramCurMapKsaResponse struct {
	ID                      uint                      `json:"id"`
	ProgramCourseID         *uint                     `json:"program_course_id,omitempty"`
	CourseNameTH            *string                   `json:"course_name_th,omitempty"`
	CourseNameEN            *string                   `json:"course_name_en,omitempty"`
	ProgramCurMapKsaDetails []ProgramCurMapKsaDetails `json:"program_curmap_ksa_details,omitempty"`
}

type ProgramCurMapKsaDetails struct {
	CurMapID         *uint   `json:"cur_map_id"`
	PLOID            *uint   `json:"plo_id"`
	ProgramPloPrefix *string `json:"plo_prefix,omitempty"`
	KsaID            []int   `json:"ksa_id"`
}

type ProgramCurMapKsaDetail struct {
	ID                  *uint   `json:"id"`
	ProgramCourseID     *uint   `json:"program_course_id"`
	ProgramCourseType   *int    `json:"program_course_type"`
	CourseNameTH        *string `json:"course_name_th"`
	CourseNameEN        *string `json:"course_name_en"`
	ProgramHeaderNameTH *string `json:"program_header_name_th"`
	ProgramHeaderNameEN *string `json:"program_header_name_en"`
	PLOID               *uint   `json:"plo_id"`
	ProgramPloPrefix    *string `json:"plo_prefix"`
	KsaID               []int   `json:"ksa_id"`
}

type ProgramKsaDetailGetResponseDto struct {
	ProgramSubPlanID uint               `json:"program_sub_plan_id"`
	Knowledge        []ProgramKsaDetail `json:"knowledge"`
	Skill            []ProgramKsaDetail `json:"skill"`
	Attitude         []ProgramKsaDetail `json:"attitude"`
}

type CreateOrUpdateKsaDetailRequestDto struct {
	ProgramSubPlanID uint               `json:"program_sub_plan_id"`
	Knowledge        []ProgramKsaDetail `json:"knowledge"`
	Skill            []ProgramKsaDetail `json:"skill"`
	Attitude         []ProgramKsaDetail `json:"attitude"`
}

type ProgramKsaDetail struct {
	ID        *uint   `json:"id"`
	KsaType   *string `json:"ksa_type"`
	ShortCode *string `json:"short_code"`
	KsaDetail *string `json:"ksa_detail"`
	Order     *uint   `json:"order"`
}
