package dto

type ProgramQualityAssurance struct {
	HESC  QualityAssuranceData `json:"hesc"`
	AunQa QualityAssuranceData `json:"aun_qa"`
	ABET  QualityAssuranceData `json:"abet"`
	WFME  QualityAssuranceData `json:"wfme"`
	AACSB QualityAssuranceData `json:"aacsb"`
}

type CreateOrUpdateQualityAssuranceDto struct {
	HESC  QualityAssuranceData `json:"hesc"`
	AunQa QualityAssuranceData `json:"aun_qa"`
	ABET  QualityAssuranceData `json:"abet"`
	WFME  QualityAssuranceData `json:"wfme"`
	AACSB QualityAssuranceData `json:"aacsb"`
}

type QualityAssuranceData struct {
	IsCheck     *bool   `json:"is_checked"`
	Description *string `json:"description"`
}
