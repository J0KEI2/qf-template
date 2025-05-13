package approvals

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
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/helper"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	"gorm.io/gorm"
)

func (u *programApprovalUsecase) UploadApprovalDocumentFile(c *fiber.Ctx, programUID uuid.UUID, files []*multipart.FileHeader, approvalAttribute string, userUID uuid.UUID, approvalID *uint) error {
	initPath := fmt.Sprintf("%v", viper.GetString("file.document_path"))
	if viper.GetBool("debug_mode") {
		ex, err := os.Executable()
		if err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err)
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)
		fmt.Println("\n exPath: ", exPath)
		fmt.Println("\n uploadPath: ", initPath)
		fmt.Println("\n programUID: ", programUID)
	}

	uploadPath := fmt.Sprintf("%v/%v/approval/%s", initPath, programUID, approvalAttribute)
	if _, err := os.Stat(uploadPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err)
			panic(err)
		}
	}

	filePathList := make([]map[string]string, 0)
	for _, file := range files {
		fileName := fmt.Sprintf("%v_%v", time.Now().Format("20060102150405"), strings.ReplaceAll(file.Filename, " ", "_"))
		destination := fmt.Sprintf("%v/%v", uploadPath, fileName)
		if err := c.SaveFile(file, destination); err != nil {
			log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err)
			return err
		}

		fileData := map[string]string{
			"file_name": file.Filename,
			"file_path": destination,
		}

		filePathList = append(filePathList, fileData)
	}

	if err := helper.ExecuteTransaction(u.CommonRepository, u.CreateApprovalFilesSystem(filePathList, approvalAttribute, programUID, userUID, approvalID)); err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err)
		return err
	}

	return nil
}

func (u *programApprovalUsecase) CreateApprovalFilesSystem(filePathList []map[string]string, approvalAttribute string, programMainUID, createdBy uuid.UUID, approvalID *uint) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		for _, file := range filePathList {
			qfType := "program"
			categoryType := "approval"
			filePath := file["file_path"]
			fileName := file["file_name"]

			createQuery := query.FileSystemQueryEntity{
				QFType:       &qfType,
				QFMainID:     &programMainUID,
				CategoryType: &categoryType,
				Attribute:    &approvalAttribute,
				FilePath:     &filePath,
				FileName:     &fileName,
				CreatedBy:    &createdBy,
				UpdatedBy:    &createdBy,
			}

			if err = u.CommonRepository.Create(tx, &createQuery); err != nil {
				log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err)
				return err
			}

			mapFileQuery := query.MapFilesSystemQueryEntity{
				FileID:     createQuery.ID,
				ApprovalID: approvalID,
			}

			if err = u.CommonRepository.Create(tx, &mapFileQuery); err != nil {
				log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err)
				return err
			}
		}
		return
	}
}
