package usecase

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u reportUsecase) CreateOrUpdateReport(c *fiber.Ctx, request dto.CreateOrUpdateReportRequestDto, userUID uuid.UUID) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateReportTransaction(c, request, userUID))
}

func (u reportUsecase) CreateOrUpdateReportTransaction(c *fiber.Ctx, request dto.CreateOrUpdateReportRequestDto, userUID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		queryReport := query.ReportQueryEntity{
			ID: request.ReportID,
		}
		updateReport := query.ReportQueryEntity{
			QFMainID:    request.ProgramUID,
			Name:        request.Name,
			Description: request.Description,
		}

		if request.File != nil {
			// Clear file before add a new one
			queryDeleteFile := query.MapFilesSystemQueryEntity{
				ReportID: request.ReportID,
			}
			u.CommonRepository.Delete(tx, &queryDeleteFile)

			// Upload file
			initPath := fmt.Sprintf("%v", viper.GetString("file.document_path"))
			uploadPath := fmt.Sprintf("%v/%v/references/reference", initPath, updateReport.QFMainID)
			if _, err := os.Stat(uploadPath); errors.Is(err, os.ErrNotExist) {
				err := os.MkdirAll(uploadPath, os.ModePerm)
				if err != nil {
					log.Println(err)
					panic(err)
				}
			}

			fileName := fmt.Sprintf("%v_%v", time.Now().Format("20060102150405"), strings.ReplaceAll(request.File.Filename, " ", "_"))
			destination := fmt.Sprintf("%v/%v", uploadPath, fileName)
			if err := c.SaveFile(request.File, destination); err != nil {
				return err
			}

			qfType := "program"
			categoryType := "documents"
			attribute := "report"
			createQuery := query.FileSystemQueryEntity{
				QFType:       &qfType,
				QFMainID:     updateReport.QFMainID,
				CategoryType: &categoryType,
				Attribute:    &attribute,
				FilePath:     &destination,
				FileName:     &request.File.Filename,
				CreatedBy:    &userUID,
				UpdatedBy:    &userUID,
			}
			err = u.CommonRepository.Create(tx, &createQuery)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}

			err = u.CommonRepository.Update(tx, &queryReport, &updateReport)
			if err != nil {
				if queryReport.ID == nil || err == gorm.ErrRecordNotFound {
					err = u.CommonRepository.Create(tx, &updateReport)
					if err != nil {
						log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
						return err
					}

					createMapFileSystemQuery := query.MapFilesSystemQueryEntity{
						FileID:   createQuery.ID,
						ReportID: updateReport.ID,
					}
					err = u.CommonRepository.Create(tx, &createMapFileSystemQuery)
					if err != nil {
						log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
						return err
					}
					return nil
				}
				if err != nil {
					log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
					return err
				}
			}

			updateMapFile := query.MapFilesSystemQueryEntity{
				FileID:   createQuery.ID,
				ReportID: queryReport.ID,
			}
			err = u.CommonRepository.Create(tx, &updateMapFile)
			if err != nil {
				if err != nil {
					log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
					return err
				}
			}

			return
		}

		err = u.CommonRepository.Update(tx, &queryReport, &updateReport)
		if err != nil {
			if queryReport.ID == nil || err == gorm.ErrRecordNotFound {
				err = u.CommonRepository.Create(tx, &updateReport)
			}
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}

		return
	}

	// initPath := fmt.Sprintf("%v", viper.GetString("file.document_path"))

	// uploadPath := fmt.Sprintf("%v/%v/documents/report", initPath, *request.ProgramUID)
	// if _, err := os.Stat(uploadPath); errors.Is(err, os.ErrNotExist) {
	// 	err := os.MkdirAll(uploadPath, os.ModePerm)
	// 	if err != nil {
	// 		log.Println(err)
	// 		panic(err)
	// 	}
	// }

	// filePathList := make([]map[string]string, 0)
	// for _, file := range request.Files {
	// 	fileName := fmt.Sprintf("%v_%v", time.Now().Format("20060102150405"), strings.ReplaceAll(file.Filename, " ", "_"))
	// 	destination := fmt.Sprintf("%v/%v", uploadPath, fileName)
	// 	if err := c.SaveFile(file, destination); err != nil {
	// 		return err
	// 	}

	// 	fileData := map[string]string{
	// 		"file_name": fileName,
	// 		"file_path": destination,
	// 	}

	// 	filePathList = append(filePathList, fileData)
	// }

	// err := helper.ExecuteTransaction(u.CommonRepository, u.CreateReportFilesSystem(filePathList, *request.ProgramUID, userUID, request.ReportID))
	// if err != nil {
	// 	return err
	// }

	// return nil
}

func (u reportUsecase) CreateReportFilesSystem(filePathList []map[string]string, programMainUID, userUID uuid.UUID, reportID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, file := range filePathList {
			qfType := "program"
			categoryType := "documents"
			attribute := "report"
			filePath := file["file_path"]
			fileName := file["file_name"]

			createQuery := query.FileSystemQueryEntity{
				QFType:       &qfType,
				QFMainID:     &programMainUID,
				CategoryType: &categoryType,
				Attribute:    &attribute,
				FilePath:     &filePath,
				FileName:     &fileName,
				CreatedBy:    &userUID,
				UpdatedBy:    &userUID,
			}
			err = u.CommonRepository.Create(tx, &createQuery)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}

			mapFileSystemQuery := query.MapFilesSystemQueryEntity{
				FileID:   createQuery.ID,
				ReportID: reportID,
			}

			err = u.CommonRepository.Create(tx, &mapFileSystemQuery)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}
		return
	}
}
