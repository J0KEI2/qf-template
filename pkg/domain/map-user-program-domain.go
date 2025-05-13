package domain

type MapUserProgramUsecase interface {
}

type MapUserProgramRepository interface {
	DbMapUserProgramSVCMigrator() (err error)
}
