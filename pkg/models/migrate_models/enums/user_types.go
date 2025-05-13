package enums

import (
	"database/sql/driver"

	"github.com/zercle/kku-qf-services/pkg/utils"
)

type UserType string

const (
	INTERNAL UserType = "INTERNAL"
	EXTERNAL UserType = "EXTERNAL"
)

func CreateUserTypesEnum() string {
	return utils.CreateEnum("user_type", string(INTERNAL), string(EXTERNAL))
}

func (ut *UserType) Scan(value interface{}) error {
	if ut != nil && *ut != "" {
		*ut = UserType(value.([]byte))
	}
	return nil
}

func (ut UserType) Value() (driver.Value, error) {
	return string(ut), nil
}

func (ut UserType) ToString() string {
	return string(ut)
}
