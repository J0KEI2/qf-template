package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/course"
)

func (u courseUsecase) GetMainCoursePagination(options models.PaginationOptions) (result *dto.GetMainCoursePaginationResponseDto, err error) {

	statement := query.CourseMainQueryEntity{}

	snapshot := []query.CourseMainQueryEntity{}

	err = u.CommonRepository.GetList(&statement, &snapshot, &options, "CourseInfo")

	if err != nil {
		return nil, err
	}
	items := make([]dto.CourseMainWithInfoResponseDto, 0)

	for _, course := range snapshot {
		main := course
		items = append(items, dto.CourseMainWithInfoResponseDto{
			CourseID:       main.CourseID,
			CourseNumber:   main.CourseNumber,
			Version:        main.Version,
			DepartmentName: main.DepartmentName,
			EducationYear:  main.EducationYear,
			Status:         main.Status,
			CourseCode:     main.CourseInfo.CourseCode,
			CourseNameTh:   main.CourseInfo.CourseNameTH,
			CourseNameEn:   main.CourseInfo.CourseNameEN,
			CategoryName:   main.CourseInfo.CategoryName,
			TotalCredit:    main.CourseInfo.TotalCredit,
			Credit1:        main.CourseInfo.Credit1,
			Credit2:        main.CourseInfo.Credit2,
			Credit3:        main.CourseInfo.Credit3,
		})
	}

	result = &dto.GetMainCoursePaginationResponseDto{
		Items:             items,
		PaginationOptions: options,
	}

	return
}
