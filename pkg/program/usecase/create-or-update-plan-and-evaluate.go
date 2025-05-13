package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdatePlanAndEvaluate(planAndEvaluate dto.ProgramPlanAndEvaluateRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdatePlanAndEvaluateTransaction(planAndEvaluate))
}

func (u programUsecase) CreateOrUpdatePlanAndEvaluateTransaction(planAndEvaluate dto.ProgramPlanAndEvaluateRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		var studentCharacteristic *string
		if planAndEvaluate.StudentCharacteristics != nil {
			jsonByte, _ := json.Marshal(planAndEvaluate.StudentCharacteristics)
			studentCharacteristic = pointer.ToStringOrNil(string(jsonByte))
		} else {
			initStudentCharacteristicList := []map[string]interface{}{}
			jsonByte, _ := json.Marshal(initStudentCharacteristicList)
			studentCharacteristic = pointer.ToStringOrNil(string(jsonByte))
		}

		var receiveStudentPlan *string
		if planAndEvaluate.ReceiveStudentPlan != nil {
			jsonByte, _ := json.Marshal(planAndEvaluate.ReceiveStudentPlan)
			receiveStudentPlan = pointer.ToStringOrNil(string(jsonByte))
		} else {
			initReceiveStudentPlanList := []map[string]interface{}{}
			result, err := u.GetYearAndCourse(*planAndEvaluate.ProgramSubPlanID, nil)
			if err != nil {
				return err
			}
			numYear := len(result.Items)
			for i := 1; i <= numYear+1; i++ {
				lastIndex := numYear + 1
				if i == lastIndex {
					initReceiveStudentPlanLastItem := map[string]interface{}{
						"name":     "คาดว่าจะสำเร็จการศึกษา", // hardcode
						"position": "footer",
						"data":     []int{0, 0, 0, 0, 0},
					}
					initReceiveStudentPlanList = append(initReceiveStudentPlanList, initReceiveStudentPlanLastItem)
				} else {
					initReceiveStudentPlanItem := map[string]interface{}{
						"name":     fmt.Sprintf("ชั้นปีที่ %v", i),
						"position": "body",
						"data":     []int{0, 0, 0, 0, 0},
					}
					initReceiveStudentPlanList = append(initReceiveStudentPlanList, initReceiveStudentPlanItem)
				}
			}
			jsonByte, _ := json.Marshal(initReceiveStudentPlanList)
			receiveStudentPlan = pointer.ToStringOrNil(string(jsonByte))
		}

		var programIncome *string
		if planAndEvaluate.ProgramIncome != nil {
			jsonByte, _ := json.Marshal(planAndEvaluate.ProgramIncome)
			programIncome = pointer.ToStringOrNil(string(jsonByte))
		}

		var programOutcome *string
		if planAndEvaluate.ProgramOutcome != nil {
			jsonByte, _ := json.Marshal(planAndEvaluate.ProgramOutcome)
			programOutcome = pointer.ToStringOrNil(string(jsonByte))
		}

		update := query.ProgramPlanAndEvaluateQueryEntity{
			ProgramSubPlanID:      planAndEvaluate.ProgramSubPlanID,
			StudentCharacteristic: studentCharacteristic,
			ReceiveStudentPlan:    receiveStudentPlan,
			ProgramIncome:         programIncome,
			ProgramOutcome:        programOutcome,
			AcademicEvaluation:    planAndEvaluate.AcademicEvaluation,
			GraduationCriteria:    planAndEvaluate.GraduationCriteria,
		}

		queryTb := query.ProgramPlanAndEvaluateQueryEntity{
			ProgramSubPlanID: planAndEvaluate.ProgramSubPlanID,
		}

		err = u.CommonRepository.GetFirst(&queryTb)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = u.CommonRepository.Create(tx, &update)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		err = u.CommonRepository.Update(tx, queryTb, &update)
		if err != nil {
			return err
		}

		return
	}
}

func (u programUsecase) InitPlanAndEvaluate(subPlanID uint, numberOfYear int) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		initReceiveStudentPlanList := []map[string]interface{}{}
		var receiveStudentPlan *string
		for i := 0; i < numberOfYear+1; i++ {
			if i == numberOfYear {
				initReceiveStudentPlanLastItem := map[string]interface{}{
					"name":     "คาดว่าจะสำเร็จการศึกษา", // hardcode
					"position": "footer",
					"data":     []int{0, 0, 0, 0, 0},
				}
				initReceiveStudentPlanList = append(initReceiveStudentPlanList, initReceiveStudentPlanLastItem)
			} else {
				initReceiveStudentPlanItem := map[string]interface{}{
					"name":     fmt.Sprintf("ชั้นปีที่ %v", i+1),
					"position": nil,
					"data":     []int{0, 0, 0, 0, 0},
				}
				initReceiveStudentPlanList = append(initReceiveStudentPlanList, initReceiveStudentPlanItem)
			}
			jsonByte, _ := json.Marshal(initReceiveStudentPlanList)
			receiveStudentPlan = pointer.ToStringOrNil(string(jsonByte))
		}

		nilString := "[]"
		update := query.ProgramPlanAndEvaluateQueryEntity{
			ProgramSubPlanID:      &subPlanID,
			ReceiveStudentPlan:    receiveStudentPlan,
			StudentCharacteristic: &nilString,
			ProgramIncome:         &nilString,
			ProgramOutcome:        &nilString,
			AcademicEvaluation:    &nilString,
			GraduationCriteria:    &nilString,
		}
		err := u.CommonRepository.Create(tx, &update)
		if err != nil {
			return err
		}

		return nil
	}
}
