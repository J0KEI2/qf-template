package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"

)

func (u *qf4Usecase) GetQF4CoursePlan(qf4UidString string) (res []dto.QF4CoursePlanResponseDto, err error) {
	qf4Uid, err := uuid.Parse(qf4UidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.QF4CoursePlanQueryEntity{
		Qf4MainUid: &qf4Uid,
	}

	dest := []query.QF4CoursePlanQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil)

	res = make([]dto.QF4CoursePlanResponseDto, len(dest))

	for i, coursePlan := range dest {
		res[i] = dto.QF4CoursePlanResponseDto{
			ID:               coursePlan.ID,
			Qf4MainUid:       coursePlan.Qf4MainUid,
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
