package _5_fuzz

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'s', 's', 'd'}, []byte{'s', 's', 'd'}) {
		t.Error("Equal failed")
	}
}

// go test -65_fuzz FuzzEqual -fuzztime 10s
func FuzzEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
