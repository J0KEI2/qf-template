package dto

type GetOpenablePageResponseDto struct {
	Items []OpenablePageDto `json:"items"`
}

type OpenablePageDto struct {
	PageName          string `json:"page_name"`
	Openable          bool   `json:"openable"`
	OpenableBySubPlan []int  `json:"openable_by_sub_plan"`
}
