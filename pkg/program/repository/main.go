package repository

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
	"gorm.io/gorm"
)

type programRepository struct {
	MainDbConn *gorm.DB
}

func NewProgramRepository(mainDbConn *gorm.DB) domain.ProgramRepository {
	return &programRepository{
		MainDbConn: mainDbConn,
	}
}

// func for migrate table and forget to call this func in router.go
func (repo *programRepository) DbProgramSVCMigrator() (err error) {
	err = repo.MainDbConn.AutoMigrate(
		&migrateModels.ProgramMain{},
		&migrateModels.ProgramGeneralDetail{},
		&migrateModels.ProgramCompetency{},
		&migrateModels.ProgramMajor{},
		&migrateModels.ProgramPlanDetail{},
		&migrateModels.ProgramSubPlan{},
		&migrateModels.ProgramStructureDetail{},
		&migrateModels.ProgramYearAndSemester{},
		&migrateModels.ProgramCourseDetail{},
		&migrateModels.ProgramPolicyAndStrategic{},
		&migrateModels.ProgramPLOFormat{},
		&migrateModels.ProgramPlo{},
		&migrateModels.ProgramMapPloWithKsec{},
		&migrateModels.ProgramKsecDetail{},
		&migrateModels.ProgramPLOLearningSolution{},
		&migrateModels.ProgramPLOLearningEvaluation{},
		&migrateModels.ProgramQualityAssurance{},
		&migrateModels.ProgramSystemAndMechanic{},
		&migrateModels.ProgramYloKsec{},
		&migrateModels.ProgramYloWithPlo{},
		&migrateModels.ProgramYloWithKsec{},
		&migrateModels.ProgramThesisLecturer{},
		&migrateModels.ProgramLecturer{},
		&migrateModels.ProgramOwner{},
		&migrateModels.ProgramPlanAndEvaluate{},
		&migrateModels.ProgramReference{},
		&migrateModels.ProgramMapCurMapResp{},
		&migrateModels.ProgramMapCurMapKsa{},
		&migrateModels.ProgramKsaDetail{},
	)
	return
}
