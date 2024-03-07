package __random

import "testing"

// go test -37_slice FuzzRandom -fuzztime 10s
var fuz map[string]struct{}

func FuzzRandom(f *testing.F) {
	fuz = make(map[string]struct{})
	f.Fuzz(func(t *testing.T, a int) {
		str := RandomStr(a)
		if _, ok := fuz[str]; ok {
			t.Fatal("repeated str")
		}
	})
}
