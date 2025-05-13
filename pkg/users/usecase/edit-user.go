package usecase

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/zercle/kku-qf-services/pkg/models"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

func (u *userUsecase) EditUser(userID string, user models.PatchUserRequest) (err error) {
	newUserData := new(migrateModels.Users)
	byteData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(byteData, newUserData); err != nil {
		return err
	}

	log.Println(userID)
	log.Println(user)

	newUserData.NameTH = strings.Join([]string{user.TitleTH, user.FirstNameTH, user.LastNameTH}, " ")
	newUserData.NameEN = strings.Join([]string{user.TitleEN, user.FirstNameEN, user.LastNameEN}, " ")

	err = u.userRepo.EditUser(userID, *newUserData)
	if err != nil {
		return err
	}

	return nil
}
