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
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u programUsecase) UploadGeneralDetailMouDocumentsFile(c *fiber.Ctx, programUID uuid.UUID, generalDetailId *uint, files []dto.MouFileDto, userUID uuid.UUID) ([]uint, error) {
	// init data
	initPath := fmt.Sprintf("%v", viper.GetString("file.document_path"))

	uploadPath := fmt.Sprintf("%v/%v/general-detail/mou", initPath, programUID)
	if _, err := os.Stat(uploadPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}

	// Get file list from generalID
	queryMapFile := query.MapFilesSystemQueryEntity{
		GeneralDetailID: generalDetailId,
	}

	mapFileArray := make([]query.MapFilesSystemQueryEntity, 0)
	err := u.CommonRepository.GetList(&queryMapFile, &mapFileArray, nil, "FileSystem")
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
		return nil, err
	}

	deleteFileIDList := make([]uint, 0)
newLoop:
	for _, mapFileData := range mapFileArray {
		// check input
		for _, mouFileData := range files {
			if mouFileData.FileID != nil {
				if *mapFileData.FileID == *mouFileData.FileID {
					continue newLoop
				}
			}
		}
		deleteFileIDList = append(deleteFileIDList, *mapFileData.FileID)
	}

	// delete file
	helper.ExecuteTransaction(u.CommonRepository, u.DeleteMapFiles(deleteFileIDList))

	filePathList := make([]models.FileSystemData, 0)
	for _, file := range files {
		if file.File != nil {
			fileName := fmt.Sprintf("%v_%v", time.Now().Format("20060102150405"), strings.ReplaceAll(file.File.Filename, " ", "_"))
			destination := fmt.Sprintf("%v/%v", uploadPath, fileName)
			if err := c.SaveFile(file.File, destination); err != nil {
				return nil, err
			}

			fileData := models.FileSystemData{
				FilePath: &destination,
				FileName: &file.File.Filename,
				FileID:   file.FileID,
			}

			filePathList = append(filePathList, fileData)
		}

	}

	err = helper.ExecuteTransaction(u.CommonRepository, u.CreateGeneralDetailFilesSystem(filePathList, programUID, userUID, generalDetailId))
	if err != nil {
		return nil, err
	}

	return deleteFileIDList, nil
}

func (u programUsecase) CreateGeneralDetailFilesSystem(filePathList []models.FileSystemData, programMainUID, userUID uuid.UUID, generalDetailId *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, file := range filePathList {
			qfType := "program"
			categoryType := "general_detail"
			attribute := "mou"
			filePath := file.FilePath
			fileName := file.FileName

			createQuery := query.FileSystemQueryEntity{
				QFType:       &qfType,
				QFMainID:     &programMainUID,
				CategoryType: &categoryType,
				Attribute:    &attribute,
				FilePath:     filePath,
				FileName:     fileName,
				CreatedBy:    &userUID,
				UpdatedBy:    &userUID,
			}
			err = u.CommonRepository.Create(tx, &createQuery)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}

			if file.FileID != nil {
				mapFileSystemQuery := query.MapFilesSystemQueryEntity{
					FileID: file.FileID,
				}
				mapFileSystemUpdateQuery := query.MapFilesSystemQueryEntity{
					GeneralDetailID: generalDetailId,
					FileID:          createQuery.ID,
				}

				err = u.CommonRepository.Update(tx, &mapFileSystemQuery, &mapFileSystemUpdateQuery)
				if err != nil {
					log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
					return err
				}
				continue
			}

			mapFileSystemQuery := query.MapFilesSystemQueryEntity{
				FileID:          createQuery.ID,
				GeneralDetailID: generalDetailId,
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

func (u programUsecase) DeleteMapFiles(fileIDs []uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, fileId := range fileIDs {
			mapFileQuery := query.MapFilesSystemQueryEntity{
				FileID: &fileId,
			}
			err = u.CommonRepository.Delete(tx, &mapFileQuery)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}
		return
	}
}

func (u programUsecase) DeleteFiles(fileIDs []uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, fileId := range fileIDs {
			mapFileQuery := query.FileSystemQueryEntity{
				ID: &fileId,
			}
			err = u.CommonRepository.Delete(tx, &mapFileQuery)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}
		return
	}
}
