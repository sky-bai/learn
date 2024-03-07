package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx8 := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer rdb.Close()
	// 只能说在某一个特定距离下，就显示没有

	//geo(rdb, ctx8)

	//now := time.Now()
	//// 查找该对象距离最近的对象
	//result, err := rdb.GeoRadiusByMember(ctx8, "geo_hash_test", "天府广场", &redis.GeoRadiusQuery{
	//	Radius: 100, // 范围越大耗时越大
	//	Unit:   "km",
	//	//WithCoord:   true,
	//	WithDist:    true,
	//	WithGeoHash: true,
	//	Count:       6,
	//}).Result()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("haoShi", time.Since(now))
	//fmt.Println("dis", result)

}

func geo(rdb *redis.Client, ctx8 context.Context) {

	res, _ := rdb.GeoAdd(ctx8, "geo_hash_test", &redis.GeoLocation{
		Name:      "天府广场",
		Longitude: 104.072833,
		Latitude:  30.663422,
	}, &redis.GeoLocation{
		Name:      "四川大剧院",
		Longitude: 104.074378,
		Latitude:  30.664804,
	}, &redis.GeoLocation{
		Name:      "新华文轩",
		Longitude: 104.070084,
		Latitude:  30.664649,
	}, &redis.GeoLocation{
		Name:      "手工茶",
		Longitude: 104.072402,
		Latitude:  30.664121,
	}, &redis.GeoLocation{
		Name:      "宽窄巷子",
		Longitude: 104.059826,
		Latitude:  30.669883,
	}, &redis.GeoLocation{
		Name:      "奶茶",
		Longitude: 104.06085,
		Latitude:  30.670054,
	}, &redis.GeoLocation{
		Name:      "钓鱼台",
		Longitude: 104.058424,
		Latitude:  30.670737,
	}).Result()

	for i := 0; i < 3000000; i++ {

		lon := 104.072833 + float64(i)/1000000

		_, _ = rdb.GeoAdd(ctx8, "geo_hash_test", &redis.GeoLocation{
			Name:      "天府广场" + fmt.Sprintf("%d", i),
			Longitude: lon,
			Latitude:  30.663422,
		}).Result()
	}

	fmt.Println(res)

	// 查找两个坐标的距离
	resDist, _ := rdb.GeoDist(ctx8, "geo_hash_test", "天府广场", "宽窄巷子", "m").Result()
	fmt.Println(resDist)

}
