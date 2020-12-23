// Package main repeatedly calls initfunc.GetInstance()
package main

import (
	"fmt"

	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/06-singleton/initfunc"
)

func main() {
	for i := 0; i < 40; i++ {
		go initfunc.GetInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
