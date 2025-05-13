package usecase

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
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
	commonQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/common"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	"gorm.io/gorm"
)

func (u programUsecase) CreateOrUpdateReference(reference dto.ProgramReferenceDto, programMainUID uuid.UUID, file *multipart.FileHeader, c *fiber.Ctx, userUID uuid.UUID) (err error) {
	return helper.ExecuteTransaction(u.CommonRepository, u.CreateOrUpdateReferenceTransaction(reference, programMainUID, file, c, userUID))
}

func (u programUsecase) CreateOrUpdateReferenceTransaction(reference dto.ProgramReferenceDto, programMainUID uuid.UUID, file *multipart.FileHeader, c *fiber.Ctx, userUID uuid.UUID) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) (err error) {
		queryRefOption := commonQuery.ReferenceOptionQueryEntity{
			ID:   reference.ReferenceTypeID,
			Name: reference.ReferenceTypeName,
		}

		if reference.ReferenceTypeName != nil {
			err = u.CommonRepository.GetFirstOrCreate(&queryRefOption)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}

		queryRef := programQuery.ProgramReferenceQueryEntity{
			ID: reference.ID,
		}

		updateRef := programQuery.ProgramReferenceQueryEntity{
			ProgramID:            &programMainUID,
			ReferenceDescription: reference.ReferenceDescription,
			ReferenceFileName:    reference.ReferenceFileName,
			ReferenceFilePath:    reference.ReferenceFilePath,
			ReferenceName:        reference.ReferenceName,
			ReferenceTypeID:      queryRefOption.ID,
		}

		if file != nil {
			// Clear file before add a new one
			queryDeleteFile := query.MapFilesSystemQueryEntity{
				ReferenceID: queryRef.ID,
			}
			u.CommonRepository.Delete(tx, &queryDeleteFile)

			// Upload file
			initPath := fmt.Sprintf("%v", viper.GetString("file.document_path"))
			uploadPath := fmt.Sprintf("%v/%v/references/reference", initPath, programMainUID)
			if _, err := os.Stat(uploadPath); errors.Is(err, os.ErrNotExist) {
				err := os.MkdirAll(uploadPath, os.ModePerm)
				if err != nil {
					log.Println(err)
					panic(err)
				}
			}

			fileName := fmt.Sprintf("%v_%v", time.Now().Format("20060102150405"), strings.ReplaceAll(file.Filename, " ", "_"))
			destination := fmt.Sprintf("%v/%v", uploadPath, fileName)
			if err := c.SaveFile(file, destination); err != nil {
				return err
			}

			qfType := "program"
			categoryType := "references"
			attribute := "reference"
			createQuery := query.FileSystemQueryEntity{
				QFType:       &qfType,
				QFMainID:     &programMainUID,
				CategoryType: &categoryType,
				Attribute:    &attribute,
				FilePath:     &destination,
				FileName:     &file.Filename,
				CreatedBy:    &userUID,
				UpdatedBy:    &userUID,
			}
			err = u.CommonRepository.Create(tx, &createQuery)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}

			err = u.CommonRepository.Update(tx, queryRef, &updateRef)
			if err != nil {
				if reference.ID == nil || err == gorm.ErrRecordNotFound {
					err = u.CommonRepository.Create(tx, &updateRef)
					if err != nil {
						log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
						return err
					}

					createMapFileSystemQuery := query.MapFilesSystemQueryEntity{
						FileID:      createQuery.ID,
						ReferenceID: updateRef.ID,
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
				FileID:      createQuery.ID,
				ReferenceID: queryRef.ID,
			}
			err = u.CommonRepository.Create(tx, &updateMapFile)
			if err != nil {
				if err != nil {
					log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
					return err
				}
			}

			if queryRefOption.Name == nil {
				err = u.ProgramRepository.UpdateRefNilOption(tx, queryRef.ID)
			}

			return
		}

		err = u.CommonRepository.Update(tx, queryRef, &updateRef)
		if err != nil {
			if reference.ID == nil || err == gorm.ErrRecordNotFound {
				err = u.CommonRepository.Create(tx, &updateRef)
				if err != nil {
					log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
					return err
				}

				if queryRefOption.Name == nil {
					err = u.ProgramRepository.UpdateRefNilOption(tx, updateRef.ID)
					if err != nil {
						log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
						return err
					}
				}
				return nil
			}
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}

		if reference.ReferenceTypeName == nil && queryRef.ID != nil {
			err = u.ProgramRepository.UpdateRefNilOption(tx, queryRef.ID)
			if err != nil {
				log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err)
				return err
			}
		}

		return
	}
}
