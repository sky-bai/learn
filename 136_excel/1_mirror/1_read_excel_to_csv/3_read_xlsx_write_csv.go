package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

{"system_time":1739168149742,"gps_time":1739168151000,"latitude":34.074385,"longitude":114.839656,"accuracy":15,"gps_bearing":265,"speed":1.13066,"location_type":0,"info":[{"x":163,"y1":165,"x2":234,"y2":211,"score":0.994081,"label":10005010,"id":""}],"version":"6.0.1.abcd"}

// 读取excel 写csv

func main() {
	//ctx := context.Background()

	f, err := excelize.OpenFile("/Users/blj/Downloads/skybai/learn/136_excel/1_mirror/1_read_excel_to_csv/和麦谷重复设备.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 0 // 统计计算数量
	arr := []string{}
	for _, row := range rows {
		// 获取第一列的数据
		i++
		if len(row) == 0 {
			continue
		}
		arr = append(arr, row[1])
		//fmt.Printf(row[1])
	}

	for _, v := range arr {
		temp := v
		filter := bson.M{"imei": temp}
		option := options.FindOne().SetSort(bson.D{{"create_time", -1}}).SetProjection(bson.M{"imei": 1, "tid": 1})
		data, err := model.TidImei.FindOne(filter, option)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				continue
			}

			logger.Logger.Error("ImeiToQQTid model.TidImei.FindOne err:%v", err)
			c.Send(config.REST_CODE_EXCEPTION, "系统异常", nil)
			return
		}
		if data != nil {
			dataList = append(dataList, imeiAndTid{Imei: data.Imei, Tid: data.Tid})
		}
		time.Sleep(100 * time.Millisecond)
	}

	//file, err := os.Create("./orderInfo.csv")
	//if err != nil {
	//	logger.Logger.Errorf("os.create err: %v", err)
	//	return
	//}
	//defer file.Close()
	//writer := csv.NewWriter(file)
	//defer writer.Flush()
	//err = writer.Write([]string{"imei", "task_id", "third_order_id", "create_time", "status", "num_count", "num_succ"})
	//if err != nil {
	//	logger.Logger.Errorf("writer.Write err: %v", err)
	//	return
	//}
	//
	//err = writer.Write([]string{tmpData.Imei, tmpData.TaskId, tmpData.ThirdOrderId, tmpData.CreateTime.String(), strconv.Itoa(tmpData.Status), strconv.Itoa(tmpData.NumCount), strconv.Itoa(tmpData.NumSucc)})
	//if err != nil {
	//	logger.Logger.Errorf("writer.Write err: %v", err)
	//	return
	//}
}
