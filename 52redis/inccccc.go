package incr

import (
	"context"
	"cutframe/config"
	"cutframe/lib/goredis"
	"cutframe/lib/logger"
	"cutframe/lib/tools"
	"sync/atomic"
	"time"
)

var Incr *incr

func init() {
	channels := []string{"maigu", "yika", "wanji", "baiduMap", "qqMap", "lingdu"}
	// 获取到所有需要做统计的key
	var keys []string

	for _, v := range channels {
		temp := v
		keys = append(keys, config.REDIS_PREFIX_cfStat+"track_flow_"+temp+tools.Date("Ymd", time.Now().Unix()))

	}

	Incr = NewIncr(keys)

}

// incr 是一个简单的累加器
type incr struct {
	keys []string

	keyAndCount map[string]int64

	maiGuTrackNum  int64
	maiGuTrackFlow int64
}

// NewIncr 创建一个新的 Incr 实例
func NewIncr(keys []string) *incr {
	return &incr{
		keys: keys,
	}
}

// IncrMaiGuTrackNum 对给定的 key 进行累加操作
func (i *incr) IncrMaiGuTrackNum(val int64) {
	atomic.AddInt64(&i.maiGuTrackNum, val) // 这个做累加 但是实际写入redis是根据key来判断 每一个统计需要一个atomic值
}

// TakeoutToRedis 获取给定 key 的当前累加值并重置为0
func (i *incr) TakeoutToRedis() {
	// 1.获取值
	value := atomic.SwapInt64(&i.maiGuTrackFlow, 0)

	// 2.更新redis 降低redis qps
	err := goredis.Backend.IncrBy(context.Background(), config.REDIS_PREFIX_cfStat+"track_flow_maigu_"+tools.Date("Ymd", time.Now().Unix()), value).Err()
	if err != nil {
		logger.Logger.Errorf("TakeoutToRedis err:%v", err)
	}

	// 3. ...

}

func (i *incr) Stop() {

}
