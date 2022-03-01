package loadBalance

import (
	"errors"
	"math/rand"
)

func main() {

}

type RandomBalance struct {
	currentIndex int
	server       []string
	//conf         LoadBalanceConf
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	r.server = append(r.server, addr)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.server) == 0 {
		return ""
	}

	r.currentIndex = rand.Intn(len(r.server))
	return r.server[r.currentIndex]
}

func (r *RandomBalance) Get(key string) (string, error) {
	return r.Next(), nil
}

//func (r *RandomBalance) SetConf(conf LoadBalanceConf) *RandomBalance {
//
//	r1 := &RandomBalance{conf: conf}
//
//	return r1
//}

// 可变参数
