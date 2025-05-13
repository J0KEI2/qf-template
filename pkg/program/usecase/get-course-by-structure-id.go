package usecase

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetCourseByStructureId(structureId uint, paginationOptions *models.PaginationOptions) (result *dto.GetCourseByStructureResponseDto, err error) {
	structureQuery := query.ProgramStructureDetailQueryEntity{
		ID: &structureId,
	}

	err = u.CommonRepository.GetFirst(&structureQuery, "CourseDetail")

	if err != nil {
		return nil, err
	}

	courses := make([]dto.ProgramCourseDetailResponseDto, 0)

	CascualCourseCodeStyleRegex, _ := regexp.Compile(`^([A-Za-z]{2})(\d{3})(\d{3})$`)
	FormalCourseCodeStyleRegex, _ := regexp.Compile(`^[A-Z]{2} \d{3} \d{3}$`)
	for _, course := range structureQuery.CourseDetail {
		courseCode := pointer.GetString(course.CourseCode)
		if match := FormalCourseCodeStyleRegex.MatchString(courseCode); !match {
			courseCode = strings.ReplaceAll(courseCode, " ", "")
			if match := CascualCourseCodeStyleRegex.FindStringSubmatch(courseCode); match != nil {
				match[1] = strings.ToUpper(match[1])
				courseCode = strings.Join(match[1:len(match)], " ")
			}
		}
		creditDetail := fmt.Sprintf("(%d - %d - %d)", pointer.GetUint(course.Credit1), pointer.GetUint(course.Credit2), pointer.GetUint(course.Credit3))
		courses = append(courses, dto.ProgramCourseDetailResponseDto{
			ID:                course.ID,
			YearAndSemesterID: course.YearAndSemesterID,
			CourseTypeID:      course.CourseTypeID,
			CourseType:        course.CourseType,
			CourseCode:        &courseCode,
			CourseNameTH:      course.CourseNameTH,
			CourseNameEN:      course.CourseNameEN,
			CourseCredit:      course.CourseCredit,
			CreditDetail:      &creditDetail,
			Credit1:           course.Credit1,
			Credit2:           course.Credit2,
			Credit3:           course.Credit3,
			IsCreditCalc:      course.IsCreditCalc,
			CreatedAt:         course.CreatedAt,
			UpdatedAt:         course.UpdatedAt,
			DeletedAt:         course.DeletedAt,
		})
	}

	sort.SliceStable(courses, func(i, j int) bool {
		return *courses[i].ID < *courses[j].ID
	})

	result = &dto.GetCourseByStructureResponseDto{
		Items:             courses,
		PaginationOptions: paginationOptions,
	}
	return result, nil
}
