package main

import "github.com/ffo32167/calculator/processer"

type config struct {
	notUsedField string
	processer.AdditionalData
}

func newConfig() config {
	return config{"val", processer.AdditionalData{Precision: 2}}
}
