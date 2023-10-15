// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTest = "test"

// Test mapped from table <test>
type Test struct {
	ID         int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键Id" json:"id"` // 主键Id
	CreatedAt  *time.Time     `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`          // 创建时间
	UpdatedAt  *time.Time     `gorm:"column:updated_at;type:datetime(3);comment:修改时间" json:"updated_at"`          // 修改时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at"`          // 删除时间
	AccountID  *int64         `gorm:"column:account_id;type:bigint;comment:审核人ID" json:"account_id"`              // 审核人ID
	Result     *int32         `gorm:"column:result;type:int;comment:人工审核结果 1是碰撞 2不是碰撞" json:"result"`             // 人工审核结果 1是碰撞 2不是碰撞
	State      *int32         `gorm:"column:state;type:int;comment:审核状态 1待审核 2审核完成" json:"state"`                 // 审核状态 1待审核 2审核完成
	TaskID     string         `gorm:"column:task_id;type:varchar(255);not null;comment:任务Id" json:"task_id"`      // 任务Id
	Imei       *string        `gorm:"column:imei;type:varchar(255);comment:设备imei" json:"imei"`                   // 设备imei
	CrashLevel *bool          `gorm:"column:crash_level;type:tinyint(1);comment:碰撞等级" json:"crash_level"`         // 碰撞等级
	Ts         *int64         `gorm:"column:ts;type:bigint;comment:时间戳" json:"ts"`                                // 时间戳
	FileType   *string        `gorm:"column:fileType;type:varchar(255);comment:文件类型" json:"fileType"`             // 文件类型
	Camera     *bool          `gorm:"column:camera;type:tinyint(1);comment:摄像头" json:"camera"`                    // 摄像头
	URL        *string        `gorm:"column:url;type:varchar(255);comment:文件链接" json:"url"`                       // 文件链接
	Longitude  *float64       `gorm:"column:longitude;type:decimal(10,6);comment:事故发生时的经度" json:"longitude"`      // 事故发生时的经度
	Latitude   *float64       `gorm:"column:latitude;type:decimal(10,6);comment:事故发生时的纬度" json:"latitude"`        // 事故发生时的纬度
}

// TableName Test's table name
func (*Test) TableName() string {
	return TableNameTest
}
