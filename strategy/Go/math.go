package math

import (
	"fmt"

	"github.com/kwongjose/DesignPatterns/strategy/Go/addition"
	"github.com/kwongjose/DesignPatterns/strategy/Go/subtract"
)

type operation interface {
	Calculate(x int, y int) int
}

type calcStrategy struct {
	operation
}

var operationStrategy = calcStrategy{}

func Calculate(x int, y int) int {
	return operationStrategy.Calculate(x, y)
}

func SetStrategy(strategy string) {
	switch strategy {
	case "add":
		operationStrategy.Calculate = addition.Calculate
	case "sub":
		operationStrategy.Calculate = subtract.Calculate
	default:
		fmt.Println("Invalid strategy")
	}
}
