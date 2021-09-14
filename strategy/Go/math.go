package main

import (
	"fmt"
)

// operation is the interface our concrete strategies will use
type operation interface {
	calculate(x int, y int) int
}

// CalcStrategy is the proxy wraping around the concrete addition/subtraction strategy
type calcStrategy struct {
	operation operation
}

// Run executes the calcStrategy's operation.
func (proxy *calcStrategy) Run(x int, y int) int {
	return proxy.operation.calculate(x, y)
}

// SetStrategy sets the calcStrategy's operation to either addition or subtractions
func (proxy *calcStrategy) SetStrategy(strategy string) {
	switch strategy {
	case "add":
		proxy.operation = Addition{}
	case "sub":
		proxy.operation = Subtraction{}
	default:
		fmt.Println("Invalid strategy")
	}
}
