package usecase

import (
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
)

func (usecase *hrUseCase) GetLecturerPagination(options models.PaginationOptions) (*dto.GetEmployeesPaginationResponseDto, error) {

	hrToken, err := usecase.GetHRToken()

	if err != nil {
		return nil, err
	}

	firstname := ""
	lastname := ""
	if options.Search != nil {
		search := strings.Split(*options.Search, " ")
		if len(search) > 1 {
			firstname = search[len(search)-2]
			lastname = search[len(search)-1]
		} else {
			firstname = search[0]
		}
	}

	HrEmployees, err := usecase.repo.GetLecturerPagination(*hrToken, firstname, lastname, options)

	if err != nil {
		return nil, err
	}

	employees := make([]dto.HREmployeesResponseDto, 0)

	for _, hrEmpemployee := range HrEmployees.Data.Items {
		name := strings.Join([]string{hrEmpemployee.TitleTh, hrEmpemployee.FirstnameTh, hrEmpemployee.LastnameTh}, " ")
		nameEn := strings.Join([]string{hrEmpemployee.TitleEn, hrEmpemployee.FirstnameEn, hrEmpemployee.LastnameEn}, " ")
		employees = append(employees, dto.HREmployeesResponseDto{
			TitleTh:     hrEmpemployee.TitleTh,
			FirstnameTh: hrEmpemployee.FirstnameTh,
			LastnameTh:  hrEmpemployee.LastnameTh,
			Name:        name,
			TitleEn:     hrEmpemployee.TitleEn,
			FirstnameEn: hrEmpemployee.FirstnameEn,
			LastnameEn:  hrEmpemployee.LastnameEn,
			Email:       hrEmpemployee.Email,
			NameEn:      nameEn,
			Position: []dto.LecturerPositionDto{
				{
					Position: &hrEmpemployee.Position,
					Year:     pointer.ToString("-"),
				},
			},
			Faculty: hrEmpemployee.Faculty,
		})
	}
	response := dto.GetEmployeesPaginationResponseDto{
		Options: options,
		Items:   employees,
	}

	return &response, nil
}
