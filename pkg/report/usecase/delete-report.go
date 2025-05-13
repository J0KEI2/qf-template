package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u reportUsecase) DeleteReport(id uint) (err error) {

	query := query.ReportQueryEntity{
		ID: &id,
	}

	return helper.ExecuteTransaction(u.CommonRepository, u.DeleteReportTransaction(&query))
}

func (u reportUsecase) DeleteReportTransaction(reportQuery *query.ReportQueryEntity) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		// fileQuery := query.MapFilesSystemQueryEntity{
		// 	ReportID: reportQuery.ID,
		// }
		// if err := u.CommonRepository.GetFirst(&fileQuery, "FileSystem"); err != nil {
		// 	log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		// 	return err
		// }

		queryMapFile := query.MapFilesSystemQueryEntity{
			ReportID: reportQuery.ID,
		}
		if err := u.CommonRepository.Delete(tx, &queryMapFile); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
			return err
		}

		// queryFile := query.FileSystemQueryEntity{
		// 	ID: fileQuery.FileID,
		// }
		// if err := u.CommonRepository.Delete(tx, &queryFile); err != nil {
		// 	log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		// 	return err
		// }

		if err := u.CommonRepository.Delete(tx, reportQuery); err != nil {
			log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
			return err
		}

		return nil
	}
}
