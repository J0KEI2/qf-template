package dto

import (
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

type ProgramPLOGetResponseDto struct {
	ID               uint                  `json:"id"`
	ProgramSubPlanID uint                  `json:"program_sub_plan_id"`
	PLOFormat        *string               `json:"plo_format"`
	Ksec             *KsecResponseDto      `json:"ksec"`
	PLODetails       []ProgramPLODetailDto `json:"plo_details"`
}

type ProgramMapPloWithKsec struct {
	ID        *uint           `json:"id"`
	PloID     *uint           `json:"plo_id"`
	KsecID    *uint           `json:"ksec_id"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

type CreateOrUpdatePLORequestDto struct {
	ID            uint                  `json:"id"`
	ProgramMainID uuid.UUID             `json:"program_main_id"`
	PLOFormat     *string               `json:"plo_format"`
	Ksec          KsecRequestDto        `json:"ksec"`
	PLODetails    []ProgramPLODetailDto `json:"plo_details"`
}

type ProgramPLODetailDto struct {
	ID        *uint                 `json:"id"`
	Order     *uint                 `json:"order"`
	ParentID  *uint                 `json:"parent_id"`
	Children  []ProgramPLODetailDto `json:"children"`
	PLOPrefix *string               `json:"plo_prefix"`
	PLODetail *string               `json:"plo_detail"`
	CreatedAt *time.Time            `json:"created_at"`
	UpdatedAt *time.Time            `json:"updated_at"`
	DeletedAt *gorm.DeletedAt       `json:"deleted_at,omitempty"`
}

type KsecResponseDto struct {
	Knowledge []KsecDetail `json:"knowledge"`
	Skill     []KsecDetail `json:"skill"`
	Ethic     []KsecDetail `json:"ethic"`
	Character []KsecDetail `json:"character"`
}

func (ksec *KsecResponseDto) Init() {
	ksec.Knowledge = make([]KsecDetail, 0)
	ksec.Skill = make([]KsecDetail, 0)
	ksec.Ethic = make([]KsecDetail, 0)
	ksec.Character = make([]KsecDetail, 0)
}

func (ksec *KsecRequestDto) Init() {
	ksec.Knowledge = make([]KsecDetail, 0)
	ksec.Skill = make([]KsecDetail, 0)
	ksec.Ethic = make([]KsecDetail, 0)
	ksec.Character = make([]KsecDetail, 0)
}

func (ksec *KsecResponseDto) SortByOrder() {
	sort.SliceStable(ksec.Knowledge, func(i, j int) bool {
		return *ksec.Knowledge[i].Order < *ksec.Knowledge[j].Order
	})
	sort.SliceStable(ksec.Skill, func(i, j int) bool {
		return *ksec.Skill[i].Order < *ksec.Skill[j].Order
	})
	sort.SliceStable(ksec.Ethic, func(i, j int) bool {
		return *ksec.Ethic[i].Order < *ksec.Ethic[j].Order
	})
	sort.SliceStable(ksec.Character, func(i, j int) bool {
		return *ksec.Character[i].Order < *ksec.Character[j].Order
	})
}

func (ksec *KsecRequestDto) SortByOrder() {
	sort.SliceStable(ksec.Knowledge, func(i, j int) bool {
		return *ksec.Knowledge[i].Order < *ksec.Knowledge[j].Order
	})
	sort.SliceStable(ksec.Skill, func(i, j int) bool {
		return *ksec.Skill[i].Order < *ksec.Skill[j].Order
	})
	sort.SliceStable(ksec.Ethic, func(i, j int) bool {
		return *ksec.Ethic[i].Order < *ksec.Ethic[j].Order
	})
	sort.SliceStable(ksec.Character, func(i, j int) bool {
		return *ksec.Character[i].Order < *ksec.Character[j].Order
	})
}

func (ksec *KsecResponseDto) FromKsecDetail(ksecDetailArray []query.ProgramKsecDetailQueryEntity) {
	for _, ksecDetail := range ksecDetailArray {
		if ksecDetail.Type == nil {
			continue
		}
		switch *ksecDetail.Type {
		case "K":
			ksec.Knowledge = append(ksec.Knowledge, newKsecDetailFromQueryKsecDetail(ksecDetail))
		case "S":
			ksec.Skill = append(ksec.Skill, newKsecDetailFromQueryKsecDetail(ksecDetail))
		case "E":
			ksec.Ethic = append(ksec.Ethic, newKsecDetailFromQueryKsecDetail(ksecDetail))
		case "C":
			ksec.Character = append(ksec.Character, newKsecDetailFromQueryKsecDetail(ksecDetail))
		default:
			continue
		}
	}
}

func (ksec *KsecRequestDto) FromKsecDetail(ksecDetailArray []query.ProgramKsecDetailQueryEntity) {
	for _, ksecDetail := range ksecDetailArray {
		if ksecDetail.Type == nil {
			continue
		}
		switch *ksecDetail.Type {
		case "K":
			ksec.Knowledge = append(ksec.Knowledge, newDupKsecDetailFromQueryKsecDetail(ksecDetail))
		case "S":
			ksec.Skill = append(ksec.Skill, newDupKsecDetailFromQueryKsecDetail(ksecDetail))
		case "E":
			ksec.Ethic = append(ksec.Ethic, newDupKsecDetailFromQueryKsecDetail(ksecDetail))
		case "C":
			ksec.Character = append(ksec.Character, newDupKsecDetailFromQueryKsecDetail(ksecDetail))
		default:
			continue
		}
	}
}

type KsecRequestDto struct {
	Knowledge []KsecDetail `json:"knowledge"`
	Skill     []KsecDetail `json:"skill"`
	Ethic     []KsecDetail `json:"ethic"`
	Character []KsecDetail `json:"character"`
}

func (ksec *KsecRequestDto) ToKsecDetail() []KsecDetail {
	ksecDetails := make([]KsecDetail, 0)
	ksecDetails = append(ksecDetails, formatKsecType(ksec.Knowledge, "K")...)
	ksecDetails = append(ksecDetails, formatKsecType(ksec.Skill, "S")...)
	ksecDetails = append(ksecDetails, formatKsecType(ksec.Ethic, "E")...)
	ksecDetails = append(ksecDetails, formatKsecType(ksec.Character, "C")...)
	return ksecDetails
}

func formatKsecType(ksecDetails []KsecDetail, Key string) []KsecDetail {
	for index := range ksecDetails {
		ksecDetails[index].Type = &Key
	}
	return ksecDetails
}

type KsecDetail struct {
	ID        *uint      `json:"id"`
	MapPLOID  *uint      `json:"map_plo_id"`
	Type      *string    `json:"type"`
	Order     *uint      `json:"order"`
	Detail    *string    `json:"detail"`
	IsChecked *bool      `json:"is_checked"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func newKsecDetailFromQueryKsecDetail(queryKsecDetail query.ProgramKsecDetailQueryEntity) KsecDetail {
	return KsecDetail{
		ID:        queryKsecDetail.ID,
		MapPLOID:  queryKsecDetail.ProgramMapPloWithKsecID,
		Type:      queryKsecDetail.Type,
		Order:     queryKsecDetail.Order,
		Detail:    queryKsecDetail.Detail,
		IsChecked: queryKsecDetail.IsChecked,
		CreatedAt: queryKsecDetail.CreatedAt,
		UpdatedAt: queryKsecDetail.UpdatedAt,
	}
}

func newDupKsecDetailFromQueryKsecDetail(queryKsecDetail query.ProgramKsecDetailQueryEntity) KsecDetail {
	return KsecDetail{
		// ID:        queryKsecDetail.ID,
		// MapPLOID:  queryKsecDetail.ProgramMapPloWithKsecID,
		Type:      queryKsecDetail.Type,
		Order:     queryKsecDetail.Order,
		Detail:    queryKsecDetail.Detail,
		IsChecked: queryKsecDetail.IsChecked,
		CreatedAt: queryKsecDetail.CreatedAt,
		UpdatedAt: queryKsecDetail.UpdatedAt,
	}
}

type GetProgramPLOMapKsecResponseDto struct {
	ID               uint                  `json:"id"`
	ProgramSubPlanID uint                  `json:"program_sub_plan_id"`
	PLOFormat        *string               `json:"plo_format"`
	Ksec             *KsecResponseDto      `json:"ksec"`
	PLODetails       []ProgramPLODetailDto `json:"plo_details"`
}

type LearningSolutionResponseDto struct {
	Items []LearningSolution `json:"items"`
	*models.PaginationOptions
}

type LearningSolution struct {
	ID        *uint      `json:"id"`
	PloID     *uint      `json:"plo_id"`
	Order     *uint      `json:"order"`
	Detail    *string    `json:"detail"`
	Key       *string    `json:"key"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateOrUpdateLearningSolutionRequestDto struct {
	ID        *uint      `json:"id"`
	PloID     *uint      `json:"plo_id"`
	Order     *uint      `json:"order"`
	Detail    *string    `json:"detail"`
	Key       *string    `json:"key"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type LearningEvaluationResponseDto struct {
	Items []LearningEvaluation `json:"items"`
	*models.PaginationOptions
}

type LearningEvaluation struct {
	ID        *uint      `json:"id"`
	PloID     *uint      `json:"plo_id"`
	Order     *uint      `json:"order"`
	Detail    *string    `json:"detail"`
	Key       *string    `json:"key"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateOrUpdateLearningEvaluationRequestDto struct {
	ID        *uint      `json:"id"`
	PloID     *uint      `json:"plo_id"`
	Order     *uint      `json:"order"`
	Detail    *string    `json:"detail"`
	Key       *string    `json:"key"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type MapKsecList struct {
	KsecList KsecRequestDto
	Children []MapKsecList
}

type LearningEvaList struct {
	LEList   CreateOrUpdateLearningEvaluationRequestDto
	Children []LearningEvaList
}

type LearningSolList struct {
	LSList   CreateOrUpdateLearningSolutionRequestDto
	Children []LearningSolList
}
