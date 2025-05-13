package usecase

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u programUsecase) UploadReferenceDocumentsFile(c *fiber.Ctx, programUID uuid.UUID, files []*multipart.FileHeader, userUID uuid.UUID) error {
	initPath := fmt.Sprintf("%v", viper.GetString("file.document_path"))
	if viper.GetBool("debug_mode") {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)
		fmt.Println("\n exPath: ", exPath)
		fmt.Println("\n uploadPath: ", initPath)
		fmt.Println("\n programUID: ", programUID)
	}

	uploadPath := fmt.Sprintf("%v/%v/references/reference", initPath, programUID)
	if _, err := os.Stat(uploadPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}

	filePathList := make([]map[string]string, 0)
	for _, file := range files {
		fileName := fmt.Sprintf("%v_%v", time.Now().Format("20060102150405"), strings.ReplaceAll(file.Filename, " ", "_"))
		destination := fmt.Sprintf("%v/%v", uploadPath, fileName)
		if err := c.SaveFile(file, destination); err != nil {
			return err
		}

		fileData := map[string]string{
			"file_name": file.Filename,
			"file_path": destination,
		}

		filePathList = append(filePathList, fileData)
	}

	err := helper.ExecuteTransaction(u.CommonRepository, u.CreateReferenceFilesSystem(filePathList, programUID, userUID))
	if err != nil {
		return err
	}

	return nil
}

func (u programUsecase) CreateReferenceFilesSystem(filePathList []map[string]string, programMainUID, userUID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, file := range filePathList {
			qfType := "program"
			categoryType := "references"
			attribute := "reference"
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

			// programIdStr := programMainUID.String()
			// mapFileSystemQuery := query.MapFilesSystemQueryEntity{
			// 	FileID:      createQuery.ID,
			// 	Attribute:   &attribute,
			// 	AttributeID: &programIdStr,
			// }

			// err = u.CommonRepository.Create(tx, &mapFileSystemQuery)
			// if err != nil {
			// 	log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
			// 	return err
			// }
		}
		return
	}
}
