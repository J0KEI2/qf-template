package query

import "time"

type ActionLog struct {
	ID        *uint      `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID    *string    `gorm:"column:user_id;type:varchar"`
	Method    *string    `gorm:"column:method;type:varchar;size:255"`
	Action    *string    `gorm:"column:action;type:varchar;size:255"`
	Payload   *string    `gorm:"column:payload;type:text"`
	Params    *string    `gorm:"column:params;type:text"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (s *ActionLog) TableName() string {
	return "action_logs"
}
