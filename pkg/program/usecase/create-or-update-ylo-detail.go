package usecase

import (
	"log"

	"github.com/AlekSi/pointer"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateYLODetail(yloDetail dto.CreateOrUpdateYLODetailRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateYLODetailTransaction(yloDetail))
}

func (u programUsecase) CreateOrUpdateYLODetailTransaction(yloDetail dto.CreateOrUpdateYLODetailRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, yloData := range yloDetail.YLODetails {
			yloQuery := query.ProgramYloKsecQueryEntity{
				ID:                       yloData.ID,
				ProgramSubPlanID:         &yloDetail.ProgramSubPlanID,
				ProgramYearAndSemesterID: yloData.ProgramYearAndSemesterID,
			}

			yloUpdate := query.ProgramYloKsecQueryEntity{
				ID:                       yloData.ID,
				ProgramSubPlanID:         &yloDetail.ProgramSubPlanID,
				ProgramYearAndSemesterID: yloData.ProgramYearAndSemesterID,
				Knowledge:                yloData.Knowledge,
				Skill:                    yloData.Skill,
				Ethic:                    yloData.Ethic,
				Character:                yloData.Character,
			}

			if yloQuery.ID != nil {
				if err = u.CommonRepository.Update(tx, yloQuery, &yloUpdate); err != nil {
					log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
					return err
				}
			} else {
				if err = u.CommonRepository.Create(tx, &yloUpdate); err != nil {
					log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
					return err
				}
			}

			if err = u.ProgramRepository.CreateOrUpdateYLOWithPLO(tx, yloData.YLOData.PLODetails, *yloData.ProgramYearAndSemesterID); err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}

			if err = u.ProgramRepository.CreateOrUpdateYLOWithKsec(tx, yloData.YLOData.PLODetails, *yloData.ProgramYearAndSemesterID); err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}

		return nil
	}
}

func (u programUsecase) CreateInitYLODetail(programSubPlanID *uint) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateInitYLODetailTransaction(programSubPlanID))
}

func (u programUsecase) CreateInitYLODetailTransaction(programSubPlanID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		queryDb := query.ProgramYearAndSemesterQueryEntity{
			ProgramSubPlanID: programSubPlanID,
		}

		mapCourseYear := []query.ProgramYearAndSemesterQueryEntity{}

		err = u.CommonRepository.GetList(queryDb, &mapCourseYear, nil, "CourseDetail", "CourseDetail.CourseMain")

		yearList := map[string]uint{}
		yearStrList := []string{}
		for index, yearData := range mapCourseYear {
			if index == 0 {
				yearList[*yearData.Year] = *yearData.ID
				yearStrList = append(yearStrList, *yearData.Year)
				continue
			} else {
				if _, exist := yearList[*yearData.Year]; !exist {
					yearList[*yearData.Year] = *yearData.ID
					yearStrList = append(yearStrList, *yearData.Year)
				}
			}
		}

		programPLO := query.ProgramPloFormatQueryEntity{
			ProgramSubPlanID: programSubPlanID,
		}

		if err = u.CommonRepository.GetFirst(&programPLO); err != nil {
			return
		}

		initYLODetailDto := new(dto.CreateOrUpdateYLODetailRequestDto)
		programYLODetailList := []dto.ProgramYLODetailDto{}
		for _, yearItem := range yearStrList {
			yearID := yearList[yearItem]

			ploDetails, err := u.recursivePloCheckDetail(programPLO.ID, pointer.ToUint(0), &yearID, programSubPlanID)
			if err != nil {
				return err
			}

			newYloDetail := dto.ProgramYLODetailDto{
				ID:                       nil,
				ProgramYearAndSemesterID: &yearID,
				Year:                     &yearItem,
				Knowledge:                nil,
				Skill:                    nil,
				Ethic:                    nil,
				Character:                nil,
				YLOData: dto.ProgramYLODataDto{
					ID:         programPLO.ID,
					PLOFormat:  programPLO.PLOFormat,
					PLODetails: ploDetails,
				},
			}
			programYLODetailList = append(programYLODetailList, newYloDetail)
		}

		initYLODetailDto.ProgramSubPlanID = *programSubPlanID
		initYLODetailDto.YLODetails = programYLODetailList

		err = helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateYLODetailTransaction(*initYLODetailDto))
		if err != nil {
			return err
		}
		return nil
	}
}
