package usecase

import (
	"fmt"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	programConstant "github.com/zercle/kku-qf-services/pkg/constant/program"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	courseQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateCourse(request dto.ProgramCourseDetailRequestDto, subPlanID uint, structureId uint) (result *dto.ProgramCourseDetailResponseDto, err error) {
	courseCredit, _ := request.CourseCredit.Int64()
	courseQuery := query.ProgramCourseDetailQueryEntity{
		YearAndSemesterID:   nil,
		ProgramStructureID:  &structureId,
		ProgramSubPlanID:    &subPlanID,
		CourseSource:        request.CourseSource,
		REGKkuKey:           request.REGKkuKey,
		CourseKey:           request.CourseKey,
		CourseTypeID:        request.CourseTypeID,
		CourseType:          request.CourseType,
		CourseCode:          request.CourseCode,
		CourseYear:          request.CourseYear,
		CourseNameTH:        request.CourseNameTH,
		CourseNameEN:        request.CourseNameEN,
		Version:             request.Version,
		CourseCredit:        pointer.ToUint(uint(courseCredit)),
		Credit1:             request.Credit1,
		Credit2:             request.Credit2,
		Credit3:             request.Credit3,
		CourseConditionTH:   request.CourseConditionTH,
		CourseConditionEN:   request.CourseConditionEN,
		CourseDescriptionEN: request.CourseDescriptionEN,
		CourseDescriptionTH: request.CourseDescriptionTH,
		CourseObjective:     request.CourseObjective,
		IsCreditCalc:        request.IsCreditCalc,
		IsEditedCourse:      request.IsEditedCourse,
		IsNewCourse:         request.IsNewCourse,
	}
	err = helper.ExecuteTransaction(u.CommonRepository, u.CreateCourseTransaction(&courseQuery))
	if err != nil {
		return nil, err
	}

	result = &dto.ProgramCourseDetailResponseDto{
		ID:                courseQuery.ID,
		YearAndSemesterID: courseQuery.YearAndSemesterID,
		CourseTypeID:      courseQuery.CourseTypeID,
		CourseType:        courseQuery.CourseType,
		CourseCode:        courseQuery.CourseCode,
		CourseYear:        courseQuery.CourseYear,
		CourseNameTH:      courseQuery.CourseNameTH,
		CourseNameEN:      courseQuery.CourseNameEN,
		CourseCredit:      courseQuery.CourseCredit,
		Credit1:           courseQuery.Credit1,
		Credit2:           courseQuery.Credit2,
		Credit3:           courseQuery.Credit3,
		IsCreditCalc:      courseQuery.IsCreditCalc,
		CreatedAt:         courseQuery.CreatedAt,
		UpdatedAt:         courseQuery.UpdatedAt,
		DeletedAt:         courseQuery.DeletedAt,
	}
	return
}

func (u programUsecase) CreateCourseTransaction(course *query.ProgramCourseDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		isEditedCourse := pointer.GetBool(course.IsEditedCourse)
		isNewCourse := pointer.GetBool(course.IsNewCourse)
		if isEditedCourse || isNewCourse {
			//Todo : create course
			courseCourseInfo := courseQuery.CourseInfoQueryEntity{
				CategoryName:        pointer.ToString("Lecture"),
				CourseCode:          course.CourseCode,
				CourseNameTH:        course.CourseNameTH,
				CourseNameEN:        course.CourseNameEN,
				TotalCredit:         course.CourseCredit,
				Credit1:             course.Credit1,
				Credit2:             course.Credit2,
				Credit3:             course.Credit3,
				CourseTypeID:        course.CourseTypeID,
				CourseConditionTH:   course.CourseConditionTH,
				CourseConditionEN:   course.CourseConditionEN,
				CourseDescriptionTH: course.CourseDescriptionTH,
				CourseDescriptionEN: course.CourseDescriptionEN,
				CourseObjective:     course.CourseObjective,
				Location:            pointer.ToString(constant.UNIVERSITY_NAME_TH),
			}

			tx.Create(&courseCourseInfo)

			version := 1
			if isEditedCourse {
				oldversion, _ := strconv.Atoi(pointer.GetString(course.Version))
				version = oldversion + 1
			}
			versionString := fmt.Sprintf("%d", version)
			courseNewPointer := uuid.New()
			courseStatement := courseQuery.CourseQueryEntity{
				CourseID:       &courseNewPointer,
				CourseNumber:   pointer.ToInt(0),
				FacultyID:      pointer.ToUint(0),
				DepartmentName: pointer.ToString(""),
				EducationYear:  course.CourseYear,
				CourseInfoID:   courseCourseInfo.ID,
				Status:         pointer.ToString("draft"),
				Version:        pointer.ToString(versionString),
			}

			tx.Create(&courseStatement)

			course.CourseSource = pointer.ToString(programConstant.COURSE_SOURCE_COURSE)
			course.CourseKey = courseStatement.CourseID
			course.Version = pointer.ToString(versionString)
		}
		return u.CommonRepository.Create(tx, course)
	}
}
