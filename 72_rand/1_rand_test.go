package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestTimeDurtion(t *testing.T) {
	time, err := RandomSeconds(10)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(time)
	t.Log(time)
}

func TestCryptoRandSecure(t *testing.T) {
	cryptoRandSecure, err := cryptoRandSecure(0, 10)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(cryptoRandSecure)
}

func TestRandomStr(t *testing.T) {
	str := RandomStr(10)
	fmt.Println(str)
}

func TestRandomIntStr(t *testing.T) {
	str := RandomIntStr(10)
	fmt.Println(str)
}

// go test   -fuzztime 10s
func FuzzCryptoRandSecure(f *testing.F) {
	f.Fuzz(func(t *testing.T, a1 int64, b1 int64) {
		var a int64 = 0
		var b int64 = 10
		cryptoRandSecure, err := cryptoRandSecure(a, b)
		if err != nil {
			t.Log(err)
		}
		if cryptoRandSecure < a || cryptoRandSecure >= b {
			log.Fatal("cryptoRandSecure < 0 || cryptoRandSecure >= 10")
			t.Error("cryptoRandSecure < 0")
		}
	})
}

func TestRandomStrWith13Ip(t *testing.T) {
	str := RandomStrWith13Ip(10)
	fmt.Println(str)
}
