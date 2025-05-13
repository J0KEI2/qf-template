package usecase

import (
	"log"

	helpers "github.com/zercle/gofiber-helpers"
	constant "github.com/zercle/kku-qf-services/pkg/constant/common"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
)

func (u commonUsecase) GetFilePathByID(fileID uint) (filePath *string, err error) {
	query := query.FileSystemQueryEntity{
		ID: &fileID,
	}

	if err := u.CommonRepository.GetFirst(&query); err != nil {
		log.Printf(constant.WHERE_AM_I_LOG, helpers.WhereAmI(), err.Error())
		return nil, err
	}

	return query.FilePath, nil
}
