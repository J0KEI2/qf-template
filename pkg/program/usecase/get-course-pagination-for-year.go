package usecase

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u programUsecase) GetCourseDetailPagination(paginationOptions models.PaginationOptions, subPlanID int) (result *dto.GetCourseDetailPaginationResponseDto, err error) {

	courseDetails, err := u.ProgramRepository.GetCourseDetailPaginationForYear(&paginationOptions, uint(subPlanID))
	if err != nil {
		return nil, err
	}

	courseDetailResult := make([]dto.ProgramCourseDetailResponseDto, 0)

	CascualCourseCodeStyleRegex, _ := regexp.Compile(`^([A-Za-z]{2})(\d{3})(\d{3})$`)
	FormalCourseCodeStyleRegex, _ := regexp.Compile(`^[A-Z]{2} \d{3} \d{3}$`)
	for _, courseDetail := range courseDetails {
		course := courseDetail
		name := strings.Join([]string{
			pointer.GetString(course.CourseSource),
			pointer.GetString(course.CourseYear),
			pointer.GetString(course.CourseCode),
			pointer.GetString(course.CourseNameTH),
			pointer.GetString(course.CourseNameEN),
		}, " | ")
		courseCode := pointer.GetString(course.CourseCode)
		if match := FormalCourseCodeStyleRegex.MatchString(courseCode); !match {
			courseCode = strings.ReplaceAll(courseCode, " ", "")
			if match := CascualCourseCodeStyleRegex.FindStringSubmatch(courseCode); match != nil {
				match[1] = strings.ToUpper(match[1])
				courseCode = strings.Join(match[1:len(match)], " ")
			}
		}
		creditDetail := fmt.Sprintf("(%d - %d - %d)", pointer.GetUint(course.Credit1), pointer.GetUint(course.Credit2), pointer.GetUint(course.Credit3))
		courseDetailResult = append(courseDetailResult, dto.ProgramCourseDetailResponseDto{
			ID:                course.ID,
			YearAndSemesterID: course.YearAndSemesterID,
			CourseType:        course.CourseType,
			CourseCode:        &courseCode,
			CourseYear:        course.CourseYear,
			CourseNameTH:      course.CourseNameTH,
			CourseNameEN:      course.CourseNameEN,
			Name:              &name,
			CourseCredit:      course.CourseCredit,
			CreditDetail:      &creditDetail,
			Credit1:           course.Credit1,
			Credit2:           course.Credit2,
			Credit3:           course.Credit3,
			IsCreditCalc:      course.IsCreditCalc,
			IsEditedCourse:    course.IsEditedCourse,
			IsNewCourse:       course.IsNewCourse,
			CreatedAt:         course.CreatedAt,
			UpdatedAt:         course.UpdatedAt,
			DeletedAt:         course.DeletedAt,
		})
	}

	result = &dto.GetCourseDetailPaginationResponseDto{
		Items:             courseDetailResult,
		PaginationOptions: &paginationOptions,
	}

	return
}
