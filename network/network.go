package network

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BoB1Edition/gobrain/neuron"
)

type Link map[int][]int

type network struct {
	Inputlayer  []neuron.Neuron
	Hiddenlayer [][]neuron.Neuron
	Outputlayer []neuron.Neuron
	Link        []Link
}

func (n *network) fillLayer(layer []neuron.Neuron, count uint) {
	for i := range layer {
		layer[i] = neuron.NewNeuron(count)
	}
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
