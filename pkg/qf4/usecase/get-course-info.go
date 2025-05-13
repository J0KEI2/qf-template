package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"
)

func (u *qf4Usecase) GetQF4CourseInfo(courseUidString string) (res *dto.QF4CourseInfoResponseDto, err error) {
	courseUid, err := uuid.Parse(courseUidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.QF4MainQueryEntity{
		QF4ID: &courseUid,
	}

	dest := query.QF4MainQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil, "CourseInfo")

	res = &dto.QF4CourseInfoResponseDto{
		ID:                  dest.CourseInfo.ID,
		CategoryName:        dest.CourseInfo.CategoryName,
		CourseCode:          dest.CourseInfo.CourseCode,
		CourseNameTH:        dest.CourseInfo.CourseNameEN,
		CourseNameEN:        dest.CourseInfo.CourseNameEN,
		NumberOfCredits:     dest.CourseInfo.NumberOfCredits,
		CourseTypeID:        dest.CourseInfo.CourseTypeID,
		CourseConditionTH:   dest.CourseInfo.CourseConditionTH,
		CourseConditionEN:   dest.CourseInfo.CourseConditionEN,
		CourseDescriptionTH: dest.CourseInfo.CourseDescriptionTH,
		CourseDescriptionEN: dest.CourseInfo.CourseDescriptionEN,
		CourseObjective:     dest.CourseInfo.CourseObjective,
		StudentActivity:     dest.CourseInfo.StudentActivity,
		FacilitatorTask:     dest.CourseInfo.FacilitatorTask,
		ConsultantTask:      dest.CourseInfo.ConsultantTask,
		StudentGuideline:    dest.CourseInfo.StudentGuideline,
		Location:            dest.CourseInfo.Location,
		StudentSupport:      dest.CourseInfo.StudentSupport,
	}

	return
}
