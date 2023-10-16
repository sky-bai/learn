// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameWxUser = "wx_user"

// WxUser mapped from table <wx_user>
type WxUser struct {
	ID            int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键Id" json:"id"` // 主键Id
	CreatedAt     *time.Time     `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`          // 创建时间
	UpdatedAt     *time.Time     `gorm:"column:updated_at;type:datetime(3);comment:修改时间" json:"updated_at"`          // 修改时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at"`          // 删除时间
	OpenID        *string        `gorm:"column:open_id;type:varchar(255)" json:"open_id"`
	Customer      *string        `gorm:"column:customer;type:varchar(255)" json:"customer"`
	FollowStatus  *int32         `gorm:"column:follow_status;type:int" json:"follow_status"`
	HeadURL       *string        `gorm:"column:head_url;type:varchar(255)" json:"head_url"`
	LastLoginTime *time.Time     `gorm:"column:last_login_time;type:datetime(3)" json:"last_login_time"`
	Nickname      *string        `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
	Sex           *int32         `gorm:"column:sex;type:int" json:"sex"`
	UnionID       *string        `gorm:"column:union_id;type:varchar(255)" json:"union_id"`
}

// TableName WxUser's table name
func (*WxUser) TableName() string {
	return TableNameWxUser
}