package main

import (
	"fmt"

	"github.com/ffo32167/calculator/server"
)

func main() {
	config := newConfig()
	v, err := server.Serve(config.AdditionalData)
	if err != nil {
		fmt.Printf("can't process Serve: %s", err)
	} else {
		fmt.Printf("converted value: %f and it type is %T", v, v)
	}
}
