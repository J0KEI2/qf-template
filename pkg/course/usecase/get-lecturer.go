package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

func (u *courseUsecase) GetCourseLecturer(courseUidString string) (res *dto.CourseLecturerResponseDto, err error) {
	courseUid, err := uuid.Parse(courseUidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.CourseMainQueryEntity{
		CourseID: &courseUid,
	}

	dest := query.CourseMainQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil, "Lecturer")

	mapCourseData, err := u.courseRepository.GetMapCourseLecturer(dest.Lecturer.ID)
	if err != nil {
		return nil, err
	}

	res = &dto.CourseLecturerResponseDto{
		ID:                 &dest.Lecturer.ID,
		CategoryName:       &dest.Lecturer.CategoryName,
		CourseOwnerID:      &dest.Lecturer.CourseOwnerID,
		CourseMapLecturers: mapCourseData,
		CreatedAt:          &dest.Lecturer.CreatedAt,
		UpdatedAt:          &dest.Lecturer.UpdatedAt,
	}

	return
}
