package approvals

import (
	"sort"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u *programApprovalUsecase) GetAllApproval(programUID *uuid.UUID) (*[]dto.GetAllApprovalDto, error) {
	var record []query.ProgramApprovalQueryEntity

	if err := u.CommonRepository.GetList(&query.ProgramApprovalQueryEntity{ProgramUID: *programUID}, &record, nil); err != nil {
		return nil, err
	}

	var result []dto.GetAllApprovalDto

	for _, approval := range record {
		result = append(result, dto.GetAllApprovalDto{
			ID:            approval.ID,
			ProgramUID:    approval.ProgramUID,
			NameTh:        approval.NameTh,
			NameEn:        approval.NameEn,
			ApprovalLevel: approval.ApprovalStatusLevel,
			IsCurrent:     approval.IsCurrent,
			IsApproved:    approval.IsApproved,
			IsRejected:    approval.IsRejected,
			CreatedAt:     approval.CreatedAt,
			UpdatedAt:     approval.UpdatedAt,
		})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return &result, nil
}
