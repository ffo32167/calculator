package server

import (
	"errors"
	"fmt"

	"github.com/ffo32167/calculator/processer"
)

func Serve(p processer.AdditionalData) (float64, error) {
	num1, num2, err := getNumbers()
	if err != nil {
		return 0, fmt.Errorf("cant Serve: %w", err)
	}

	actionName := getActionName()

	action, err := processer.NewAction(actionName, p)
	if err != nil {
		return 0, fmt.Errorf("cant Serve: %w", err)
	}

	return action.Calculate(num1, num2)
}

func getNumbers() (num1, num2 float64, err error) {
	if false {
		return 10, 3, errors.New("failed to getNumbers")
	}
	return 10, 3, nil
}

func getActionName() string {
	return "division"
}
