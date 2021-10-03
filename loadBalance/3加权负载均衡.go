package loadBalance

import (
	"errors"
	"strconv"
)

type WeightRoundRobinBalance struct {
	currentIndex int
	server       []*WeightNode // 管理所有节点数组
	rsw          []int

	// 观察主题
	//conf LoadBalanceConf
}

type WeightNode struct {
	addr            string
	weight          int // 权重值
	currentWeight   int // 节点当前权重
	effectiveWeight int // 有效权重
}

// 增加方法 1.构造节点 2.添加节点
func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) == 2 {
		return errors.New("param len need 2")
	}
	parseInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{addr: params[0], weight: int(parseInt)}
	node.effectiveWeight = node.weight

	r.server = append(r.server, node)
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	total := 0
	var best *WeightNode
	for i := 0; i < len(r.server); i++ {
		w := r.server[i]

		// step1 计算所有有效权重之和
		total += w.effectiveWeight
		// step2 变更节点临时权重为的节点临时权重+节点有效权重
		w.currentWeight += w.effectiveWeight
		// step3 有效权重默认与权重相同，通讯异常时-1，通讯成功+1，直到回复到weight大小
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		// step4 选择最大临时权重节点
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}

	// step5 变更临时权重 临时权重 - 有效权重之和
	best.currentWeight -= total
	return best.addr

}

func (r *WeightRoundRobinBalance) Get(key string) (string, error) {
	return r.Next(), nil
}
