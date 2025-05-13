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

func (u programUsecase) GetYearAndCourse(subPlanID uint, paginationOptions *models.PaginationOptions) (result *dto.GetYearCourseResponseDto, err error) {
	paginationOptions.DefaultLimit(0)
	
	queryDb := query.ProgramYearAndSemesterQueryEntity{
		ProgramSubPlanID: &subPlanID,
	}

	mapCourseYear := []query.ProgramYearAndSemesterQueryEntity{}

	err = u.CommonRepository.GetList(queryDb, &mapCourseYear, paginationOptions, "CourseDetail")

	
	yearsMap := map[string][]dto.ProgramSemesterResponseDto{}
	CascualCourseCodeStyleRegex, _ := regexp.Compile(`^([A-Za-z]{2})(\d{3})(\d{3})$`)
	FormalCourseCodeStyleRegex, _ := regexp.Compile(`^[A-Z]{2} \d{3} \d{3}$`)
	for _, courseYear := range mapCourseYear {
		courses := make([]dto.ProgramCourseDetailResponseDto, 0)
		for _, course := range courseYear.CourseDetail {
			u.UpdateCourseFromCoursetoRegKKU(course)
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
		year := pointer.GetString(courseYear.Year)
		yearsMap[year] = append(yearsMap[year], dto.ProgramSemesterResponseDto{
			ID:            courseYear.ID,
			SubPlanID:     courseYear.ProgramSubPlanID,
			Semester:      courseYear.Semester,
			CourseDetails: courses,
			CreatedAt:     courseYear.CreatedAt,
			UpdatedAt:     courseYear.UpdatedAt,
			DeletedAt:     courseYear.DeletedAt,
		})
	}

	years := make([]dto.ProgramYearResponseDto, 0)

	for year, semester := range yearsMap {
		sort.SliceStable(semester, func(i, j int) bool {
			return *semester[i].Semester < *semester[j].Semester
		})
		years = append(years, dto.ProgramYearResponseDto{
			Year:      pointer.ToString(year),
			Semesters: semester,
		})
	}

	sort.SliceStable(years, func(i, j int) bool {
		return *years[i].Year < *years[j].Year
	})

	result = &dto.GetYearCourseResponseDto{
		Items:             years,
		PaginationOptions: paginationOptions,
	}
	return
}
