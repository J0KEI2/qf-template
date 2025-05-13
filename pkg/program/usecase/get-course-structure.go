package usecase

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) GetStructureAndCourse(subPlanID uint, paginationOptions *models.PaginationOptions) (result *dto.GetCourseStructureResponseDto, err error) {
	paginationOptions.DefaultLimit(0)

	structureResult, err := u.recursiveStructureFormat(subPlanID, pointer.ToUint(0), paginationOptions) // get nil parentID
	if err != nil {
		return nil, err
	}
	result = &dto.GetCourseStructureResponseDto{
		Items:             structureResult,
		PaginationOptions: paginationOptions,
	}
	return result, nil
}

func (u programUsecase) recursiveStructureFormat(subPlanID uint, parentID *uint, paginationOptions *models.PaginationOptions) (structuresResult []dto.ProgramStructureResponseDto, err error) {
	queryDb := query.ProgramStructureDetailQueryEntity{
		ProgramSubPlanID: &subPlanID,
		ParentID:         parentID,
	}

	structures := []query.ProgramStructureDetailQueryEntity{}

	if err = u.CommonRepository.GetListWithNilSearch(&queryDb, &structures, paginationOptions, "CourseDetail", "CourseDetail.CourseMain"); err != nil {
		return nil, err
	}

	CascualCourseCodeStyleRegex, _ := regexp.Compile(`^([A-Za-z]{2})(\d{3})(\d{3})$`)
	FormalCourseCodeStyleRegex, _ := regexp.Compile(`^[A-Z]{2} \d{3} \d{3}$`)
	for _, structureCourse := range structures {
		courses := make([]dto.ProgramCourseDetailResponseDto, 0)
		for _, course := range structureCourse.CourseDetail {
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
				IsEditedCourse:    course.IsEditedCourse,
				IsNewCourse:       course.IsNewCourse,
				CreatedAt:         course.CreatedAt,
				UpdatedAt:         course.UpdatedAt,
				DeletedAt:         course.DeletedAt,
			})
		}
		children, err := u.recursiveStructureFormat(subPlanID, structureCourse.ID, paginationOptions)
		if err != nil {
			return nil, err
		}
		if children == nil {
			children = make([]dto.ProgramStructureResponseDto, 0)
		}
		sort.SliceStable(courses, func(i, j int) bool {
			return *courses[i].ID < *courses[j].ID
		})
		structuresResult = append(structuresResult, dto.ProgramStructureResponseDto{
			ID:               structureCourse.ID,
			ProgramSubPlanID: structureCourse.ProgramSubPlanID,
			Name:             structureCourse.Name,
			Order:            structureCourse.Order,
			ParentID:         structureCourse.ParentID,
			Children:         children,
			CourseDetails:    courses,
			Qualification:    structureCourse.Qualification,
			StructureCredit:  structureCourse.StructureCredit,
			CreatedAt:        structureCourse.CreatedAt,
			UpdatedAt:        structureCourse.UpdatedAt,
			DeletedAt:        structureCourse.DeletedAt,
		})
	}
	sort.SliceStable(structuresResult, func(i, j int) bool {
		return pointer.GetUint(structuresResult[i].Order) < pointer.GetUint(structuresResult[j].Order)
	})
	return structuresResult, nil
}

func (u programUsecase) UpdateCourseFromCoursetoRegKKU(course query.ProgramCourseDetailQueryEntity) (err error) {
	if course.CourseSource == nil || *course.CourseSource != "Course" || course.CourseMain.REGKkuKey == nil {
		return nil
	}
	log.Println("Update start from id : ", course.ID)
	return helper.ExecuteTransaction(u.CommonRepository, u.updateREGKkuKeyTransaction(course))
}

func (u programUsecase) updateREGKkuKeyTransaction(course query.ProgramCourseDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		err = u.CommonRepository.Update(tx, query.ProgramCourseDetailQueryEntity{
			ID: course.ID,
		}, &query.ProgramCourseDetailQueryEntity{
			CourseSource: pointer.ToString("REG"),
			REGKkuKey:    course.CourseMain.REGKkuKey,
		})
		return err
	}
}
