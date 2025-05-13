package usecase

import "github.com/zercle/kku-qf-services/pkg/models"

func (u *userUsecase) GetUserByIDWithRoleSystem(uid string) (*models.UserFetchWithSystemRoleModel, error) {
	return u.userRepo.GetUserByIDWithRoleSystem(uid)
}
