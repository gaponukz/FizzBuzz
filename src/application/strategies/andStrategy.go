package strategies

type strategy interface {
	IsTruthy(int) bool
}

type andStrategy struct {
	strategies []strategy
}

func NewAndStrategy(strategies ...strategy) andStrategy {
	return andStrategy{strategies: strategies}
}

func (s andStrategy) IsTruthy(num int) bool {
	for _, s := range s.strategies {
		if !s.IsTruthy(num) {
			return false
		}
	}
	return true
}
