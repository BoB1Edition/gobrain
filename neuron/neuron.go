package neuron

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Neuron struct {
	inputs           []float64
	Weights          []float64
	output           float64
	nu               float64
	nuChanged        int
	FunctionActivate functionActivate
}
type functionActivate func(inputs, weights []float64) float64

func randFloatBeetween(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

func NewNeuronWithActivate(count uint, fActivate functionActivate) Neuron {
	n := Neuron{}
	n.inputs = make([]float64, count)
	n.Weights = make([]float64, count)
	n.FunctionActivate = fActivate
	for i := range n.Weights {
		n.Weights[i] = randFloatBeetween(-2, 2)
	}
	n.nu = randFloatBeetween(-2, 2)
	n.nuChanged = rand.Intn(100)
	return n
}

func NewNeuron(count uint) Neuron {
	return NewNeuronWithActivate(count, defaultActivate)
}

func (n *Neuron) NuChanged() {
	n.nu = randFloatBeetween(-2, 2)
	n.nuChanged = rand.Intn(100)
}

func (n *Neuron) Forward(data []float64) (float64, error) {
	if len(n.inputs) != len(data) {
		return 0.0, fmt.Errorf("Количество входных данных: %d не соответствует количеству выходных %d\n", len(n.inputs), len(data))
	}
	for i := range n.inputs {
		n.inputs[i] = data[i]
	}
	n.output = n.FunctionActivate(n.inputs, n.Weights)
	/*if n.nuChanged <= 0 {
		n.NuChanged()
	}
	n.nuChanged -= 1*/
	return n.output, nil
}

func defaultActivate(inputs, weights []float64) float64 {
	var sum float64
	for i := range inputs {
		sum += inputs[i] * weights[i]
	}
	sum = 1 / (1 + math.Exp(-sum))
	/*if sum > 0.5 {
		sum = 1
	} else {
		sum = 0
	}*/
	return sum
}

func (n *Neuron) Adjustment(t float64) {
	fi := t - n.output
	for i := range n.Weights {
		dweight := n.inputs[i] * fi * n.nu
		n.Weights[i] = n.Weights[i] + dweight
	}
}
