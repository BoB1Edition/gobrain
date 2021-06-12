package neuron

import (
	"math/rand"
	"time"
)

type neuron struct {
	input  []float64
	weight []float64
}

func (n neuron) init(count uint) neuron {
	n.input = make([]float64, count)
	n.weight = make([]float64, count)
	rand.Seed(time.Now().UnixNano())
	for i := range n.weight {
		n.weight[i] = rand.NormFloat64()
		rand.Seed(time.Now().UnixNano())
	}
	return n
}
