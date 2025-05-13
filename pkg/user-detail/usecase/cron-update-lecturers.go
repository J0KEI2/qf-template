package usecase

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/entity"
	migrateModels "github.com/zercle/kku-qf-services/pkg/models/migrate_models"
)

// TODO: create the new func don't forget to embed in domain use case interface
func (u *userDetailUseCase) CronUpdateLecturers(accessToken string) error {
	userData, err := u.userDetailRepo.GetAllHrKeyLecturers()
	if err != nil {
		return err
	}

	for _, user := range userData {
		newEmployeeModel := new(models.EmployeeDetail)

		responseBody, err := u.userDetailRepo.GetSingleEmployee(user.HRKey, accessToken)
		if err != nil {
			return err
		}

		err = json.Unmarshal(responseBody, newEmployeeModel)
		if err != nil {
			return err
		}

		if strings.ToLower(newEmployeeModel.Status) == "fail" {
			err = models.NotFoundError()
			return err
		}

		// Mock
		position := []models.UserDetailPostion{
			{
				Position: "PROFESSOR",
				Year:     time.Now().Year(),
			},
		}

		positionByte, err := json.Marshal(position)
		if err != nil {
			return errors.New("position is invalid")
		}
		positionString := string(positionByte)
		educationBG := entity.EducationBackGround{
			Qualification:    "",
			Department:       "โรงพยาบาลศรีนครินทร์",
			InstituteName:    "",
			InstituteCountry: "",
			GraduateYear:     2020,
		}

		educationBgByteData, err := json.Marshal(educationBG)
		if err != nil {
			return err
		}

		educationBgStr := string(educationBgByteData)

		newLecturerData := migrateModels.UserDetail{
			TitleTh:             newEmployeeModel.EmployeeData.Title,
			FirstnameTh:         newEmployeeModel.EmployeeData.FirstName,
			LastnameTh:          newEmployeeModel.EmployeeData.LastName,
			TitleEn:             newEmployeeModel.EmployeeData.TitleEng,
			FirstnameEn:         newEmployeeModel.EmployeeData.FirstNameEng,
			LastnameEn:          newEmployeeModel.EmployeeData.LastNameEng,
			Position:            &positionString,
			EducationBackGround: &educationBgStr,
		}

		err = u.userDetailRepo.UpdateLecturer(user.UID, newLecturerData)
		if err != nil {
			err = models.InternalServerError()
			return err
		}
	}

	return nil
}
