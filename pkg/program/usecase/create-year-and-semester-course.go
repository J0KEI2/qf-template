package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateYearAndSemesterCourse(request dto.ProgramYearAndSemesterCourseRequestDto, yearAndSemesterId uint) (result *dto.ProgramYearAndSemesterResponseDto, err error) {
	courseStatement := query.ProgramCourseDetailQueryEntity{
		ID: request.CourseID,
	}

	update := query.ProgramCourseDetailQueryEntity{
		YearAndSemesterID: &yearAndSemesterId,
	}

	helper.ExecuteTransaction(u.CommonRepository, u.UpdateCourseYearAndSemesterTransaction(&courseStatement, &update))

	if err != nil {
		return nil, err
	}

	yearAndSemester := query.ProgramYearAndSemesterQueryEntity{
		ID: &yearAndSemesterId,
	}

	u.CommonRepository.GetFirst(&yearAndSemester)

	course := []dto.ProgramCourseDetailResponseDto{
		{
			ID:                update.ID,
			YearAndSemesterID: update.YearAndSemesterID,
			CourseTypeID:      update.CourseTypeID,
			CourseType:        update.CourseType,
			CourseYear:        update.CourseYear,
			CourseCode:        update.CourseCode,
			CourseNameTH:      update.CourseNameTH,
			CourseNameEN:      update.CourseNameEN,
			CourseCredit:      update.CourseCredit,
			Credit1:           update.Credit1,
			Credit2:           update.Credit2,
			Credit3:           update.Credit3,
			IsCreditCalc:      update.IsCreditCalc,
			CreatedAt:         update.CreatedAt,
			UpdatedAt:         update.UpdatedAt,
			DeletedAt:         update.DeletedAt,
		},
	}

	result = &dto.ProgramYearAndSemesterResponseDto{
		ID:            &yearAndSemesterId,
		SubPlanID:     yearAndSemester.ProgramSubPlanID,
		Year:          yearAndSemester.Year,
		Semester:      yearAndSemester.Semester,
		CourseDetails: course,
		CreatedAt:     yearAndSemester.CreatedAt,
		UpdatedAt:     yearAndSemester.UpdatedAt,
		DeletedAt:     yearAndSemester.DeletedAt,
	}

	return result, err
}

func (u programUsecase) UpdateCourseYearAndSemesterTransaction(query *query.ProgramCourseDetailQueryEntity, update *query.ProgramCourseDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, query, update)
	}
}
