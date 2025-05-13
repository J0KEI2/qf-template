package usecase

import (
	"encoding/json"
	"sort"

	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetCurriculumMapping(paginationOptions models.PaginationOptions, subPlanID uint) (result *dto.ProgramCurMapRespGetResponseDto, err error) {
	curMapQuery := query.ProgramMapCurMapRespQueryEntity{
		ProgramSubPlanID: &subPlanID,
	}

	curMapList := []query.ProgramMapCurMapRespQueryEntity{}
	err = u.CommonRepository.GetList(&curMapQuery, &curMapList, nil, "ProgramCourseDetail", "ProgramPlo", "ProgramPlo.Children")
	if err != nil {
		return nil, err
	}

	if len(curMapList) == 0 {
		err = helper.ExecuteTransaction(u.CommonRepository, u.InitCurriculumMapping(paginationOptions, subPlanID))
		if err != nil {
			return nil, err
		}
		err = u.CommonRepository.GetList(&curMapQuery, &curMapList, nil, "ProgramCourseDetail", "ProgramPlo", "ProgramPlo.Children")
		if err != nil {
			return nil, err
		}
	}

	fundamentalListResp := []dto.ProgramCurMapDetail{}
	compulsoryListResp := []dto.ProgramCurMapDetail{}
	enrichmentListResp := []dto.ProgramCurMapDetail{}

	for _, curMapItem := range curMapList {
		bodyData := dto.ProgramCurMapDetail{
			ID:            curMapItem.ID,
			ProgramCourseID:   curMapItem.ProgramCourseDetailID,
			ProgramCourseType: curMapItem.ProgramCourseDetail.CourseTypeID,
			CourseNameTH:  curMapItem.ProgramCourseDetail.CourseNameTH,
			CourseNameEN:  curMapItem.ProgramCourseDetail.CourseNameEN,
			PLOID:         curMapItem.ProgramPlo.ID,
			ProgramPloPrefix:  curMapItem.ProgramPlo.PLOPrefix,
			Status:        curMapItem.Status,
		}
		if *curMapItem.ProgramCourseDetail.CourseTypeID == 1 {
			nameEN := "Fundamental module"
			nameTH := "วิชาศึกษาทั้วไป"
			bodyData.ProgramHeaderNameEN = &nameEN
			bodyData.ProgramHeaderNameTH = &nameTH

			fundamentalListResp = append(fundamentalListResp, bodyData)
		} else if *curMapItem.ProgramCourseDetail.CourseTypeID == 2 {
			nameEN := "Compulsory Modules"
			nameTH := "วิชาในหลักสูตร"
			bodyData.ProgramHeaderNameEN = &nameEN
			bodyData.ProgramHeaderNameTH = &nameTH

			compulsoryListResp = append(compulsoryListResp, bodyData)
		} else {
			nameEN := "Enrichment Modules"
			nameTH := "วิชาเลือก"
			bodyData.ProgramHeaderNameEN = &nameEN
			bodyData.ProgramHeaderNameTH = &nameTH

			enrichmentListResp = append(enrichmentListResp, bodyData)
		}
	}

	fundamentalResp := u.mappingResponse(fundamentalListResp)
	compulsoryResp := u.mappingResponse(compulsoryListResp)
	enrichmentResp := u.mappingResponse(enrichmentListResp)

	result = &dto.ProgramCurMapRespGetResponseDto{
		ProgramSubPlanID:          subPlanID,
		ProgramFundamentalDetails: fundamentalResp,
		ProgramCompulsoryDetails:  compulsoryResp,
		ProgramEnrichmentDetails:  enrichmentResp,
	}

	return
}

func (u programUsecase) GetCurriculumMappingKsa(paginationOptions models.PaginationOptions, subPlanID uint) (result *dto.ProgramCurMapKsaGetResponseDto, err error) {

	curMapKsaQuery := query.ProgramMapCurMapKsaQueryEntity{
		ProgramSubPlanID: &subPlanID,
	}

	curMapKsaList := []query.ProgramMapCurMapKsaQueryEntity{}
	err = u.CommonRepository.GetList(&curMapKsaQuery, &curMapKsaList, nil, "ProgramCourseDetail", "ProgramPlo", "ProgramPlo.Children")
	if err != nil {
		return nil, err
	}

	if len(curMapKsaList) == 0 {
		err = helper.ExecuteTransaction(u.CommonRepository, u.InitCurriculumMappingKsa(paginationOptions, subPlanID))
		if err != nil {
			return nil, err
		}
		err = u.CommonRepository.GetList(&curMapKsaQuery, &curMapKsaList, nil, "ProgramCourseDetail", "ProgramPlo", "ProgramPlo.Children")
		if err != nil {
			return nil, err
		}
	}

	fundamentalListResp := []dto.ProgramCurMapKsaDetail{}
	compulsoryListResp := []dto.ProgramCurMapKsaDetail{}
	enrichmentListResp := []dto.ProgramCurMapKsaDetail{}

	for _, curMapItem := range curMapKsaList {
		ksaID := []int{}
		if curMapItem.KsaID != nil {
			json.Unmarshal([]byte(*curMapItem.KsaID), &ksaID)
		}

		bodyData := dto.ProgramCurMapKsaDetail{
			ID:            curMapItem.ID,
			ProgramCourseID:   curMapItem.ProgramCourseDetailID,
			ProgramCourseType: curMapItem.ProgramCourseDetail.CourseTypeID,
			CourseNameTH:  curMapItem.ProgramCourseDetail.CourseNameTH,
			CourseNameEN:  curMapItem.ProgramCourseDetail.CourseNameEN,
			PLOID:         curMapItem.ProgramPlo.ID,
			ProgramPloPrefix:  curMapItem.ProgramPlo.PLOPrefix,
			KsaID:         ksaID,
		}
		if *curMapItem.ProgramCourseDetail.CourseTypeID == 1 {
			nameEN := "Fundamental module"
			nameTH := "วิชาศึกษาทั้วไป"
			bodyData.ProgramHeaderNameEN = &nameEN
			bodyData.ProgramHeaderNameTH = &nameTH

			fundamentalListResp = append(fundamentalListResp, bodyData)
		} else if *curMapItem.ProgramCourseDetail.CourseTypeID == 2 {
			nameEN := "Compulsory Modules"
			nameTH := "วิชาในหลักสูตร"
			bodyData.ProgramHeaderNameEN = &nameEN
			bodyData.ProgramHeaderNameTH = &nameTH

			compulsoryListResp = append(compulsoryListResp, bodyData)
		} else {
			nameEN := "Enrichment Modules"
			nameTH := "วิชาเลือก"
			bodyData.ProgramHeaderNameEN = &nameEN
			bodyData.ProgramHeaderNameTH = &nameTH

			enrichmentListResp = append(enrichmentListResp, bodyData)
		}
	}

	fundamentalResp := u.mappingKsaResponse(fundamentalListResp)
	compulsoryResp := u.mappingKsaResponse(compulsoryListResp)
	enrichmentResp := u.mappingKsaResponse(enrichmentListResp)

	result = &dto.ProgramCurMapKsaGetResponseDto{
		ProgramSubPlanID:          subPlanID,
		ProgramFundamentalDetails: fundamentalResp,
		ProgramCompulsoryDetails:  compulsoryResp,
		ProgramEnrichmentDetails:  enrichmentResp,
	}

	return
}

func (u programUsecase) SeparateCourseType(courseList []query.ProgramCourseDetailQueryEntity, ploDetails []dto.ProgramPLODetailDto, subPlanID uint) ([]query.ProgramMapCurMapRespQueryEntity, []query.ProgramMapCurMapRespQueryEntity, []query.ProgramMapCurMapRespQueryEntity, error) {
	initStatus := 1

	fundamentalList := []query.ProgramMapCurMapRespQueryEntity{}
	compulsoryList := []query.ProgramMapCurMapRespQueryEntity{}
	enrichmentList := []query.ProgramMapCurMapRespQueryEntity{}

	for _, courseDetail := range courseList {
		if *courseDetail.CourseTypeID == 1 {
			for _, ploDetail := range ploDetails {
				if len(ploDetail.Children) > 0 {
					for _, child := range ploDetail.Children {
						fundamentalList = append(fundamentalList, query.ProgramMapCurMapRespQueryEntity{
							ProgramSubPlanID:      &subPlanID,
							ProgramCourseDetailID: courseDetail.ID,
							ProgramPloID:          child.ID,
							Status:            &initStatus,
						})
					}
				}
			}
		} else if *courseDetail.CourseTypeID == 2 {
			for _, ploDetail := range ploDetails {
				if len(ploDetail.Children) > 0 {
					for _, child := range ploDetail.Children {
						compulsoryList = append(compulsoryList, query.ProgramMapCurMapRespQueryEntity{
							ProgramSubPlanID:      &subPlanID,
							ProgramCourseDetailID: courseDetail.ID,
							ProgramPloID:          child.ID,
							Status:            &initStatus,
						})
					}
				}
			}
		} else {
			for _, ploDetail := range ploDetails {
				if len(ploDetail.Children) > 0 {
					for _, child := range ploDetail.Children {
						enrichmentList = append(enrichmentList, query.ProgramMapCurMapRespQueryEntity{
							ProgramSubPlanID:      &subPlanID,
							ProgramCourseDetailID: courseDetail.ID,
							ProgramPloID:          child.ID,
							Status:            &initStatus,
						})
					}
				}
			}
		}
	}

	return fundamentalList, compulsoryList, enrichmentList, nil
}

func (u programUsecase) SeparateCourseTypeKsa(courseList []query.ProgramCourseDetailQueryEntity, ploDetails []dto.ProgramPLODetailDto, subPlanID uint) ([]query.ProgramMapCurMapKsaQueryEntity, []query.ProgramMapCurMapKsaQueryEntity, []query.ProgramMapCurMapKsaQueryEntity, error) {
	fundamentalList := []query.ProgramMapCurMapKsaQueryEntity{}
	compulsoryList := []query.ProgramMapCurMapKsaQueryEntity{}
	enrichmentList := []query.ProgramMapCurMapKsaQueryEntity{}

	for _, courseDetail := range courseList {
		if *courseDetail.CourseTypeID == 1 {
			for _, ploDetail := range ploDetails {
				if len(ploDetail.Children) > 0 {
					for _, child := range ploDetail.Children {
						fundamentalList = append(fundamentalList, query.ProgramMapCurMapKsaQueryEntity{
							ProgramSubPlanID:      &subPlanID,
							ProgramCourseDetailID: courseDetail.ID,
							ProgramPloID:          child.ID,
							KsaID:             nil,
						})
					}
				}
			}
		} else if *courseDetail.CourseTypeID == 2 {
			for _, ploDetail := range ploDetails {
				if len(ploDetail.Children) > 0 {
					for _, child := range ploDetail.Children {
						compulsoryList = append(compulsoryList, query.ProgramMapCurMapKsaQueryEntity{
							ProgramSubPlanID:      &subPlanID,
							ProgramCourseDetailID: courseDetail.ID,
							ProgramPloID:          child.ID,
							KsaID:             nil,
						})
					}
				}
			}
		} else {
			for _, ploDetail := range ploDetails {
				if len(ploDetail.Children) > 0 {
					for _, child := range ploDetail.Children {
						enrichmentList = append(enrichmentList, query.ProgramMapCurMapKsaQueryEntity{
							ProgramSubPlanID:      &subPlanID,
							ProgramCourseDetailID: courseDetail.ID,
							ProgramPloID:          child.ID,
							KsaID:             nil,
						})
					}
				}
			}
		}
	}

	return fundamentalList, compulsoryList, enrichmentList, nil
}

func (u programUsecase) mappingResponse(listData []dto.ProgramCurMapDetail) dto.ProgramCurMapResp {
	courseMap := make(map[string]*dto.ProgramCurMapRespResponse)
	for _, input := range listData {
		if course, exists := courseMap[*input.CourseNameEN]; exists {
			// If course exists, append PloDetail
			course.ProgramCurMapRespDetails = append(course.ProgramCurMapRespDetails, dto.ProgramCurMapRespDetails{
				CurMapID:     input.ID,
				PLOID:        input.PLOID,
				ProgramPloPrefix: input.ProgramPloPrefix,
				Status:       input.Status,
			})
		} else {
			// If course does not exist, create new Output and add to map
			courseMap[*input.CourseNameEN] = &dto.ProgramCurMapRespResponse{
				ID:           *input.ID,
				ProgramCourseID:  input.ProgramCourseID,
				CourseNameTH: input.CourseNameTH,
				CourseNameEN: input.CourseNameEN,
				ProgramCurMapRespDetails: []dto.ProgramCurMapRespDetails{
					{
						CurMapID:     input.ID,
						PLOID:        input.PLOID,
						ProgramPloPrefix: input.ProgramPloPrefix,
						Status:       input.Status,
					},
				},
			}
		}
	}

	var outputs []dto.ProgramCurMapRespResponse
	for _, course := range courseMap {
		outputs = append(outputs, *course)
	}

	sort.Slice(outputs, func(i, j int) bool {
		return outputs[i].ID < outputs[j].ID
	})

	result := dto.ProgramCurMapResp{
		ProgramHeaderNameEN: listData[0].ProgramHeaderNameEN,
		ProgramHeaderNameTH: listData[0].ProgramHeaderNameTH,
		ProgramCourseType:   listData[0].ProgramCourseType,
		Items:           outputs,
	}

	return result
}

func (u programUsecase) mappingKsaResponse(listData []dto.ProgramCurMapKsaDetail) dto.ProgramCurMapKsa {
	courseMap := make(map[string]*dto.ProgramCurMapKsaResponse)
	for _, input := range listData {
		if course, exists := courseMap[*input.CourseNameEN]; exists {
			// If course exists, append PloDetail
			course.ProgramCurMapKsaDetails = append(course.ProgramCurMapKsaDetails, dto.ProgramCurMapKsaDetails{
				CurMapID:     input.ID,
				PLOID:        input.PLOID,
				ProgramPloPrefix: input.ProgramPloPrefix,
				KsaID:        input.KsaID,
			})
		} else {
			// If course does not exist, create new Output and add to map
			courseMap[*input.CourseNameEN] = &dto.ProgramCurMapKsaResponse{
				ID:           *input.ID,
				ProgramCourseID:  input.ProgramCourseID,
				CourseNameTH: input.CourseNameTH,
				CourseNameEN: input.CourseNameEN,
				ProgramCurMapKsaDetails: []dto.ProgramCurMapKsaDetails{
					{
						CurMapID:     input.ID,
						PLOID:        input.PLOID,
						ProgramPloPrefix: input.ProgramPloPrefix,
						KsaID:        input.KsaID,
					},
				},
			}
		}
	}

	var outputs []dto.ProgramCurMapKsaResponse
	for _, course := range courseMap {
		outputs = append(outputs, *course)
	}

	sort.Slice(outputs, func(i, j int) bool {
		return outputs[i].ID < outputs[j].ID
	})

	result := dto.ProgramCurMapKsa{
		ProgramHeaderNameEN: listData[0].ProgramHeaderNameEN,
		ProgramHeaderNameTH: listData[0].ProgramHeaderNameTH,
		ProgramCourseType:   listData[0].ProgramCourseType,
		Items:           outputs,
	}

	return result
}
