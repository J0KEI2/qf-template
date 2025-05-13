package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

	"gorm.io/gorm"
)

func (u *qf4Usecase) UpdateQF4CoursePlan(data *dto.QF4UpdateCoursePlanRequestDto) (res *dto.QF4CoursePlanResponseDto, err error) {

	queryTb := query.QF4CoursePlanQueryEntity{
		ID: data.ID,
	}

	updateValue := query.QF4CoursePlanQueryEntity{
		Qf4MainUid:       data.Qf4MainUid,
		Week:             data.Week,
		Title:            data.Title,
		PlanDescription:  data.PlanDescription,
		TheoryHour:       data.TheoryHour,
		OperationHour:    data.OperationHour,
		SelfLearningHour: data.SelfLearningHour,
		LearningOutcome:  data.LearningOutcome,
		TeachingMedia:    data.TeachingMedia,
		LeaningActivity:  data.LeaningActivity,
		EvaluationMethod: data.EvaluationMethod,
		Lecturer:         data.Lecturer,
		AssessmentScore:  data.AssessmentScore,
		AssessmentNote:   data.AssessmentNote,
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.updateCoursePlanAction(&queryTb, &updateValue))

	if err != nil {
		return nil, err
	}

	res = &dto.QF4CoursePlanResponseDto{
		ID:               queryTb.ID,
		Qf4MainUid:       queryTb.Qf4MainUid,
		Week:             queryTb.Week,
		Title:            queryTb.Title,
		PlanDescription:  queryTb.PlanDescription,
		TheoryHour:       queryTb.TheoryHour,
		OperationHour:    queryTb.OperationHour,
		SelfLearningHour: queryTb.SelfLearningHour,
		LearningOutcome:  queryTb.LearningOutcome,
		TeachingMedia:    queryTb.TeachingMedia,
		LeaningActivity:  queryTb.LeaningActivity,
		EvaluationMethod: queryTb.EvaluationMethod,
		Lecturer:         queryTb.Lecturer,
		AssessmentScore:  queryTb.AssessmentScore,
		AssessmentNote:   queryTb.AssessmentNote,
		CreatedAt:        queryTb.CreatedAt,
		UpdatedAt:        queryTb.UpdatedAt,
	}

	return
}

func (u *qf4Usecase) updateCoursePlanAction(query *query.QF4CoursePlanQueryEntity, updateValue *query.QF4CoursePlanQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, query, updateValue)
	}
}
