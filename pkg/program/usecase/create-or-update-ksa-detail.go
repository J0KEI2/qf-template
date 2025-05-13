package usecase

import (
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateKsaDetail(ksaDetail dto.CreateOrUpdateKsaDetailRequestDto) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateKsaDetailTransaction(ksaDetail))
}

func (u programUsecase) CreateOrUpdateKsaDetailTransaction(ksaDetail dto.CreateOrUpdateKsaDetailRequestDto) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		// Knowledge
		err = u.updateOrCreateKsaDetail(tx, ksaDetail.Knowledge, ksaDetail.ProgramSubPlanID)
		if err != nil {
			return err
		}

		err = u.updateOrCreateKsaDetail(tx, ksaDetail.Skill, ksaDetail.ProgramSubPlanID)
		if err != nil {
			return err
		}

		err = u.updateOrCreateKsaDetail(tx, ksaDetail.Attitude, ksaDetail.ProgramSubPlanID)
		if err != nil {
			return err
		}

		return
	}
}

func (u programUsecase) updateOrCreateKsaDetail(tx *gorm.DB, ksaDetails []dto.ProgramKsaDetail, subPlanID uint) (err error) {
	for _, ksaItem := range ksaDetails {
		queryTb := query.ProgramKsaDetailQueryEntity{
			ID: ksaItem.ID,
		}

		updateData := query.ProgramKsaDetailQueryEntity{
			ProgramSubPlanID: &subPlanID,
			KsaType:      ksaItem.KsaType,
			Order:        ksaItem.Order,
			ShortCode:    ksaItem.ShortCode,
			KsaDetail:    ksaItem.KsaDetail,
		}

		if queryTb.ID != nil {
			err = u.CommonRepository.Update(tx, queryTb, &updateData)
			if err != nil {
				return err
			}
		} else {
			err = u.CommonRepository.Create(tx, &updateData)
			if err != nil {
				return err
			}
		}
	}

	return
}

func (u programUsecase) InitKsaDetail(subPlanID uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		initType := []string{
			"k", "s", "a",
		}
		kOrderNum := []uint{
			1, 2, 3, 4, 5, 6,
		}
		kShortCode := []string{
			"K1", "K2", "K3", "K4", "K5", "K6",
		}
		kDetail := []string{
			"Remembering (จำได้)",
			"Understanding (เข้าใจได้)",
			"Applying (ประยุกต์ได้)",
			"Analysis (วิเคราะห์ได้)",
			"Evaluating (ประเมินได้)",
			"Creating (คิดสร้างสรรค์)",
		}

		sOrderNum := []uint{
			1, 2, 3, 4, 5,
		}
		sShortCode := []string{
			"S1", "S2", "S3", "S4", "S5",
		}
		sDetail := []string{
			"Imitating (รับรู้)",
			"Manipulating (ทำตามได้)",
			"Precising (ทำถูกต้อง หาความถูกต้องได้)",
			"Articulating (ทำได้อย่างคล่องแคล่ว)",
			"Naturalizing (ทำได้เป็นธรรมชาติ)",
		}

		aOrderNum := []uint{
			1, 2, 3, 4, 5,
		}
		aShortCode := []string{
			"A1", "A2", "A3", "A4", "A5",
		}
		aDetail := []string{
			"Receiving (รับรู้)",
			"Responding (ตอบสนอง)",
			"Valuating (เห็นคุณค่า)",
			"Organizing (จัดการ)",
			"Characterizing (บุคลิก นิสัย)",
		}

		for i := 0; i < len(kShortCode); i++ {
			initData := query.ProgramKsaDetailQueryEntity{
				ProgramSubPlanID: &subPlanID,
				KsaType:      &initType[0],
				Order:        &kOrderNum[i],
				ShortCode:    &kShortCode[i],
				KsaDetail:    &kDetail[i],
			}

			err := u.CommonRepository.Create(tx, &initData)
			if err != nil {
				return err
			}
		}

		for i := 0; i < len(sShortCode); i++ {
			initData := query.ProgramKsaDetailQueryEntity{
				ProgramSubPlanID: &subPlanID,
				KsaType:      &initType[1],
				Order:        &sOrderNum[i],
				ShortCode:    &sShortCode[i],
				KsaDetail:    &sDetail[i],
			}

			err := u.CommonRepository.Create(tx, &initData)
			if err != nil {
				return err
			}
		}

		for i := 0; i < len(aShortCode); i++ {
			initData := query.ProgramKsaDetailQueryEntity{
				ProgramSubPlanID: &subPlanID,
				KsaType:      &initType[2],
				Order:        &aOrderNum[i],
				ShortCode:    &aShortCode[i],
				KsaDetail:    &aDetail[i],
			}

			err := u.CommonRepository.Create(tx, &initData)
			if err != nil {
				return err
			}
		}

		return
	}
}
