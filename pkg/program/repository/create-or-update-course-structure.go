package repository

import (
	"fmt"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	programConstant "github.com/zercle/kku-qf-services/pkg/constant/program"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	courseQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (r programRepository) CreateOrUpdateCourseStructure(tx *gorm.DB, structures []dto.ProgramStructureRequestDto, subPlanID *uint, parentID *uint) (err error) {
	for _, structure := range structures {
		var structureCredit *uint = nil
		if structure.StructureCredit != nil {
			creditInt, _ := structure.StructureCredit.Int64()
			structureCredit = pointer.ToUint(uint(creditInt))
		}
		queryStatement := query.ProgramStructureDetailQueryEntity{
			ID: structure.ID,
		}
		update := query.ProgramStructureDetailQueryEntity{
			Name:             structure.Name,
			ProgramSubPlanID: subPlanID,
			Order:            structure.Order,
			ParentID:         parentID,
			Qualification:    structure.Qualification,
			StructureCredit:  structureCredit,
		}
		if err = tx.Where(queryStatement).Updates(&update).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				err = tx.Create(&update).Error
				queryStatement.ID = update.ID
			}
			if err != nil {
				return err
			}
		}
		
		if err = r.CreateOrUpdateCourse(tx, structure.CourseDetails, queryStatement.ID, subPlanID); err != nil {
			return err
		}
		if err = r.CreateOrUpdateCourseStructure(tx, structure.Children, subPlanID, queryStatement.ID); err != nil {
			return err
		}
	}
	return nil
}

func (r programRepository) CreateOrUpdateCourse(tx *gorm.DB, courses []dto.ProgramCourseDetailRequestDto, structureID *uint, subPlanID *uint) (err error) {
	for _, course := range courses {
		var courseCreditPointer *uint = nil
		if course.CourseCredit != nil {
			courseCredit, _ := (course.CourseCredit.Int64())
			courseCreditPointer = pointer.ToUint(uint(courseCredit))
		}
		update := query.ProgramCourseDetailQueryEntity{
			ProgramStructureID:  structureID,
			ProgramSubPlanID:    subPlanID,
			CourseType:          course.CourseType,
			CourseTypeID:        course.CourseTypeID,
			CourseCode:          course.CourseCode,
			CourseSource:        course.CourseSource,
			CourseYear:          course.CourseYear,
			REGKkuKey:           course.REGKkuKey,
			CourseKey:           course.CourseKey,
			CourseNameTH:        course.CourseNameTH,
			CourseNameEN:        course.CourseNameEN,
			Version:             course.Version,
			CourseCredit:        courseCreditPointer,
			Credit1:             course.Credit1,
			Credit2:             course.Credit2,
			Credit3:             course.Credit3,
			CourseConditionTH:   course.CourseConditionTH,
			CourseConditionEN:   course.CourseConditionEN,
			CourseDescriptionEN: course.CourseDescriptionEN,
			CourseDescriptionTH: course.CourseDescriptionTH,
			CourseObjective:     course.CourseObjective,
			IsEditedCourse:      course.IsEditedCourse,
			IsNewCourse:         course.IsNewCourse,
			IsCreditCalc:        course.IsCreditCalc,
		}
		whereClause := &query.ProgramCourseDetailQueryEntity{
			ID: course.ID,
		}
		if course.ID != nil {
			err = tx.Where(whereClause).Updates(&update).Error
		} else {
			err = r.CreateCourse(tx, &update)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (r programRepository) CreateCourse(tx *gorm.DB, course *query.ProgramCourseDetailQueryEntity) (err error) {
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
	err = tx.Create(&course).Error
	return err
}
