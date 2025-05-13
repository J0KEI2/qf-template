package usecase

import (
	"log"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
	"github.com/zercle/kku-qf-services/pkg/utils"
)

func (u *userUsecase) GetUserByID(uid uuid.UUID) (*models.UserFetchModel, error) {
	userStatement := query.UserQueryEntity{
		UID: &uid,
	}
	err := u.commonRepo.GetFirst(&userStatement, "Faculty")
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	nameTHList := append(strings.SplitN(pointer.GetString(userStatement.NameTH), " ", 3), []string{"", ""}...)
	nameENList := append(strings.SplitN(pointer.GetString(userStatement.NameEN), " ", 3), []string{"", ""}...)

	user := &models.UserFetchModel{
		CreatedAt:               utils.PtoTime(userStatement.CreatedAt),
		LastAccessAt:            utils.PtoTime(userStatement.LastAccessAt),
		UpdatedAt:               utils.PtoTime(userStatement.UpdatedAt),
		Email:                   userStatement.Email,
		TitleTH:                 pointer.ToString(nameTHList[0]),
		FirstNameTH:             pointer.ToString(nameTHList[1]),
		LastNameTH:              pointer.ToString(nameTHList[2]),
		TitleEN:                 pointer.ToString(nameENList[0]),
		FirstNameEN:             pointer.ToString(nameENList[1]),
		LastNameEN:              pointer.ToString(nameENList[2]),
		Status:                  userStatement.Status,
		SSN:                     userStatement.SSN,
		Type:                    userStatement.Type,
		FacultyID:               userStatement.FacultyID,
		SystemPermissionUID:     userStatement.SystemPermissionUID,
		UID:                     userStatement.UID,
		ProgramApprovalMinLevel: userStatement.ProgramApprovalMinLevel,
		ProgramApprovalMaxLevel: userStatement.ProgramApprovalMaxLevel,
		CurrentRoleID:           *userStatement.CurrentRoleID,
	}
	if userStatement.Faculty != nil {
		user.FacultyName = userStatement.Faculty.FacultyNameTH
	}
	return user, nil
}
