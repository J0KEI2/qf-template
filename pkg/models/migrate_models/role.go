package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleSystems struct {
	RoleNameTH      string    `gorm:"column:role_name_th;type:varchar;size:200;not null" json:"roleNameTH"`
	RoleNameEN      string    `gorm:"column:role_name_en;type:varchar;size:200;not null" json:"roleNameEN"`
	UID             uuid.UUID `gorm:"column:uid;type:uuid;primaryKey" json:"uid"`
	ViewMyProfile   bool      `gorm:"column:view_my_profile;type:boolean;not null" json:"viewMyProfile"`
	UpdateMyProfile bool      `gorm:"column:update_my_profile;type:boolean;not null" json:"updateMyProfile"`
	ListAllLecture  bool      `gorm:"column:list_all_lecture;type:boolean;not null" json:"listAllLecture"`
	CreateLecture   bool      `gorm:"column:create_lecture;type:boolean;not null" json:"createLecture"`
	UpdateLecture   bool      `gorm:"column:update_lecture;type:boolean;not null" json:"updateLecture"`
	DeleteLecture   bool      `gorm:"column:delete_lecture;type:boolean;not null" json:"deleteLecture"`
	Status          bool      `gorm:"column:status;type:boolean;not null" json:"status"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type RoleCourses struct {
	RoleNameTH         string    `gorm:"column:role_name_th;type:varchar;size:200;not null" json:"roleNameTH"`
	RoleNameEN         string    `gorm:"column:role_name_en;type:varchar;size:200;not null" json:"roleNameEN"`
	UID                uuid.UUID `gorm:"column:uid;type:uuid;primaryKey" json:"uid"`
	CreateCourse       bool      `gorm:"column:create_course;type:boolean;not null" json:"createCourse"`
	ListCourse         bool      `gorm:"column:list_course;type:boolean;not null" json:"listCourse"`
	ViewCourse         bool      `gorm:"column:view_course;type:boolean;not null" json:"viewCourse"`
	ListAllCourse      bool      `gorm:"column:list_all_course;type:boolean;not null" json:"listAllCourse"`
	UpdateCourse       bool      `gorm:"column:update_course;type:boolean;not null" json:"updateCourse"`
	UpdateCourseStatus bool      `gorm:"column:update_course_status;type:boolean;not null" json:"updateCourseStatus"`
	DeleteCourse       bool      `gorm:"column:delete_course;type:boolean;not null" json:"deleteCourse"`
	CommentCourse      bool      `gorm:"column:comment_course;type:boolean;not null" json:"commentCourse"`
	ApproveCourse      bool      `gorm:"column:approve_course;type:boolean;not null" json:"approveCourse"`
	Status             bool      `gorm:"column:status;type:boolean;not null" json:"status"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}

type RolePrograms struct {
	RoleNameTH          string    `gorm:"column:role_name_th;type:varchar;size:200;not null" json:"roleNameTH"`
	RoleNameEN          string    `gorm:"column:role_name_en;type:varchar;size:200;not null" json:"roleNameEN"`
	UID                 uuid.UUID `gorm:"column:uid;type:uuid;primaryKey" json:"uid"`
	CreateProgram       bool      `gorm:"column:create_program;type:boolean;not null" json:"createProgram"`
	ListProgram         bool      `gorm:"column:list_program;type:boolean;not null" json:"listProgram"`
	ViewProgram         bool      `gorm:"column:view_program;type:boolean;not null" json:"viewProgram"`
	ListAllProgram      bool      `gorm:"column:list_all_program;type:boolean;not null" json:"listAllProgram"`
	UpdateProgram       bool      `gorm:"column:update_program;type:boolean;not null" json:"updateProgram"`
	UpdateProgramStatus bool      `gorm:"column:update_program_status;type:boolean;not null" json:"updateProgramStatus"`
	DeleteProgram       bool      `gorm:"column:delete_program;type:boolean;not null" json:"deleteProgram"`
	CommentProgram      bool      `gorm:"column:comment_program;type:boolean;not null" json:"commentProgram"`
	ApproveProgram      bool      `gorm:"column:approve_program;type:boolean;not null" json:"approveProgram"`
	Status              bool      `gorm:"column:status;type:boolean;not null" json:"status"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
}
