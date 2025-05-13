package models

type TbRole struct {
	// RoleName    string `form:"role_name" json:"role_name" gorm:"size:64;index"`
	// Description string `form:"description" json:"description" gorm:"size:64"`
	RoleId uint64 `form:"role_id" json:"role_id" gorm:"primaryKey"`
	// SystemLevel   uint64 `form:"system_level" json:"system_level" gorm:"size:1;default:0"`
	// DivisionLevel uint64 `form:"division_level" json:"division_level" gorm:"size:1;default:0"`
	// ProjectLevel  uint64 `form:"project_level" json:"project_level" gorm:"size:1;default:0"`
	// UserLevel     uint64 `form:"user_level" json:"user_level" gorm:"size:1;default:0"`
}

type UserConfig struct {
	Language string `form:"language" json:"language" gorm:"-"`
	RoleId   uint64 `form:"role_id" json:"role_id" gorm:"-"`
}
