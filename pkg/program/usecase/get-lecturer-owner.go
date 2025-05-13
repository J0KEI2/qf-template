package usecase

import (
	"encoding/json"
	"errors"
	"log"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetLecturerOwner(ProgramID uuid.UUID) (result *dto.GetLecturerOwnerDto, err error) {
	mainQuery := query.ProgramMainQueryEntity{
		ID: &ProgramID,
	}

	if err = u.CommonRepository.GetFirst(&mainQuery, "ProgramOwner", "ProgramOwner.Owner", "ProgramLecturer", "ProgramLecturer.Lecturer", "ProgramThesisLecturer", "ProgramThesisLecturer.ThesisLecturer"); err != nil {
		return nil, errors.New("lecturer not found")
	}

	owners := make([]dto.LecturerDto, 0)

	for _, owner := range mainQuery.ProgramOwner {
		if owner.Owner != nil {
			user := owner.Owner
			name := strings.Join([]string{*user.TitleTh, *user.FirstnameTh, *user.LastnameTh}, " ")
			nameEN := strings.Join([]string{*user.TitleEn, *user.FirstnameEn, *user.LastnameEn}, " ")
			position := []dto.LecturerPositionDto{}
			educationBackground := []dto.EducationalBackgroundDto{}
			if user.Position != nil {
				json.Unmarshal([]byte(*user.Position), &position)
				sort.SliceStable(position, func(i, j int) bool {
					return *position[i].Year > *position[j].Year
				})
			}
			if user.EducationBackGround != nil {
				json.Unmarshal([]byte(*user.EducationBackGround), &educationBackground)
				sort.SliceStable(educationBackground, func(i, j int) bool {
					return *educationBackground[i].GraduateYear > *educationBackground[j].GraduateYear
				})
			}
			owners = append(owners, dto.LecturerDto{
				ID:                   owner.ID,
				UserUID:              user.UID,
				Email:                user.Email,
				Name:                 &name,
				NameEN:               &nameEN,
				TitleTh:              user.TitleTh,
				TitleEn:              user.TitleEn,
				FirstnameTh:          user.FirstnameTh,
				FirstnameEn:          user.FirstnameEn,
				LastnameTh:           user.LastnameTh,
				LastnameEn:           user.LastnameEn,
				Positions:            position,
				EducationBackgrounds: educationBackground,
			})
		}
	}

	lecturers := make([]dto.LecturerDto, 0)
	for _, lecturer := range mainQuery.ProgramLecturer {
		if lecturer.Lecturer != nil {
			user := lecturer.Lecturer
			name := strings.Join([]string{*user.TitleTh, *user.FirstnameTh, *user.LastnameTh}, " ")
			nameEN := strings.Join([]string{*user.TitleEn, *user.FirstnameEn, *user.LastnameEn}, " ")
			position := []dto.LecturerPositionDto{}
			educationBackground := []dto.EducationalBackgroundDto{}
			if user.Position != nil {
				json.Unmarshal([]byte(*user.Position), &position)
				sort.SliceStable(position, func(i, j int) bool {
					return *position[i].Year > *position[j].Year
				})
			}
			if user.EducationBackGround != nil {
				json.Unmarshal([]byte(*user.EducationBackGround), &educationBackground)
				log.Println(educationBackground)
				sort.SliceStable(educationBackground, func(i, j int) bool {
					return *educationBackground[i].GraduateYear > *educationBackground[j].GraduateYear
				})
			}
			lecturers = append(lecturers, dto.LecturerDto{
				ID:                   lecturer.ID,
				UserUID:              user.UID,
				Email:                user.Email,
				Name:                 &name,
				NameEN:               &nameEN,
				TitleTh:              user.TitleTh,
				TitleEn:              user.TitleEn,
				FirstnameTh:          user.FirstnameTh,
				FirstnameEn:          user.FirstnameEn,
				LastnameTh:           user.LastnameTh,
				LastnameEn:           user.LastnameEn,
				Positions:            position,
				EducationBackgrounds: educationBackground,
			})
		}
	}

	thesisLecturers := make([]dto.LecturerDto, 0)
	for _, thesisLecturer := range mainQuery.ProgramThesisLecturer {
		if thesisLecturer.ThesisLecturer != nil {
			user := thesisLecturer.ThesisLecturer
			name := strings.Join([]string{*user.TitleTh, *user.FirstnameTh, *user.LastnameTh}, " ")
			nameEN := strings.Join([]string{*user.TitleEn, *user.FirstnameEn, *user.LastnameEn}, " ")
			position := []dto.LecturerPositionDto{}
			educationBackground := []dto.EducationalBackgroundDto{}
			if user.Position != nil {
				json.Unmarshal([]byte(*user.Position), &position)
				sort.SliceStable(position, func(i, j int) bool {
					return *position[i].Year > *position[j].Year
				})
			}
			if user.EducationBackGround != nil {
				json.Unmarshal([]byte(*user.EducationBackGround), &educationBackground)
				sort.SliceStable(educationBackground, func(i, j int) bool {
					return *educationBackground[i].GraduateYear > *educationBackground[j].GraduateYear
				})
			}
			thesisLecturers = append(thesisLecturers, dto.LecturerDto{
				ID:                   thesisLecturer.ID,
				UserUID:              user.UID,
				Email:                user.Email,
				Name:                 &name,
				NameEN:               &nameEN,
				TitleTh:              user.TitleTh,
				TitleEn:              user.TitleEn,
				FirstnameTh:          user.FirstnameTh,
				FirstnameEn:          user.FirstnameEn,
				LastnameTh:           user.LastnameTh,
				LastnameEn:           user.LastnameEn,
				Positions:            position,
				EducationBackgrounds: educationBackground,
			})
		}
	}

	result = &dto.GetLecturerOwnerDto{
		Owner:          owners,
		ThesisLecturer: thesisLecturers,
		Lecturer:       lecturers,
	}
	return
}
