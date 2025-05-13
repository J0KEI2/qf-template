package usecase

import (
	"log"

	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
)

func (u commonUsecase) GetFacultiesPagination(options models.PaginationOptions) (result *dto.GetFacultyPaginationResponseDto, err error) {
	statement := query.Faculty{}
	faculties := []query.Faculty{}

	log.Println("order : ", *options.Order)

	options.SetSearchFields([]string{"faculty_name_th", "faculty_name_en"})

	err = u.CommonRepository.GetList(statement, &faculties, &options)

	if err != nil {
		return nil, err
	}
	items := make([]dto.FacultyResponseDto, 0)

	for _, faculty := range faculties {
		items = append(items, dto.FacultyResponseDto{
			ID:            faculty.ID,
			FacultyNameEN: faculty.FacultyNameEN,
			FacultyNameTH: faculty.FacultyNameTH,
			University:    faculty.University,
		})
	}

	result = &dto.GetFacultyPaginationResponseDto{
		Items:   items,
		Options: options,
	}

	return
}
