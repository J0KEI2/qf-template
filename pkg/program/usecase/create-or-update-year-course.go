package usecase

import (
	"fmt"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateYearCourse(years []dto.ProgramYearRequestDto, subPlanID int) (err error) {
	planDetailIDUint := uint(subPlanID)
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateYearCourseTransaction(years, &planDetailIDUint))
}

func (u programUsecase) CreateOrUpdateYearCourseTransaction(years []dto.ProgramYearRequestDto, subPlanID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, year := range years {
			for _, semester := range year.Semesters {
				queryYear := query.ProgramYearAndSemesterQueryEntity{
					ID: semester.ID,
				}
				update := query.ProgramYearAndSemesterQueryEntity{
					ProgramSubPlanID: subPlanID,
					Year:             year.Year,
					Semester:         semester.Semester,
				}
				if err = u.CommonRepository.Update(tx, queryYear, &update); err != nil {
					if semester.ID == nil || err != gorm.ErrRecordNotFound {
						err = u.CommonRepository.Create(tx, &update)
					}
					if err != nil {
						return err
					}
				}

				for _, course := range semester.CourseDetails {
					queryCourse := query.ProgramCourseDetailQueryEntity{
						ID: course.ID,
					}
					updateCourse := query.ProgramCourseDetailQueryEntity{
						YearAndSemesterID: semester.ID,
					}
					if err = u.CommonRepository.Update(tx, queryCourse, &updateCourse); err != nil {
						return err
					}
				}
			}
		}
		return
	}
}

func (u programUsecase) CreateYearAndSemesterByEducationYear(subPlanId, academicYear uint) error {
	return helper.ExecuteTransaction(u.CommonRepository, u.createYearAndSemesterByEducationYearTransaction(subPlanId, academicYear))
}

func (u programUsecase) createYearAndSemesterByEducationYearTransaction(subPlanId, years uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for year := uint(1); year <= years; year++ {
			for semester := 1; semester <= 2; semester++ {
				yearAndSemester := query.ProgramYearAndSemesterQueryEntity{
					ProgramSubPlanID: &subPlanId,
					Year:             pointer.ToString(fmt.Sprintf("%d", year)),
					Semester:         pointer.ToString(fmt.Sprintf("%d", semester)),
				}
				if err := u.CommonRepository.Create(tx, &yearAndSemester); err != nil {
					return err
				}
			}
		}
		return
	}
}
