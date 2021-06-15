package neuron

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type neuron struct {
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

func NewNeuronWithActivate(count uint, fActivate functionActivate) *neuron {
	n := new(neuron)
	n.inputs = make([]float64, count+1)
	n.Weights = make([]float64, count+1)
	n.FunctionActivate = fActivate
	for i := range n.Weights {
		n.Weights[i] = randFloatBeetween(-2, 2)
	}
	n.nu = randFloatBeetween(-2, 2)
	n.nuChanged = rand.Intn(100)
	return n
}

func NewNeuron(count uint) *neuron {
	return NewNeuronWithActivate(count, defaultActivate)
}

func (n *neuron) NuChanged() {
	n.nu = randFloatBeetween(-2, 2)
	n.nuChanged = rand.Intn(100)
}

func (n *neuron) Forward(data []float64) (float64, error) {
	if len(n.inputs) != len(data)+1 {
		return 0.0, fmt.Errorf("Количество входных данных: %d не соответствует количеству выходных %d\n", len(n.inputs), len(data))
	}
	for i := range n.inputs {
		if i == 0 {
			continue
		}
		n.inputs[i] = data[i-1]
	}
	n.inputs[0] = 1
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

func (n *neuron) Adjustment(t float64) {
	fi := t - n.output
	for i := range n.Weights {
		dweight := n.inputs[i] * fi * n.nu
		n.Weights[i] = n.Weights[i] + dweight
	}
}
