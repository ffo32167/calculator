package server

import (
	"errors"
	"fmt"

	"github.com/ffo32167/calculator/processer"
)

func Serve(p processer.AdditionalData) (float64, error) {
	a, b, err := getNumbers()
	if err != nil {
		return 0, fmt.Errorf("cant Serve: %w", err)
	}
	actionName := getActionName()
	action, err := processer.PickAction(actionName)
	if err != nil {
		return 0, fmt.Errorf("cant Serve: %w", err)
	}
	return action.Calculate(a, b)
}

func getNumbers() (a, b float64, err error) {
	if false {
		return 10, 3, errors.New("failed to getNumbers")
	}
	return 10, 3, nil
}

func getActionName() string {
	return "division"
}
