package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

func (u *courseUsecase) GetCoursePlan(courseUidString string) (res []dto.CoursePlanResponseDto, err error) {
	courseUid, err := uuid.Parse(courseUidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.CoursePlanQueryEntity{
		Qf3MainUid: &courseUid,
	}

	dest := []query.CoursePlanQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil)

	res = make([]dto.CoursePlanResponseDto, len(dest))

	for i, coursePlan := range dest {
		res[i] = dto.CoursePlanResponseDto{
			ID:               coursePlan.ID,
			Qf3MainUid:       coursePlan.Qf3MainUid,
			Week:             coursePlan.Week,
			Title:            coursePlan.Title,
			PlanDescription:  coursePlan.PlanDescription,
			TheoryHour:       coursePlan.TheoryHour,
			OperationHour:    coursePlan.OperationHour,
			SelfLearningHour: coursePlan.SelfLearningHour,
			LearningOutcome:  coursePlan.LearningOutcome,
			TeachingMedia:    coursePlan.TeachingMedia,
			LeaningActivity:  coursePlan.LeaningActivity,
			EvaluationMethod: coursePlan.EvaluationMethod,
			Lecturer:         coursePlan.Lecturer,
			AssessmentScore:  coursePlan.AssessmentScore,
			AssessmentNote:   coursePlan.AssessmentNote,
			CreatedAt:        coursePlan.CreatedAt,
			UpdatedAt:        coursePlan.UpdatedAt,
		}
	}

	return
}
