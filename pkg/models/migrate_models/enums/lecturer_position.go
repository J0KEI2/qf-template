package enums

import (
	"database/sql/driver"

	"github.com/zercle/kku-qf-services/pkg/utils"
)

type LecturerPosition string

const (
	PROFESSOR           LecturerPosition = "ศาสตราจารย์"
	ASSOCIATE_PROFESSOR LecturerPosition = "รองศาสตราจารย์"
	ASSISTANT_PROFESSOR LecturerPosition = "ผู้ช่วยศาสตราจารย์"
	LECTURER            LecturerPosition = "อาจารย์"
	STAFF               LecturerPosition = "เจ้าหน้าที่"
	// PROFESSORTH           LecturerPosition = "ศาสตราจารย์"
	// ASSOCIATE_PROFESSORTH LecturerPosition = "ASSOCIATE_PROFESSOR"
	// ASSISTANT_PROFESSORTH LecturerPosition = "ASSISTANT_PROFESSOR"
)

func CreateLecturerPositionEnum() string {
	return utils.CreateEnum("lecturer_position", string(PROFESSOR), string(ASSOCIATE_PROFESSOR), string(ASSISTANT_PROFESSOR), string(LECTURER), string(STAFF))
}

func (lp *LecturerPosition) Scan(value interface{}) error {
	*lp = LecturerPosition(value.([]byte))
	return nil
}

func (lp LecturerPosition) Value() (driver.Value, error) {
	return string(lp), nil
}

func (lp LecturerPosition) ToString() string {
	return string(LecturerPosition(lp))
}