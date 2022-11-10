package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func (c *Config) T() {

}

func BenchmarkAtomic(b *testing.B) {
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i++
			cfg := &Config{a: []int{i, i + 1, i + 2}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < b.N; i++ {
				cfg := v.Load().(*Config)
				cfg.T()
				fmt.Println(cfg)
			}
			wg.Done()
		}()

	}

}
