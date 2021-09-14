package main

import (
	"fmt"

	"github.com/kwongjose/DesignPatterns/strategy/Go/math"
)

func main() {
	math.SetStrategy("add")
	r := math.Calculate(1, 2)
	fmt.Println(r)
}
