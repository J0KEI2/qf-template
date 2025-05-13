package repository

import (
	"fmt"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/common/repository"
	"github.com/zercle/kku-qf-services/pkg/models"
	"github.com/zercle/kku-qf-services/pkg/models/dto"
	query "github.com/zercle/kku-qf-services/pkg/models/query-model"
	programQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/program"
	rapQuery "github.com/zercle/kku-qf-services/pkg/models/query-model/role-and-permission"
)

func (r programRepository) GetProgramMainPagination(user *query.UserQueryEntity, role *rapQuery.RoleQueryEntity, paginationOptions *models.PaginationOptions, param dto.GetMainProgramPaginationQueryParam) (record []programQuery.ProgramMainQueryEntity, err error) {
	if r.MainDbConn == nil {
		err = fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		return
	}
	dbTx := r.MainDbConn

	record = make([]programQuery.ProgramMainQueryEntity, 0)

	accessLevel := pointer.GetUint(role.ProgramAccessLevel)

	if accessLevel == 0 {
		return record, nil
	}
	db := dbTx.Model(programQuery.ProgramMainQueryEntity{})
	db.Joins("JOIN program_general_detail as qgd ON program_main.program_general_detail_id = qgd.id")
	db.Joins("JOIN common_faculties as fct ON qgd.faculty_id = fct.id")

	switch accessLevel {
	case 3:
		{
			break
		}
	case 2:
		{
			facultyIdList := make([]uint, 0)
			for _, mapFaculty := range role.MapFacultiesRoles {
				facultyIdList = append(facultyIdList, pointer.GetUint(mapFaculty.FacultyID))
			}
			db.Where("qgd.faculty_id in (?)", facultyIdList)
		}
	case 1:
		{
			programIdList := make([]uuid.UUID, 0)
			for _, mapProgram := range role.MapProgramRoles {
				programId := mapProgram.ProgramID
				programIdList = append(programIdList, *programId)
			}
			db.Where("program_main.id in (?)", programIdList)
		}
	default:
		{
			return record, nil
		}
	}

	whereQueryString := []string{}
	whereQueryParam := []interface{}{}
	if param.FacultyID != nil {
		whereQueryString = append(whereQueryString, "qgd.faculty_id = ?")
		whereQueryParam = append(whereQueryParam, param.FacultyID)
	}

	if param.FacultyName != nil {
		search := "%" + pointer.GetString(param.FacultyName) + "%"
		whereQueryString = append(whereQueryString, "fct.faculty_name_th LIKE ? OR fct.faculty_name_en LIKE ?")
		whereQueryParam = append(whereQueryParam, search)
		whereQueryParam = append(whereQueryParam, search)
	}

	if param.ProgramCode != nil {
		search := "%" + pointer.GetString(param.ProgramCode) + "%"
		whereQueryString = append(whereQueryString, "qgd.program_code LIKE ?")
		whereQueryParam = append(whereQueryParam, search)
	}

	if param.ProgramId != nil {
		search := "%" + fmt.Sprintf("%v", *param.ProgramId) + "%"
		whereQueryString = append(whereQueryString, "program_main.id::text LIKE ?")
		whereQueryParam = append(whereQueryParam, search)
	}

	if param.ProgramName != nil {
		search := "%" + pointer.GetString(param.ProgramName) + "%"
		whereQueryString = append(whereQueryString, "qgd.program_name_th LIKE ? OR qgd.program_name_en LIKE ?")
		whereQueryParam = append(whereQueryParam, search)
		whereQueryParam = append(whereQueryParam, search)
	}

	db.Where(strings.Join(whereQueryString, " AND "), whereQueryParam...)

	if search := pointer.GetString(paginationOptions.Search); search != "" {
		search = "%" + search + "%"
		db.Where("(qgd.program_name_th LIKE ? OR qgd.program_name_en LIKE ?)", search, search)
	}

	db.Preload("ProgramGeneralDetail.Faculty")

	db.Count(paginationOptions.Total)
	db.Scopes(repository.Paginate(paginationOptions))

	if err = db.Find(&record).Error; err != nil {
		return nil, err
	}
	return
}
