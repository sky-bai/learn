// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRoleAPI = "role_api"

// RoleAPI mapped from table <role_api>
type RoleAPI struct {
	ID        int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键Id" json:"id"`                                 // 主键Id
	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                                          // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime(3);comment:修改时间" json:"updated_at"`                                          // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);index:idx_role_api_deleted_at,priority:1;comment:删除时间" json:"deleted_at"` // 删除时间
	Role      *string        `gorm:"column:role;type:varchar(191);index:idx_role_api_role,priority:1;comment:角色" json:"role"`                    // 角色
	URI       *string        `gorm:"column:uri;type:varchar(191);index:idx_role_api_uri,priority:1;comment:允许的请求uri" json:"uri"`                 // 允许的请求uri
	Method    *string        `gorm:"column:method;type:longtext;comment:允许的请求方法" json:"method"`                                                  // 允许的请求方法
}

// TableName RoleAPI's table name
func (*RoleAPI) TableName() string {
	return TableNameRoleAPI
}
