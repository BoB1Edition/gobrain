package network

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BoB1Edition/gobrain/neuron"
)

type Link map[int][]int

type iNeuron interface {
	Forward(data []float64) (float64, error)
}

type network struct {
	Inputlayer  []iNeuron
	Hiddenlayer [][]iNeuron
	Outputlayer []iNeuron
	Link        Link
}

func GenerateBackPropagation(inputs int, hidden []int, output int) network {
	nn := network{}
	nn.Inputlayer = make([]iNeuron, inputs)
	for i := range nn.Inputlayer {
		nn.Inputlayer[i] = neuron.NewNeuron(1)
	}
	nn.Hiddenlayer = make([][]iNeuron, len(hidden))
	for i := range nn.Hiddenlayer {
		nn.Hiddenlayer[i] = make([]iNeuron, hidden[i])
	}
	nn.Outputlayer = make([]iNeuron, output)
	for i := range nn.Outputlayer {
		nn.Outputlayer[i] = neuron.NewNeuron(uint(hidden[len(hidden)-1]))
	}
	nn.Link = make(Link)

	return nn
}

func LoadNetwork(filename string) (*network, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	Network := new(network)
	json.Unmarshal(data, Network)
	return Network, nil
}

func (n *network) SaveNetwork(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	data, err := json.Marshal(n)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return f.Close()
}

func (nn *network) Forward(inputs []float64) ([]float64, error) {
	answersInput := make([]float64, len(nn.Inputlayer))
	var err error
	for i := range nn.Inputlayer {
		answersInput[i], err = nn.Inputlayer[i].Forward(inputs)
		if err != nil {
			return nil, nil
		}
	}
	for _, layers := range nn.Hiddenlayer {
		for i := range layers {
			layers[i].Forward(answersInput)
		}
	}
}
