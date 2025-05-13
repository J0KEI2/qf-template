package usecase

func (u *userUsecase) DeleteUser(userID string) (err error) {
	err = u.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}
