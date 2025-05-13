package domain

import "github.com/zercle/kku-qf-services/pkg/models/dto"

type QF4Usecase interface {
	GetQF4Assessment(qf4UidString string) (res *dto.QF4AssessmentResponseDto, err error)
	CreateOrUpdateQF4Assessment(data *dto.QF4CreateAssessmentRequestDto) (res *dto.QF4AssessmentResponseDto, err error)
	UpdateQF4Assessment(data *dto.QF4UpdateAssessmentRequestDto) (res *dto.QF4AssessmentResponseDto, err error)
	GetQF4CoursePlan(qf4UidString string) (res []dto.QF4CoursePlanResponseDto, err error)
	CreateQF4CoursePlan(data *dto.QF4CreateCoursePlanRequestDto) (res *dto.QF4CoursePlanResponseDto, err error)
	UpdateQF4CoursePlan(data *dto.QF4UpdateCoursePlanRequestDto) (res *dto.QF4CoursePlanResponseDto, err error)
	CreateQF4Main(data *dto.QF4CreateMainRequestDto) (res *dto.QF4MainResponseDto, err error)
	UpdateQF4Main(data *dto.QF4UpdateMainRequestDto) (res *dto.QF4MainResponseDto, err error)
	DeleteQF4Main(data *dto.QF4DeleteMainRequestDto) (err error)
	GetQF4Main(data *dto.QF4GetMainRequestDto) (response *dto.QF4MainResponseDto, err error)
	CreateOrUpdateQF4CourseInfo(data *dto.QF4CreateCourseInfoRequestDto) (res *dto.QF4CourseInfoResponseDto, err error)
	UpdateQF4CourseInfo(data *dto.QF4UpdateCourseInfoRequestDto) (res *dto.QF4CourseInfoResponseDto, err error)
	GetQF4CourseInfo(courseUidString string) (res *dto.QF4CourseInfoResponseDto, err error)
	CreateOrUpdateQF4Lecturer(data *dto.QF4CreateLecturerRequestDto) (res *dto.QF4LecturerResponseDto, err error)
	UpdateQF4Lecturer(data *dto.QF4UpdateLecturerRequestDto) (res *dto.QF4LecturerResponseDto, err error)
	GetQF4Lecturer(courseUidString string) (res *dto.QF4LecturerResponseDto, err error)
	CreateOrUpdateQF4Result(data *dto.QF4CreateResultRequestDto) (res *dto.QF4ResultResponseDto, err error)
	UpdateQF4Result(data *dto.QF4UpdateResultRequestDto) (res *dto.QF4ResultResponseDto, err error)
	GetQF4Result(courseUidString string) (res *dto.QF4ResultResponseDto, err error)
}

type QF4Repository interface {
	DbQF4SVCMigrator() (err error)
	GetMapQF4Lecturer(id int) (response []dto.MapQF4Lecturer, err error)
}
