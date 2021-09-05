package substraction

type substraction struct {
}

func NewSubstraction() substraction {
	return substraction{}
}

func (substraction) Calculate(first float64, last float64) (float64, error) {
	return first - last, nil
}
