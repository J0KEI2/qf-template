package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/qf4"
)

func (u *qf4Usecase) GetQF4Lecturer(courseUidString string) (res *dto.QF4LecturerResponseDto, err error) {
	courseUid, err := uuid.Parse(courseUidString)

	if err != nil {
		return nil, err
	}

	queryTb := query.QF4MainQueryEntity{
		QF4ID: &courseUid,
	}

	dest := query.QF4MainQueryEntity{}

	u.CommonRepository.GetList(&queryTb, &dest, nil, "Lecturer")

	mapQF4Data, err := u.qf4Repo.GetMapQF4Lecturer(dest.Lecturer.ID)
	if err != nil {
		return nil, err
	}

	res = &dto.QF4LecturerResponseDto{
		ID:                &dest.Lecturer.ID,
		CategoryName:      &dest.Lecturer.CategoryName,
		CourseOwnerID:     &dest.Lecturer.CourseOwnerID,
		CourseMapLecturers: mapQF4Data,
		CreatedAt:         &dest.Lecturer.CreatedAt,
		UpdatedAt:         &dest.Lecturer.UpdatedAt,
	}

	return
}
