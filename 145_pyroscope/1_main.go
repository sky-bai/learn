package main

import (
	"fmt"
	"github.com/grafana/pyroscope-go"

	"os"
	"runtime"
)

func main() {
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	host, _ := os.Hostname()

	pyroscope.Start(pyroscope.Config{
		ApplicationName: "cf-push",

		// replace this with the address of pyroscope server
		ServerAddress: "http://cf-pyroscope.cut-frame:9095",

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,
		// you can provide static tags via a map:
		Tags: map[string]string{"hostname": host},

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	degt()

}

func degt() {
	for i := 0; i < 10000000; i++ {
		//time.Sleep(time.Second * time.Duration(i))
		fmt.Println(i)
	}
}

func startPyroScope() {
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	host, _ := os.Hostname()

	pyroscope.Start(pyroscope.Config{
		ApplicationName: "cf-push",

		// replace this with the address of pyroscope server
		ServerAddress: "http://cf-pyroscope.cut-frame:9095",

		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,
		// you can provide static tags via a map:
		Tags: map[string]string{"hostname": host},

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
}
