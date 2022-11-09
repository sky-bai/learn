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
	fmt.Println(res)

	// 查找两个坐标的距离
	resDist, _ := rdb.GeoDist(ctx8, "geo_hash_test", "天府广场", "宽窄巷子", "m").Result()
	fmt.Println(resDist)

	// 查找该对象周围的对象
	//rdb.GeoRadiusByMember()
}
