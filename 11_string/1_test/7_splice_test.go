package __test

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkAddStringWithOperator(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = hello + "," + world
	}
}
func BenchmarkAddStringWithSprintf(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%s", hello, world)
	}
}

func BenchmarkAddStringWithJoin(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{hello, world}, ",")
	}
}

func BenchmarkAddStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
		_ = buffer.String()
	}
}

func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		//builder.WriteString(str)
		ResetCurrentTimeFeatureFlag("31704750:119972455:1720666423462:-462:1297:1011:9900:215:23:0:1")
	}
	return builder.String()
}

func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, builderConcat) }

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ResetCurrentTimeFeatureFlag(gps string) string {

	gpsData := strings.Split(gps, ":")

	if len(gpsData) >= 11 && gpsData[10] == "3" {
		gpsData[10] = "0"
	} else {
		return gps
	}

	var builder strings.Builder
	for i, val := range gpsData {
		if i > 0 {
			builder.WriteString(":")
		}
		builder.WriteString(val)
	}

	return builder.String()

}

func BenchmarkClearCurrentTimeFeatureFlag(b *testing.B) {
	gpsStr := "31704750:119972455:1720666423462:-462:1297:1011:9900:215:23:0:1"

	for i := 0; i < b.N; i++ {
		ResetCurrentTimeFeatureFlag(gpsStr)
	}
}
