package usecase

import (
	"errors"
	"sort"
	"strings"

	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetKSADetail(subPlanID uint) (result *dto.ProgramKsaDetailGetResponseDto, err error) {
	programKsaDetail := query.ProgramKsaDetailQueryEntity{
		ProgramSubPlanID: &subPlanID,
	}

	ksaDetailsDest := []query.ProgramKsaDetailQueryEntity{}
	err = u.CommonRepository.GetList(&programKsaDetail, &ksaDetailsDest, nil)
	if err != nil {
		return nil, err
	}

	if len(ksaDetailsDest) == 0 {
		err = helper.ExecuteTransaction(u.CommonRepository, u.InitKsaDetail(subPlanID))
		if err != nil {
			return nil, err
		}

		err = u.CommonRepository.GetList(&programKsaDetail, &ksaDetailsDest, nil)
		if err != nil {
			return nil, err
		}
	}

	knowledgeList := []dto.ProgramKsaDetail{}
	skillList := []dto.ProgramKsaDetail{}
	attitudeList := []dto.ProgramKsaDetail{}

	for _, ksaItem := range ksaDetailsDest {
		if strings.EqualFold(*ksaItem.KsaType, "k") {
			knowledgeList = append(knowledgeList, dto.ProgramKsaDetail{
				ID:        ksaItem.ID,
				KsaType:   ksaItem.KsaType,
				ShortCode: ksaItem.ShortCode,
				KsaDetail: ksaItem.KsaDetail,
				Order:     ksaItem.Order,
			})

			sort.SliceStable(knowledgeList, func(i, j int) bool {
				return *knowledgeList[i].Order < *knowledgeList[j].Order
			})
		} else if strings.EqualFold(*ksaItem.KsaType, "s") {
			skillList = append(skillList, dto.ProgramKsaDetail{
				ID:        ksaItem.ID,
				KsaType:   ksaItem.KsaType,
				ShortCode: ksaItem.ShortCode,
				KsaDetail: ksaItem.KsaDetail,
				Order:     ksaItem.Order,
			})

			sort.SliceStable(skillList, func(i, j int) bool {
				return *skillList[i].Order < *skillList[j].Order
			})
		} else if strings.EqualFold(*ksaItem.KsaType, "a") {
			attitudeList = append(attitudeList, dto.ProgramKsaDetail{
				ID:        ksaItem.ID,
				KsaType:   ksaItem.KsaType,
				ShortCode: ksaItem.ShortCode,
				KsaDetail: ksaItem.KsaDetail,
				Order:     ksaItem.Order,
			})

			sort.SliceStable(attitudeList, func(i, j int) bool {
				return *attitudeList[i].Order < *attitudeList[j].Order
			})
		} else {
			return nil, errors.New("ksa Type not found")
		}
	}

	result = &dto.ProgramKsaDetailGetResponseDto{
		ProgramSubPlanID: subPlanID,
		Knowledge:        knowledgeList,
		Skill:            skillList,
		Attitude:         attitudeList,
	}

	return
}
