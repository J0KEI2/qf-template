package domain

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

type CourseUsecase interface {
	GetMainCoursePagination(options models.PaginationOptions) (result *dto.GetMainCoursePaginationResponseDto, err error)
	GetCourseAssessment(courseUidString string) (res *dto.CourseAssessmentResponseDto, err error)
	CreateOrUpdateCourseAssessment(data *dto.CourseCreateAssessmentRequestDto) (res *dto.CourseAssessmentResponseDto, err error)
	UpdateCourseAssessment(data *dto.CourseUpdateAssessmentRequestDto) (res *dto.CourseAssessmentResponseDto, err error)
	GetCoursePlan(courseUidString string) (res []dto.CoursePlanResponseDto, err error)
	CreateCoursePlan(data *dto.CourseCreateCoursePlanRequestDto) (res *dto.CoursePlanResponseDto, err error)
	UpdateCoursePlan(data *dto.CourseUpdateCoursePlanRequestDto) (res *dto.CoursePlanResponseDto, err error)
	GetCourseMain(data *dto.CourseGetMainRequestDto) (res *dto.CourseMainResponseDto, err error)
	CreateCourseMain(data *dto.CourseCreateMainRequestDto) (res *dto.CourseMainResponseDto, err error)
	UpdateCourseMain(data *dto.CourseUpdateMainRequestDto) (res *dto.CourseMainResponseDto, err error)
	DeleteCourseMain(data *dto.CourseDeleteMainRequestDto) (err error)
	CreateOrUpdateCourseResult(data *dto.CourseCreateResultRequestDto) (res *dto.CourseResultResponseDto, err error)
	UpdateCourseResult(data *dto.CourseUpdateResultRequestDto) (res *dto.CourseResultResponseDto, err error)
	GetCourseResult(courseUidString string) (res *dto.CourseResultResponseDto, err error)
	CreateOrUpdateCourseInfo(data *dto.CourseCreateCourseInfoRequestDto) (res *dto.CourseInfoResponseDto, err error)
	UpdateCourseInfo(data *dto.CourseUpdateCourseInfoRequestDto) (res *dto.CourseInfoResponseDto, err error)
	GetCourseInfo(courseUidString string) (res *dto.CourseInfoResponseDto, err error)
	CreateOrUpdateCourseLecturer(data *dto.CourseCreateLecturerRequestDto) (res *dto.CourseLecturerResponseDto, err error)
	UpdateCourseLecturer(data *dto.CourseUpdateLecturerRequestDto) (res *dto.CourseLecturerResponseDto, err error)
	GetCourseLecturer(courseUidString string) (res *dto.CourseLecturerResponseDto, err error)
}

type CourseRepository interface {
	DbCourseSVCMigrator() (err error)
	GetMapCourseLecturer(id int) (response []dto.MapCourseLecturer, err error)
	GetCourseByUID(criteria interface{}) (responseData interface{}, err error)
}
