package loadBalance

type LoadBalance interface {
	Add(...string) error
	Get(string) (string, error)
}

type LbType int

const (
	LbRandom LbType = iota
	LbRoundRobin
	LbWeightRoundRobin
	LbConsistentHash
)

// 简单工厂方法 根据参数的不同 返回不同的实例
// 对来的类型进行判断
func LoadBaLanceFactory(lbType LbType) LoadBalance {
	switch lbType {
	case LbRandom:
		return &RandomBalance{}
	case LbConsistentHash:
		return
	case LbRoundRobin:
		return &RoundRobinBalance{}
	case LbWeightRoundRobin:
		return &WeightRoundRobinBalance{}
	default:
		return &RandomBalance{}
	}
}
