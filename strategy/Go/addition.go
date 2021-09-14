package main

// Here is the concreate strategy for addition
type Addition struct{}

func (add Addition) calculate(x int, y int) int {
	return x + y
}
