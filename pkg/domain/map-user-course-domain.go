package domain

type MapUserCourseUsecase interface {
}

type MapUserCourseRepository interface {
	DbMapUserCourseSVCMigrator() (err error)
}
