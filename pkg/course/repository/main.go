package repository

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	model "github.com/zercle/kku-qf-services/pkg/models/migrate_models"

	"gorm.io/gorm"
)

type courseRepository struct {
	MainDbConn *gorm.DB
}

func NewCourseRepository(mainDbConn *gorm.DB) domain.CourseRepository {
	return &courseRepository{
		MainDbConn: mainDbConn,
	}
}

// func for migrate table and forget to call this func in router.go
func (repo *courseRepository) DbCourseSVCMigrator() (err error) {
	err = repo.MainDbConn.AutoMigrate(
		&model.Course{},
		&model.CourseInfo{},
		&model.CourseLecturer{},
		&model.CourseResult{},
		&model.CourseTypeAndManagement{},
		&model.CourseAssessment{},
		&model.CoursePlan{},
		&model.MapCourseLecturer{},
	)
	return
}
