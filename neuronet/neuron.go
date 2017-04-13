package neuronet

import (
	"math/rand"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"math"
	"github.com/TheBeege/Kerensky/config"
)

type Neuron struct {
	id *uuid.UUID
	weights             []float64
	nucleus             *Nucleus
	bias float64
	activationResponse float64
}

func newNeuron(nucleus *Nucleus, configData *config.Config) *Neuron {
	fmt.Println("Generating new neuron")
	id, _ := uuid.NewV4()
	neuron := &Neuron{
		id:                 id,
		weights:            make([]float64, 0),
		nucleus:            nucleus,
		bias:               configData.Bias,
		activationResponse: configData.ActivationResponse,
	}
	for i := 0; i < configData.InputCount+1; i++  { // we're using the +1 as our activation threshold
		neuron.weights = append(neuron.weights, randFloat())
	}
	return neuron
}

func (n *Neuron) process(inputs []float64) float64 {
	var result float64 = 0
	// process inputs
	for index, input := range inputs {
		// we skip the last weight, since that's being used for our bias
		result += input * (n.weights[index % (len(n.weights)-2)])
	}
	// add in the bias
	result += n.weights[len(n.weights)-1] * n.bias
	return sigmoid(result, n.activationResponse)
}

func (n *Neuron) GetWeights() []float64 {
	return n.weights
}

func (n *Neuron) SetWeights(weights []float64) {
	n.weights = weights
}

func (n *Neuron) String() string {
	return fmt.Sprintf("Neuron(id:%s, weights:%v)", n.id, n.weights)
}


func sigmoid(input float64, activationResponse float64) float64 {
	return 1 / (1 + math.Exp(-input / activationResponse))
}

func randFloat() float64 {
	//returns a random float in the range -1 < n < 1
	return rand.Float64() - rand.Float64()
}