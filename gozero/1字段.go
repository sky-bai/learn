package main

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"

	"strings"
	"time"
)

var (
	gatewayServiceInfoFieldNames          = builderx.RawFieldNames(&GatewayServiceInfo{})
	gatewayServiceInfoRows                = strings.Join(gatewayServiceInfoFieldNames, ",")
	gatewayServiceInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(gatewayServiceInfoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	gatewayServiceInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(gatewayServiceInfoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type GatewayServiceInfo struct {
	Id          int64     `db:"id"`           // 自增主键
	LoadType    int64     `db:"load_type"`    // 负载类型 0=http 1=tcp 2=grpc
	ServiceName string    `db:"service_name"` // 服务名称 6-128 数字字母下划线
	ServiceDesc string    `db:"service_desc"` // 服务描述
	CreateAt    time.Time `db:"create_at"`    // 添加时间
	UpdateAt    time.Time `db:"update_at"`    // 更新时间
	IsDelete    int64     `db:"is_delete"`    // 是否删除 1=删除
}

func main() {
	fmt.Println(gatewayServiceInfoRowsExpectAutoSet)
}
