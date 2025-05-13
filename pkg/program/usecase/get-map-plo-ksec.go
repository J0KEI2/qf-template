package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

func (u programUsecase) GetPloMapWithKsec(ProgramSubPlanId uint, ploID uint) (ksec *dto.KsecResponseDto, err error) {
	programPloQuery := query.ProgramPloQueryEntity{
		ID: &ploID,
	}

	if err = u.CommonRepository.GetFirst(&programPloQuery, "ProgramMapPloWithKsecQueryEntity"); err != nil {
		return
	}

	programKsecQuery := query.ProgramKsecDetailQueryEntity{
		ProgramSubPlanID: &ProgramSubPlanId,
	}

	programKsec := []query.ProgramKsecDetailQueryEntity{}

	if err = u.CommonRepository.GetList(&programKsecQuery, &programKsec, nil); err != nil {
		return
	}

	programKsec = checkKsecCheckedWithPLO(programKsec, programPloQuery.ProgramMapPloWithKsecQueryEntity)

	ksec = new(dto.KsecResponseDto)
	ksec.Init()
	ksec.FromKsecDetail(programKsec)
	ksec.SortByOrder()

	return
}

func checkKsecCheckedWithPLO(ksecDetails []query.ProgramKsecDetailQueryEntity, MapPloWithKsecIds []query.ProgramMapPloWithKsecQueryEntity) []query.ProgramKsecDetailQueryEntity {
	for index := range ksecDetails {
		isChecked, ProgramMapPloWithKsecID := isIdInMapPloWithKsecId(MapPloWithKsecIds, *ksecDetails[index].ID)
		ksecDetails[index].IsChecked = isChecked
		ksecDetails[index].ProgramMapPloWithKsecID = ProgramMapPloWithKsecID
	}
	return ksecDetails
}

func isIdInMapPloWithKsecId(MapPloWithKsecIds []query.ProgramMapPloWithKsecQueryEntity, id uint) (*bool, *uint) {
	found := true
	for _, MapPloWithKsecId := range MapPloWithKsecIds {
		if *MapPloWithKsecId.KsecID == id {
			return &found, MapPloWithKsecId.ID
		}
	}
	found = false
	return &found, nil
}
