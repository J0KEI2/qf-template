package enums

import (
	"database/sql/driver"

	"github.com/zercle/kku-qf-services/pkg/utils"
)

type UserStatus string

const (
	ACTIVE   UserStatus = "ACTIVE"
	INACTIVE UserStatus = "INACTIVE"
)

func CreateUserStatusEnum() string {
	return utils.CreateEnum("user_status", string(ACTIVE), string(INACTIVE))
}

func (us *UserStatus) Scan(value interface{}) error {
	if us != nil && *us != "" {
	*us = UserStatus(value.([]byte))
	}
	return nil
}

func (us UserStatus) Value() (driver.Value, error) {
	return string(us), nil
}

func (us UserStatus) ToString() (string) {
	return string(us)
}
