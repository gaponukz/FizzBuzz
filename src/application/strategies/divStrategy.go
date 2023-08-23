package strategies

type divStrategy struct {
	devider int
}

func NewDivStrategy(devider int) divStrategy {
	return divStrategy{devider: devider}
}

func (s divStrategy) IsTruthy(num int) bool {
	return num%s.devider == 0
}
