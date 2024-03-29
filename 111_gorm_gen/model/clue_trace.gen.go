// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameClueTrace = "clue_trace"

// ClueTrace mapped from table <clue_trace>
type ClueTrace struct {
	ID           int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键Id" json:"id"`          // 主键Id
	CreatedAt    *time.Time     `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                   // 创建时间
	UpdatedAt    *time.Time     `gorm:"column:updated_at;type:datetime(3);comment:修改时间" json:"updated_at"`                   // 修改时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at"`                   // 删除时间
	Status       int32          `gorm:"column:status;type:int;not null;comment:线索跟进状态" json:"status"`                        // 线索跟进状态
	TraceContent *string        `gorm:"column:trace_content;type:varchar(255);comment:跟进内容" json:"trace_content"`            // 跟进内容
	ClueID       *int64         `gorm:"column:clue_id;type:bigint;index:idx_clue_id,priority:1;comment:线索ID" json:"clue_id"` // 线索ID
	TraceType    *string        `gorm:"column:trace_type;type:varchar(255);comment:跟进类型，查阅车主信息、跟进记录" json:"trace_type"`      // 跟进类型，查阅车主信息、跟进记录
	AccountID    int64          `gorm:"column:account_id;type:bigint;not null;comment:账号ID" json:"account_id"`               // 账号ID
	EnterpriseID *string        `gorm:"column:enterprise_id;type:longtext;comment:集团id,主系统wid" json:"enterprise_id"`         // 集团id,主系统wid
	SubsidiaryID *int64         `gorm:"column:subsidiary_id;type:bigint;comment:所属4S店" json:"subsidiary_id"`                 // 所属4S店
}

// TableName ClueTrace's table name
func (*ClueTrace) TableName() string {
	return TableNameClueTrace
}
