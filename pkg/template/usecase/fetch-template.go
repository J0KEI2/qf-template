package templateUsecase

// TODO: create the new func don't forget to embed in domain use case interface
func (useCase templateUsecase) FetchTemplate(entity interface{}) (interface{}, error) {
	return useCase.TemplateRepository.FetchTemplate(entity)
}

// func (u *userUsecase) GetUser(userID string) (user models.User, err error) {
// 	return u.userRepo.GetUser(userID)
// }
