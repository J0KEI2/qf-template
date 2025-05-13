package usecase

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/AlekSi/pointer"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	programConstant "github.com/zercle/kku-qf-services/pkg/constant/program"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (u programUsecase) GetCoursePaginationForStructure(options models.PaginationOptions) (result *dto.GetCoursePaginationForStructureResponseDto, err error) {

	regex := `^\d{6}$|^[A-Z]{2}\d{6}$`
	pattern := regexp.MustCompile(regex)

	regCourses := []dto.RegCourseResponse{}
	if courseCode := pointer.GetString(options.Search); pattern.MatchString(courseCode) {
		regCourses, _ = u.RegUseCase.GetRegCourseByCourseCode(courseCode)
	}

	halfLimit := options.GetLimit() / 2
	if courseAmount := len(regCourses); courseAmount > 0 && halfLimit != 0 {
		if courseAmount > halfLimit {
			regCourses = regCourses[0 : halfLimit-1]
			options.SetLimit(halfLimit)
		} else {
			options.SetLimit(options.GetLimit() - courseAmount)
		}
	}

	courseCourses, err := u.ProgramRepository.GetCourseInfoPagnation(&options)
	if err != nil {
		return nil, err
	}

	courseInfoResult := make([]dto.CoursePaginationForStructure, 0)

	courseSource := programConstant.COURSE_SOURCE_COURSE
	CascualCourseCodeStyleRegex, _ := regexp.Compile(`^([A-Za-z]{2})(\d{3})(\d{3})$`)
	FormalCourseCodeStyleRegex, _ := regexp.Compile(`^[A-Z]{2} \d{3} \d{3}$`)
	for _, course := range courseCourses {
		courseInfo := course.CourseInfo
		courseName := strings.Join([]string{
			courseSource,
			"แก้ไขครั้งที่ " + pointer.GetString(course.Version),
			pointer.GetString(courseInfo.CourseCode),
			pointer.GetString(courseInfo.CourseNameTH),
			pointer.GetString(courseInfo.CourseNameEN),
		}, " | ")
		courseCode := pointer.GetString(courseInfo.CourseCode)
		if match := FormalCourseCodeStyleRegex.MatchString(courseCode); !match {
			courseCode = strings.ReplaceAll(courseCode, " ", "")
			if match := CascualCourseCodeStyleRegex.FindStringSubmatch(courseCode); match != nil {
				match[1] = strings.ToUpper(match[1])
				courseCode = strings.Join(match[1:len(match)], " ")
			}
		}
		creditDetail := fmt.Sprintf("(%d - %d - %d)", pointer.GetUint(courseInfo.Credit1), pointer.GetUint(courseInfo.Credit2), pointer.GetUint(courseInfo.Credit3))
		courseInfoResult = append(courseInfoResult, dto.CoursePaginationForStructure{
			ID:                  nil,
			CourseKey:           course.CourseID,
			Name:                &courseName,
			CategoryName:        courseInfo.CategoryName,
			CourseCode:          &courseCode,
			CourseSource:        pointer.ToString(courseSource),
			CourseYear:          course.EducationYear,
			CourseNameTH:        courseInfo.CourseNameTH,
			CourseNameEN:        courseInfo.CourseNameEN,
			Version:             course.Version,
			CourseCredit:        courseInfo.TotalCredit,
			CreditDetail:        &creditDetail,
			Credit1:             courseInfo.Credit1,
			Credit2:             courseInfo.Credit2,
			Credit3:             courseInfo.Credit3,
			CourseTypeID:        courseInfo.CourseTypeID,
			CourseConditionTH:   courseInfo.CourseConditionTH,
			CourseConditionEN:   courseInfo.CourseConditionEN,
			CourseDescriptionTH: courseInfo.CourseDescriptionTH,
			CourseDescriptionEN: courseInfo.CourseDescriptionEN,
			CourseObjective:     courseInfo.CourseObjective,
			Location:            courseInfo.Location,
			IsEditedCourse:      pointer.ToBool(false),
			IsNewCourse:         pointer.ToBool(false),
		})
	}

	courseSource = programConstant.COURSE_SOURCE_REG
	for _, course := range regCourses {
		revision := pointer.GetString(course.RevisionCode)
		if revision == "" {
			revision = "1"
		}
		courseName := strings.Join([]string{
			courseSource,
			"แก้ไขครั้งที่ " + revision,
			pointer.GetString(course.CourseCode),
			pointer.GetString(course.CourseName),
			pointer.GetString(course.CourseNameEng),
		}, " | ")
		period1, _ := course.Period1.Int64()
		period2, _ := course.Period2.Int64()
		period3, _ := course.Period3.Int64()
		creditTotal, _ := strconv.Atoi(pointer.GetString(course.CreditTotal))
		courseCode := pointer.GetString(course.CourseCode)
		if match := FormalCourseCodeStyleRegex.MatchString(courseCode); !match {
			courseCode = strings.ReplaceAll(courseCode, " ", "")
			if match := CascualCourseCodeStyleRegex.FindStringSubmatch(courseCode); match != nil {
				match[1] = strings.ToUpper(match[1])
				courseCode = strings.Join(match[1:len(match)], " ")
			}
		}
		creditDetail := fmt.Sprintf("(%d - %d - %d)", uint(period1), uint(period2), uint(period3))
		courseDescription := strings.Join([]string{pointer.GetString(course.Description1), pointer.GetString(course.Description2), pointer.GetString(course.Description3)}, "\n")
		courseDescriptionEng := strings.Join([]string{pointer.GetString(course.DescriptionEng1), pointer.GetString(course.DescriptionEng2), pointer.GetString(course.DescriptionEng3)}, "\n")
		courseInfoResult = append(courseInfoResult, dto.CoursePaginationForStructure{
			ID:                  nil,
			REGKkuKey:           course.CourseID,
			Name:                &courseName,
			CategoryName:        course.CourseGroup,
			CourseCode:          &courseCode,
			CourseSource:        pointer.ToString(courseSource),
			CourseYear:          pointer.ToString("-"),
			CourseNameTH:        course.CourseName,
			CourseNameEN:        course.CourseNameEng,
			Version:             course.RevisionCode,
			CourseCredit:        pointer.ToUint(uint(creditTotal)),
			CreditDetail:        &creditDetail,
			Credit1:             pointer.ToUint(uint(period1)),
			Credit2:             pointer.ToUint(uint(period2)),
			Credit3:             pointer.ToUint(uint(period3)),
			CourseDescriptionTH: &courseDescription,
			CourseDescriptionEN: &courseDescriptionEng,
			Location:            pointer.ToString(constant.UNIVERSITY_NAME_TH),
			IsEditedCourse:      pointer.ToBool(false),
			IsNewCourse:         pointer.ToBool(false),
		})
	}

	result = &dto.GetCoursePaginationForStructureResponseDto{
		Items:             courseInfoResult,
		PaginationOptions: &options,
	}

	return
}
