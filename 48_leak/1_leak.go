package main

type Leak struct {
	ch chan struct{}
}

func NewLeak(ch chan struct{}) *Leak {
	return &Leak{ch: ch}
}

func (l *Leak) leak() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
}
