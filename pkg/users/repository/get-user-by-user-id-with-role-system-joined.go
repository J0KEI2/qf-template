package repository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	helpers "github.com/zercle/gofiber-helpers"
	"github.com/zercle/kku-qf-services/pkg/models"
)

func (r *userRepository) GetUserByIDWithRoleSystem(uid string) (*models.UserFetchWithSystemRoleModel, error) {
	if r.MainDbConn == nil {
		err := fmt.Errorf("%s \nErr: %+v", helpers.WhereAmI(), "database has gone away.")
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	uidConverted, err := uuid.Parse(uid)
	if err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	var userRes models.UserFetchWithSystemRoleModel
	query := `select u.ssn, u.email ,u."type" ,u.status as user_status ,u.uid ,u.faculty_id ,u.system_permission_uid
	,rs.role_name_th ,rs.role_name_en ,rs.view_my_profile ,rs.update_my_profile 
	,rs.list_all_lecture ,rs.create_lecture,rs.create_lecture ,rs.update_lecture  
	,rs.delete_lecture ,rs.status as role_systems_status
	from users u join role_systems rs on rs.uid = u.system_permission_uid where u.uid = ?;
	`
	if err = r.MainDbConn.Raw(query, uidConverted).Scan(&userRes).Error; err != nil {
		log.Printf("source: %+v\nerr: %+v", helpers.WhereAmI(), err.Error())
		return nil, err
	}

	return &userRes, nil
}
