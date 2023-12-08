package incr

import (
	"testing"
)

var ins = NewSimpleIncrementer()

func BenchmarkSimpleIncrementer_Incr(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ins.Incr("a", 5)
	}
}
