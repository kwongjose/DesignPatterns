package main

// Here is the concreate strategy for subtraction
type Subtraction struct{}

func (sub Subtraction) calculate(x int, y int) int {
	return x - y
}
