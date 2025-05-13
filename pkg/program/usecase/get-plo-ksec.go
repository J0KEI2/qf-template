package usecase

import (
	"sort"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetPlo(ProgramSubPlanId uint) (result dto.ProgramPLOGetResponseDto, err error) {
	programPLO := query.ProgramPloFormatQueryEntity{
		ProgramSubPlanID: &ProgramSubPlanId,
	}

	if err = u.CommonRepository.GetFirst(&programPLO); err != nil {
		return
	}

	programKsecQuery := query.ProgramKsecDetailQueryEntity{
		ProgramSubPlanID: &ProgramSubPlanId,
	}

	programKsec := []query.ProgramKsecDetailQueryEntity{}

	if err = u.CommonRepository.GetList(&programKsecQuery, &programKsec, nil); err != nil {
		return
	}

	ploDetails, err := u.RecursivePloDetail(programPLO.ID, pointer.ToUint(0))

	ksec := new(dto.KsecResponseDto)
	ksec.Init()
	ksec.FromKsecDetail(programKsec)
	ksec.SortByOrder()

	result = dto.ProgramPLOGetResponseDto{
		ID:               *programPLO.ID,
		ProgramSubPlanID: ProgramSubPlanId,
		PLOFormat:        programPLO.PLOFormat,
		PLODetails:       ploDetails,
		Ksec:             ksec,
	}

	return
}

func (u programUsecase) RecursivePloDetail(PloFormatID *uint, parentID *uint) (result []dto.ProgramPLODetailDto, err error) {
	queryPLO := query.ProgramPloQueryEntity{
		ProgramPloFormatID: PloFormatID,
		ParentID:           parentID,
	}

	ploDetails := []query.ProgramPloQueryEntity{}

	err = u.CommonRepository.GetListWithNilSearch(&queryPLO, &ploDetails, nil)

	result = make([]dto.ProgramPLODetailDto, 0)

	for _, ploDetail := range ploDetails {
		chlid, err := u.RecursivePloDetail(PloFormatID, ploDetail.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, dto.ProgramPLODetailDto{
			ID:        ploDetail.ID,
			Order:     ploDetail.Order,
			ParentID:  ploDetail.ParentID,
			PLODetail: ploDetail.PLODetail,
			PLOPrefix: ploDetail.PLOPrefix,
			Children:  chlid,
			CreatedAt: ploDetail.CreatedAt,
			UpdatedAt: ploDetail.UpdatedAt,
			DeletedAt: ploDetail.DeletedAt,
		})
	}
	sort.SliceStable(result, func(i, j int) bool {
		return *result[i].Order < *result[j].Order
	})
	return
}

func (u programUsecase) RecursiveDupPloDetail(ploFormatID *uint, parentID *uint, subPlanId *uint) ([]dto.ProgramPLODetailDto, []dto.MapKsecList, error) {
	queryPLO := query.ProgramPloQueryEntity{
		ProgramPloFormatID: ploFormatID,
		ParentID:           parentID,
	}

	ploDetails := []query.ProgramPloQueryEntity{}

	u.CommonRepository.GetList(&queryPLO, &ploDetails, nil)

	result := make([]dto.ProgramPLODetailDto, 0)
	mapKsecResult := make([]dto.MapKsecList, 0)

	for _, ploDetail := range ploDetails {
		programPloQuery := query.ProgramPloQueryEntity{
			ID: ploDetail.ID,
		}

		if err := u.CommonRepository.GetFirst(&programPloQuery, "ProgramMapPloWithKsecQueryEntity", "LearningEvaluation", "LearningEvaluation"); err != nil {
			return nil, nil, err
		}

		programKsecQuery := query.ProgramKsecDetailQueryEntity{
			ProgramSubPlanID: subPlanId,
		}

		programKsec := []query.ProgramKsecDetailQueryEntity{}

		if err := u.CommonRepository.GetList(&programKsecQuery, &programKsec, nil); err != nil {
			return nil, nil, err
		}

		programKsec = checkKsecCheckedWithPLO(programKsec, programPloQuery.ProgramMapPloWithKsecQueryEntity)

		mapKsec := new(dto.KsecRequestDto)
		mapKsec.Init()
		mapKsec.FromKsecDetail(programKsec)
		mapKsec.SortByOrder()

		chlid, mapKsecList, err := u.RecursiveDupPloDetail(ploFormatID, ploDetail.ID, subPlanId)
		if err != nil {
			return nil, nil, err
		}
		result = append(result, dto.ProgramPLODetailDto{
			// ID:        ploDetail.ID,
			Order: ploDetail.Order,
			// ParentID:  ploDetail.ParentID,
			PLODetail: ploDetail.PLODetail,
			PLOPrefix: ploDetail.PLOPrefix,
			Children:  chlid,
			// CreatedAt: ploDetail.CreatedAt,
			// UpdatedAt: ploDetail.UpdatedAt,
			// DeletedAt: ploDetail.DeletedAt,
		})

		mapKsecResult = append(mapKsecResult, dto.MapKsecList{
			KsecList: *mapKsec,
			Children: mapKsecList,
		})

	}
	sort.SliceStable(result, func(i, j int) bool {
		return *result[i].Order < *result[j].Order
	})
	return result, mapKsecResult, nil
}
