package fuzz

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'s', 's', 'd'}, []byte{'s', 's', 'd'}) {
		t.Error("Equal failed")
	}
}

func FuzzEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
