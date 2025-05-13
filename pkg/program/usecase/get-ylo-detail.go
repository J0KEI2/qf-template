package usecase

import (
	"log"
	"sort"

	"github.com/AlekSi/pointer"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) GetYLODetail(programSubPlanID uint) (result dto.ProgramYLODetailGetResponseDto, err error) {
	yloDetails, err := u.RecursiveYLODetail(&programSubPlanID, pointer.ToUint(0))

	result = dto.ProgramYLODetailGetResponseDto{
		ProgramSubPlanID: programSubPlanID,
		YLODetails:       yloDetails,
	}

	return
}

func (u programUsecase) RecursiveYLODetail(programSubPlanID *uint, parentID *uint) (result []dto.ProgramYLODetailDto, err error) {
	queryYLODetail := query.ProgramYloKsecQueryEntity{
		ProgramSubPlanID: programSubPlanID,
	}

	yloDetails := []query.ProgramYloKsecQueryEntity{}

	err = u.CommonRepository.GetListWithNilSearch(&queryYLODetail, &yloDetails, nil, "ProgramYearAndSemester")
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = u.CreateInitYLODetail(programSubPlanID)
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
			err = u.CommonRepository.GetListWithNilSearch(&queryYLODetail, &yloDetails, nil, "ProgramYearAndSemester")
		} else {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
			return nil, err
		}
	}

	if len(yloDetails) == 0 {
		err = u.CreateInitYLODetail(programSubPlanID)
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		err = u.CommonRepository.GetList(&queryYLODetail, &yloDetails, nil, "ProgramYearAndSemester")
	}

	programPLO := query.ProgramPloFormatQueryEntity{
		ProgramSubPlanID: programSubPlanID,
	}

	if err = u.CommonRepository.GetFirst(&programPLO); err != nil {
		return
	}

	result = make([]dto.ProgramYLODetailDto, 0)

	for _, yloDetail := range yloDetails {
		ploDetails, err := u.recursivePloCheckDetail(programPLO.ID, pointer.ToUint(0), yloDetail.ProgramYearAndSemesterID, programSubPlanID)
		if err != nil {
			return nil, err
		}

		result = append(result, dto.ProgramYLODetailDto{
			ID:                       yloDetail.ID,
			ProgramYearAndSemesterID: yloDetail.ProgramYearAndSemesterID,
			Year:                     yloDetail.ProgramYearAndSemester.Year,
			Knowledge:                yloDetail.Knowledge,
			Skill:                    yloDetail.Skill,
			Ethic:                    yloDetail.Ethic,
			Character:                yloDetail.Character,
			YLOData: dto.ProgramYLODataDto{
				ID:         programPLO.ID,
				PLOFormat:  programPLO.PLOFormat,
				PLODetails: ploDetails,
			},
			CreatedAt: yloDetail.CreatedAt,
			UpdatedAt: yloDetail.UpdatedAt,
			DeletedAt: yloDetail.DeletedAt,
		})
	}

	return
}

func (u programUsecase) recursivePloCheckDetail(PloFormatID *uint, parentID *uint, yearAndSemesterID *uint, programSubPlanID *uint) (result []dto.ProgramYLOPLODto, err error) {
	queryPLO := query.ProgramPloQueryEntity{
		ProgramPloFormatID: PloFormatID,
		ParentID:           parentID,
	}

	ploDetails := []query.ProgramPloQueryEntity{}

	if err = u.CommonRepository.GetListWithNilSearch(&queryPLO, &ploDetails, nil, "ProgramMapPloWithKsecQueryEntity"); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		return nil, err
	}

	programKsecQuery := query.ProgramKsecDetailQueryEntity{
		ProgramSubPlanID: programSubPlanID,
	}

	programKsec := []query.ProgramKsecDetailQueryEntity{}

	if err = u.CommonRepository.GetList(&programKsecQuery, &programKsec, nil); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		return nil, err
	}

	yloWithPloDetails := []query.ProgramYloWithPloQueryEntity{}
	queryYLOWithPLO := query.ProgramYloWithPloQueryEntity{
		ProgramYearAndSemesterID: yearAndSemesterID,
	}

	if err = u.CommonRepository.GetList(&queryYLOWithPLO, &yloWithPloDetails, nil); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		return nil, err
	}

	yloWithKsecDetails := []query.ProgramYloWithKsecQueryEntity{}
	queryYLOWithKsec := query.ProgramYloWithKsecQueryEntity{
		ProgramYearAndSemesterID: yearAndSemesterID,
	}

	if err = u.CommonRepository.GetList(&queryYLOWithKsec, &yloWithKsecDetails, nil); err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		return nil, err
	}

	result = make([]dto.ProgramYLOPLODto, 0)

	for _, ploDetail := range ploDetails {
		programYLOWithKsec := filterKsecOnlyInWithPLO(programKsec, ploDetail.ProgramMapPloWithKsecQueryEntity, yloWithKsecDetails)
		mapYLOWithPLOID, isPLOCheck, ploRemark := isIdInYloWithPloId(ploDetail.ID, yloWithPloDetails)
		child, err := u.recursivePloCheckDetail(PloFormatID, ploDetail.ID, yearAndSemesterID, programSubPlanID)
		if err != nil {
			return nil, err
		}

		result = append(result, dto.ProgramYLOPLODto{
			ID:        ploDetail.ID,
			Order:     ploDetail.Order,
			ParentID:  ploDetail.ParentID,
			PLODetail: ploDetail.PLODetail,
			PLOPrefix: ploDetail.PLOPrefix,
			Children:  child,
			Ksec:      programYLOWithKsec,
			IsChecked: &isPLOCheck,
			Remark:    ploRemark,
			MapYLOID:  mapYLOWithPLOID,
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

func filterKsecOnlyInWithPLO(ksecDetails []query.ProgramKsecDetailQueryEntity, MapPloWithKsecIds []query.ProgramMapPloWithKsecQueryEntity, yloWithKsecDetails []query.ProgramYloWithKsecQueryEntity) []dto.YLOKsecDetail {
	ksecDetailList := make([]dto.YLOKsecDetail, 0)
	for _, ksecDetail := range ksecDetails {
		isChecked, ProgramMapPloWithKsecID := isIdInMapPloWithKsecId(MapPloWithKsecIds, *ksecDetail.ID)
		if *isChecked {
			mapYLOWithKsecID, isKsecCheck, ksecRemark := isIdInYloWithKsecId(ProgramMapPloWithKsecID, yloWithKsecDetails)
			ksecDetailList = append(ksecDetailList, dto.YLOKsecDetail{
				MapYLOID:  mapYLOWithKsecID,
				MapPLOID:  ProgramMapPloWithKsecID,
				Type:      ksecDetail.Type,
				Order:     ksecDetail.Order,
				Detail:    ksecDetail.Detail,
				IsChecked: &isKsecCheck,
				Remark:    ksecRemark,
				CreatedAt: ksecDetail.CreatedAt,
				UpdatedAt: ksecDetail.UpdatedAt,
			})
		}

	}
	return ksecDetailList
}

func isIdInYloWithKsecId(ksecId *uint, yloWithKsecDetails []query.ProgramYloWithKsecQueryEntity) (mapYLOWithKsecID *uint, isKsecCheck bool, ksecRemark *string) {
	isKsecCheck = false
	for _, yloWithKsecData := range yloWithKsecDetails {
		if *ksecId == *yloWithKsecData.ProgramMapPloWithKsecID {
			isKsecCheck = true
			mapYLOWithKsecID = yloWithKsecData.ID
			ksecRemark = yloWithKsecData.Remark
			return
		}
	}
	return
}

func isIdInYloWithPloId(ploId *uint, yloWithPloDetails []query.ProgramYloWithPloQueryEntity) (mapYLOWithPloID *uint, isPloCheck bool, ploRemark *string) {
	isPloCheck = false
	for _, yloWithPloData := range yloWithPloDetails {
		if *ploId == *yloWithPloData.ProgramPloID {
			isPloCheck = true
			mapYLOWithPloID = yloWithPloData.ID
			ploRemark = yloWithPloData.Remark
			return
		}
	}
	return
}
