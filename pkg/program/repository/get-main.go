package repository

import (
	"fmt"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models/query-model"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
)

// TODO: create the new func don't forget to embed in domain repository interface
func (r programRepository) GetMainProgram(user query.UserQueryEntity, id uuid.UUID) (record *programQuery.ProgramMainQueryEntity, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}

	dbTx := r.MainDbConn

	db := dbTx.Model(programQuery.ProgramMainQueryEntity{})
	db.Where(programQuery.ProgramMainQueryEntity{
		ID: &id,
	})

	db.Preload("ProgramGeneralDetail.Faculty")
	db.Preload("ProgramPermission")

	query := programQuery.ProgramMainQueryEntity{}
	if err = db.First(&query).Error; err != nil {
		return nil, err
	}

	return &query, nil
}
