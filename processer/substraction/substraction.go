package substraction

type substraction struct {
}

func NewSubstraction() substraction {
	return substraction{}
}

func (substraction) Calculate(num1, num2 float64) (float64, error) {
	return num1 - num2, nil
}
