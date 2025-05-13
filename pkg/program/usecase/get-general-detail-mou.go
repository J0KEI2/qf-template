package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u programUsecase) GetGeneralDetailMou(generalDetailID uint) (result dto.GetMouFileResponseDto, err error) {
	fileQuery := query.MapFilesSystemQueryEntity{
		GeneralDetailID: &generalDetailID,
	}
	mouFile := []query.MapFilesSystemQueryEntity{}
	if err = u.CommonRepository.GetList(&fileQuery, &mouFile, nil, "FileSystem"); err != nil {
		return dto.GetMouFileResponseDto{}, err
	}

	mouFileArray := make([]dto.MouFileResponse, 0)
	for _, mouData := range mouFile {
		mouFileArray = append(mouFileArray, dto.MouFileResponse{
			FileID:   mouData.FileID,
			FileName: mouData.FileSystem.FileName,
		})
	}

	result = dto.GetMouFileResponseDto{
		Items: mouFileArray,
	}

	return result, nil
}
