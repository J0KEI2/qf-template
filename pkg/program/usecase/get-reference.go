package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) GetReference(ProgramID uuid.UUID, paginationOptions *models.PaginationOptions) (result dto.GetReferenceResponseDto, err error) {
	queryDb := programQuery.ProgramReferenceQueryEntity{
		ProgramID: &ProgramID,
	}

	references := []programQuery.ProgramReferenceQueryEntity{}

	err = u.CommonRepository.GetList(queryDb, &references, paginationOptions, "ReferenceOption")
	if err != nil {
		return
	}

	referencesResult := make([]dto.ProgramReferenceResponseDto, 0)
	for _, reference := range references {
		fileQuery := query.MapFilesSystemQueryEntity{
			ReferenceID: reference.ID,
		}
		if err = u.CommonRepository.GetFirst(&fileQuery, "FileSystem"); err != nil {
			if err == gorm.ErrRecordNotFound {
				referencesResult = append(referencesResult, dto.ProgramReferenceResponseDto{
					ID:                   reference.ID,
					ProgramID:            reference.ProgramID,
					ReferenceName:        reference.ReferenceName,
					ReferenceDescription: reference.ReferenceDescription,
					ReferenceFileName:    nil,
					ReferenceFileID:      nil,
					ReferenceTypeID:      reference.ReferenceOption.ID,
					ReferenceTypeName:    reference.ReferenceOption.Name,
					CreatedAt:            reference.CreatedAt,
					UpdatedAt:            reference.UpdatedAt,
					DeletedAt:            &reference.DeletedAt,
				})

				continue
			} else {
				return dto.GetReferenceResponseDto{}, err
			}
		}

		referencesResult = append(referencesResult, dto.ProgramReferenceResponseDto{
			ID:                   reference.ID,
			ProgramID:            reference.ProgramID,
			ReferenceName:        reference.ReferenceName,
			ReferenceDescription: reference.ReferenceDescription,
			ReferenceFileName:    fileQuery.FileSystem.FileName,
			ReferenceFileID:      fileQuery.FileSystem.ID,
			ReferenceTypeID:      reference.ReferenceOption.ID,
			ReferenceTypeName:    reference.ReferenceOption.Name,
			CreatedAt:            reference.CreatedAt,
			UpdatedAt:            reference.UpdatedAt,
			DeletedAt:            &reference.DeletedAt,
		})
	}

	result = dto.GetReferenceResponseDto{
		Items:             referencesResult,
		PaginationOptions: paginationOptions,
	}

	return result, nil
}
