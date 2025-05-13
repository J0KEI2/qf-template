package approvals

import (
	"sort"

	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u *programApprovalUsecase) GetCHECOByProgramUID(programUID *uuid.UUID) (result *dto.GetCHECOByUIDResponseDto, err error) {
	queryDb := query.CHECOQueryEntity{
		ProgramUID: programUID,
	}

	checoData := []query.CHECOQueryEntity{}
	if err = u.CommonRepository.GetList(&queryDb, &checoData, nil); err != nil {
		return nil, err
	}

	checoItemsResp := []dto.CHECODetailsList{}
	for _, checo := range checoData {
		fileList := []query.MapFilesSystemQueryEntity{}
		fileQuery := query.MapFilesSystemQueryEntity{
			ChecoID: checo.ID,
		}
		if err = u.CommonRepository.GetList(&fileQuery, &fileList, nil, "FileSystem"); err != nil {
			return nil, err
		}

		checoFileList := []dto.CHECOFileList{}

		for _, file := range fileList {
			checoFileList = append(checoFileList, dto.CHECOFileList{
				FileID:   file.FileSystem.ID,
				FileName: file.FileSystem.FileName,
			})
		}

		checoStatus := constant.CHECO_STATUS.String(constant.CHECO_STATUS(*checo.StatusID))

		checoItemsResp = append(checoItemsResp, dto.CHECODetailsList{
			ID:            *checo.ID,
			NameEn:        *checo.NameEN,
			NameTN:        *checo.NameTH,
			Status:        checoStatus,
			ApprovedDate:  *checo.ApprovedDate,
			CHECOFileList: checoFileList,
			CreatedAt:     *checo.CreatedAt,
		})
	}

	sort.Slice(checoItemsResp, func(i, j int) bool {
		if checoItemsResp[i].ApprovedDate.Equal(checoItemsResp[j].ApprovedDate) {
			return checoItemsResp[i].CreatedAt.Before(checoItemsResp[j].CreatedAt)
		}
		return checoItemsResp[i].ApprovedDate.Before(checoItemsResp[j].ApprovedDate)
	})

	result = &dto.GetCHECOByUIDResponseDto{
		ProgramUID: *queryDb.ProgramUID,
		Items:      checoItemsResp,
	}

	return result, nil
}
