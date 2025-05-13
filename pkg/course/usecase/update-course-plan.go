package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"

	"gorm.io/gorm"
)

func (u *courseUsecase) UpdateCoursePlan(data *dto.CourseUpdateCoursePlanRequestDto) (res *dto.CoursePlanResponseDto, err error) {

	queryTb := query.CoursePlanQueryEntity{
		ID: data.ID,
	}

	updateValue := query.CoursePlanQueryEntity{
		Qf3MainUid:       data.Qf3MainUid,
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

	res = &dto.CoursePlanResponseDto{
		ID:               queryTb.ID,
		Qf3MainUid:       queryTb.Qf3MainUid,
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

func (u *courseUsecase) updateCoursePlanAction(query *query.CoursePlanQueryEntity, updateValue *query.CoursePlanQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, query, updateValue)
	}
}
