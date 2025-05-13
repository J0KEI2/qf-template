package domain

type TemplateUseCase interface {
	FetchTemplate(entity interface{}) (interface{}, error)
}

type TemplateRepository interface {
	DbTemplateSVCMigrator() (err error)
	FetchTemplate(queryEntity interface{}) (interface{}, error)
}
