package repository

import (
	"encoding/json"
	"errors"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	"gorm.io/gorm"
)

func (r programRepository) GetOrCreateEmployeeByEmail(employee dto.LecturerDto) (EmployeeUID *uuid.UUID, err error) {
	if r.MainDbConn == nil {
		return nil, errors.New("db is gone")
	}

	db := r.MainDbConn

	var employeeResult *query.EmployeeDetails
	var employeeQuery query.EmployeeDetails
	if employee.Email != nil {
		employeeQuery = query.EmployeeDetails{
			Email: employee.Email,
		}
	} else {
		employeeQuery = query.EmployeeDetails{
			TitleEn:     employee.TitleEn,
			FirstnameEn: employee.FirstnameEn,
			LastnameEn:  employee.LastnameEn,
		}
	}

	if err = db.First(&employeeResult, employeeQuery).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			position, _ := json.Marshal(employee.Positions)
			uid := uuid.New()
			employeeResult = &query.EmployeeDetails{
				UID:         &uid,
				Email:       employee.Email,
				TitleEn:     employee.TitleEn,
				FirstnameEn: employee.FirstnameEn,
				LastnameEn:  employee.LastnameEn,
				TitleTh:     employee.TitleTh,
				FirstnameTh: employee.FirstnameTh,
				LastnameTh:  employee.LastnameTh,
				Position:    pointer.ToString(string(position)),
			}
			if err = db.Model(query.EmployeeDetails{}).Create(&employeeResult).Error; err != nil {
				return nil, err
			}
		}
	}

	educationStringify, _ := json.Marshal(employee.EducationBackgrounds)
	db.Model(&employeeResult).Where(query.EmployeeDetails{
		Email: employee.Email,
	}).Updates(query.EmployeeDetails{
		EducationBackGround: pointer.ToString(string(educationStringify)),
	})
	return employeeResult.UID, nil
}
