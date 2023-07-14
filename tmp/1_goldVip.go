package tmp

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func main() {

	wids := []string{
		"siweiguangqibentian",
		"yika",
		"turui",
		"xinanzhuo",
		"lingdu",
		"pinyou",
		"yihang",
		"PAPAGO",
		"lichi",
		"renExing",
		"yilu",
		"niuman",
		"EcheEpai",
		"maigu",
		"anzhixiang",
		"lingtu",
		"fuchuang",
		"meijia",
		"siwei",
		"tonggouxinxi",
		"jinanyounitaier",
		"chexingji",
		"haoling",

		"shandongsijian",
		"siweituxin",
	}

	for _, v := range wids {
		Insert(v)
	}

}

func Insert(wid string) {
	createTime := time.Now().UTC()
	updateTime := time.Now().UTC()

	activeStartTime := time.Date(2019, 12, 16, 16, 0, 0, 0, time.UTC)
	activeEndTime := time.Date(2019, 12, 17, 16, 0, 0, 0, time.UTC)
	var space int32 = 0
	var time int32 = 0
	var space_cycle int32 = 0
	var time_cycle int32 = 36
	var price int32 = 19900
	var status int32 = 0
	var ios_price int32 = 26800
	var cost int32 = 15000
	var cloud_time int32 = 0
	var priority int32 = 1
	var active_price int32 = 19900
	var ios_active_price int32 = 26800

	data := bson.M{
		"cloud_time":        cloud_time,
		"product_id":        "VIPGM",
		"title":             "黄金会员",
		"note":              "",
		"space":             space,
		"time":              time,
		"space_cycle":       space_cycle,
		"time_cycle":        time_cycle,
		"status":            status,
		"type":              "member",
		"customer":          "TuYunHuLian",
		"price":             price,
		"ios_price":         ios_price,
		"active_price":      active_price,
		"ios_active_price":  ios_active_price,
		"active_start_time": activeStartTime,
		"active_end_time":   activeEndTime,
		"create_time":       createTime,
		"update_time":       updateTime,
		"priority":          priority,
		"wid":               wid,
		"cost":              cost,
	}
	fmt.Println(data)

	//_, err := model.Package.InsertOne(context.Background(), data)
	//if err != nil {
	//	logger.Error(context.Background(), "Package.InsertOne error:", err)
	//	return
	//}
}
