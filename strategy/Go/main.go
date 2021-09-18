package main

import (
	"fmt"
)

func main() {
	// Create the strategy wrapper. Use a pointer so that  the proxy is assigned to the reference wrapper
	var proxy = &calcStrategy{}

	proxy.SetStrategy("add")
	result := proxy.Run(1, 2)
	fmt.Printf("1 + 2 = %d\n", result)

	proxy.SetStrategy("sub")
	result = proxy.Run(result, 2)
	fmt.Printf("3 - 2 = %d\n", result)
}
