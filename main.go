package main

import (
	"fmt"
)

func main() {
	n := neuron{}
	fmt.Printf("n: %+v\n", n)
	n1 := n.init(2)
	fmt.Printf("n: %+v\n", n)
	fmt.Printf("n1: %+v\n", n1)
}
