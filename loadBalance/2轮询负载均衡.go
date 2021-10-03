package loadBalance

import "errors"

type RoundRobinBalance struct {
	currentIndex int
	server       []string
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len at 1 least")
	}
	addr := params[0]
	r.server = append(r.server, addr)
	return nil
}

func (r *RoundRobinBalance) Next() string {
	if len(r.server) == 0 {
		return ""
	}

	lens := len(r.server)
	if r.currentIndex >= lens {
		r.currentIndex = 0
	}

	currentAddr := r.server[r.currentIndex]
	r.currentIndex = (r.currentIndex + 1) % lens // 求余
	return currentAddr
}
