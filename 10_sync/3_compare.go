package main

import (
	"fmt"
	"sync"
	"time"
)

var s sync.RWMutex
var w sync.WaitGroup

func main() {
	mapTest()
	syncMapTest()
}
func mapTest() {
	m := map[int]int{1: 1}
	startTime := time.Now().Nanosecond()
	w.Add(1)
	go writeMap(m)
	w.Add(1)
	go writeMap(m)
	//w.Add(1)
	//go readMap(m)

	w.Wait()
	endTime := time.Now().Nanosecond()
	timeDiff := endTime - startTime
	seconds := float64(timeDiff) / 1e9
	fmt.Println("map:", seconds)
}

func writeMap(m map[int]int) {
	defer w.Done()
	i := 0
	for i < 500000 {
		// 加锁
		s.Lock()
		m[1] = 1
		// 解锁
		s.Unlock()
		i++
	}
}

func readMap(m map[int]int) {
	defer w.Done()
	i := 0
	for i < 10000 {
		s.RLock()
		_ = m[1]
		s.RUnlock()
		i++
	}
}

func syncMapTest() {
	m := sync.Map{}
	m.Store(1, 1)
	startTime := time.Now().Nanosecond()
	w.Add(1)
	go writeSyncMap(m)
	w.Add(1)
	go writeSyncMap(m)
	//w.Add(1)
	//go readSyncMap(m)

	w.Wait()
	endTime := time.Now().Nanosecond()
	timeDiff := endTime - startTime
	// timeDiff 转成秒
	seconds := float64(timeDiff) / 1e9
	fmt.Println("sync.Map:", seconds)
}

func writeSyncMap(m sync.Map) {
	defer w.Done()
	i := 0
	for i < 500000 {
		m.Store(1, 1)
		i++
	}
}

func readSyncMap(m sync.Map) {
	defer w.Done()
	i := 0
	for i < 10000 {
		m.Load(1)
		i++
	}
}
