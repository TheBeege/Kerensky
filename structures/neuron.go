package structures

import (
	"math/rand"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"math"
)

const INPUT_COUNT int = 10
const ACTIVATION_RESPONSE float64 = 1 // not sure what this is yet
const BIAS float64 = -1

type Neuron struct {
	id *uuid.UUID
	weights             []float64
	nucleus             *Nucleus
}

func newNeuron(nucleus *Nucleus) *Neuron {
	fmt.Println("Generating new neuron")
	id, _ := uuid.NewV4()
	neuron := &Neuron{
		id: id,
		weights: make([]float64, 0),
		nucleus: nucleus,
	}
	for i := 0; i < INPUT_COUNT+1; i++  { // we're using the +1 as our activation threshold
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
	result += n.weights[len(n.weights)-1] * BIAS
	return sigmoid(result)
}

func sigmoid(input float64) float64 {
	return 1 / (1 + math.Exp(-input / ACTIVATION_RESPONSE))
}

func randFloat() float64 {
	//returns a random float in the range -1 < n < 1
	return rand.Float64() - rand.Float64()
}

func (n *Neuron) String() string {
	return fmt.Sprintf("Neuron(id:%s, weights:%v)", n.id, n.weights)
}
