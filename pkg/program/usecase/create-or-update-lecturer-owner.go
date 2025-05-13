package usecase

import (
	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateLecturerOwner(lecturerOwner dto.CreateOrUpdateLecturerOwnerDto, programMainUID uuid.UUID) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateLecturerOwnerTransaction(lecturerOwner, programMainUID))
}

func (u programUsecase) CreateOrUpdateLecturerOwnerTransaction(lecturerOwner dto.CreateOrUpdateLecturerOwnerDto, programMainUID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {

		for _, owner := range lecturerOwner.Owner {
			EmployeeUid, err := u.updateEducation(owner)
			if err != nil {
				continue
			}
			if owner.ID == nil {
				create := query.ProgramOwnerQueryEntity{
					ID:         nil,
					ProgramMainUID: &programMainUID,
					OwnerID:    EmployeeUid,
				}
				err = u.CommonRepository.Create(tx, &create)

				if err != nil {
					continue
				}
			}
		}

		for _, lecturer := range lecturerOwner.Lecturer {
			EmployeeUid, err := u.updateEducation(lecturer)
			if err != nil {
				continue
			}
			if lecturer.ID == nil {
				create := query.ProgramLecturerQueryEntity{
					ID:         nil,
					ProgramMainUID: &programMainUID,
					LecturerID: EmployeeUid,
				}
				err = u.CommonRepository.Create(tx, &create)

				if err != nil {
					return err
				}
			}
		}

		for _, thesisLecturer := range lecturerOwner.ThesisLecturer {
			EmployeeUid, err := u.updateEducation(thesisLecturer)
			if err != nil {
				continue
			}
			if thesisLecturer.ID == nil {
				create := query.ProgramThesisLecturerQueryEntity{
					ID:               nil,
					ProgramMainUID:       &programMainUID,
					ThesisLecturerID: EmployeeUid,
				}
				err = u.CommonRepository.Create(tx, &create)

				if err != nil {
					return err
				}
			}
		}

		return nil
	}
}

func (u programUsecase) updateEducation(employee dto.LecturerDto) (*uuid.UUID, error) {
	educations, err := u.HRUseCase.GetEducationByEmail(*employee.Email)
	if err != nil {
		return nil, err
	}
	employee.EducationBackgrounds = []dto.EducationalBackgroundDto{}
	for _, educationPtr := range educations {
		education := educationPtr
		employee.EducationBackgrounds = append(employee.EducationBackgrounds, dto.EducationalBackgroundDto{
			EducationLevel:   &education.EducationLevel,
			Qualification:    &education.Qualification,
			Department:       &education.Major,
			InstituteName:    &education.Institute,
			InstituteCountry: &education.Country,
			GraduateYear:     &education.SuccessYear,
		})
	}
	EmployeeUid, err := u.ProgramRepository.GetOrCreateEmployeeByEmail(employee)
	if err != nil {
		return nil, err
	}
	return EmployeeUid, nil
}
