package usecase

import (
	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) UpdateCourse(request dto.ProgramCourseDetailRequestDto, courseId uint) (result *dto.ProgramCourseDetailResponseDto, err error) {
	courseCredit, _ := request.CourseCredit.Int64()
	courseStatement := query.ProgramCourseDetailQueryEntity{
		ID: &courseId,
	}
	courseUpdate := query.ProgramCourseDetailQueryEntity{
		YearAndSemesterID:   nil,
		ProgramStructureID:  request.ProgramStructureID,
		ProgramSubPlanID:    request.ProgramSubPlanID,
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
	err = helper.ExecuteTransaction(u.CommonRepository, u.UpdateCourseTransaction(&courseStatement, &courseUpdate))
	if err != nil {
		return nil, err
	}

	result = &dto.ProgramCourseDetailResponseDto{
		ID:                courseUpdate.ID,
		YearAndSemesterID: courseUpdate.YearAndSemesterID,
		CourseTypeID:      courseUpdate.CourseTypeID,
		CourseType:        courseUpdate.CourseType,
		CourseCode:        courseUpdate.CourseCode,
		CourseYear:        courseUpdate.CourseYear,
		CourseNameTH:      courseUpdate.CourseNameTH,
		CourseNameEN:      courseUpdate.CourseNameEN,
		CourseCredit:      courseUpdate.CourseCredit,
		Credit1:           courseUpdate.Credit1,
		Credit2:           courseUpdate.Credit2,
		Credit3:           courseUpdate.Credit3,
		IsCreditCalc:      courseUpdate.IsCreditCalc,
		CreatedAt:         courseUpdate.CreatedAt,
		UpdatedAt:         courseUpdate.UpdatedAt,
		DeletedAt:         courseUpdate.DeletedAt,
	}
	return
}

func (u programUsecase) UpdateCourseTransaction(statement *query.ProgramCourseDetailQueryEntity, update *query.ProgramCourseDetailQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return u.CommonRepository.Update(tx, statement, update)
	}
}
