package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

func (u *courseUsecase) GetCourseInfo(courseUidString string) (res *dto.CourseInfoResponseDto, err error) {
	courseUid, err := uuid.Parse(courseUidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.CourseMainQueryEntity{
		CourseID: &courseUid,
	}

	dest := query.CourseMainQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil, "CourseInfo")

	res = &dto.CourseInfoResponseDto{
		ID:                  dest.CourseInfo.ID,
		CategoryName:        dest.CourseInfo.CategoryName,
		CourseCode:          dest.CourseInfo.CourseCode,
		CourseNameTH:        dest.CourseInfo.CourseNameEN,
		CourseNameEN:        dest.CourseInfo.CourseNameEN,
		TotalCredit:         dest.CourseInfo.TotalCredit,
		Credit1:             dest.CourseInfo.Credit1,
		Credit2:             dest.CourseInfo.Credit2,
		Credit3:             dest.CourseInfo.Credit3,
		CourseTypeID:        dest.CourseInfo.CourseTypeID,
		CourseConditionTH:   dest.CourseInfo.CourseConditionTH,
		CourseConditionEN:   dest.CourseInfo.CourseConditionEN,
		CourseDescriptionTH: dest.CourseInfo.CourseDescriptionTH,
		CourseDescriptionEN: dest.CourseInfo.CourseDescriptionEN,
		CourseObjective:     dest.CourseInfo.CourseObjective,
		Location:            dest.CourseInfo.Location,
		CreatedAt:           dest.CourseInfo.CreatedAt,
		UpdatedAt:           dest.CourseInfo.UpdatedAt,
	}

	return
}
