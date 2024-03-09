package main

import (
	log "github.com/sirupsen/logrus"
)

// RunSafe runs the given fn, recovers if fn panics.
func RunSafe(fn func()) {
	defer Recover()

	fn()
}

// Recover defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		log.Error("panic: ", p)
	}
}

func main() {
	RunSafe(func() {
		panic("panic")
	})
}
