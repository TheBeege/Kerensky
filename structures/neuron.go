package structures

import (
	"math/rand"
	"fmt"
	"github.com/nu7hatch/gouuid"
)

const INPUT_COUNT int = 10
const AVG_NEURON_WEIGHTS float64 = 1.0
const WEIGHT_VARIANCE float64 = 0.2
const AVG_ACTIVATION_THRESHOLD float64 = 10
const ACTIVATION_VARIANCE float64 = 1.5

type Neuron struct {
	id *uuid.UUID
	weights             []float64
	activationThreshold float64
	nucleus             *Nucleus
}

func newNeuron(nucleus *Nucleus) *Neuron {
	fmt.Println("Generating new neuron")
	activationThreshold := AVG_ACTIVATION_THRESHOLD + (rand.Float64() * ACTIVATION_VARIANCE)
	if rand.Float64() > 0.5 {
		activationThreshold *= -1
	}
	id, _ := uuid.NewV4()
	neuron := &Neuron{
		id: id,
		weights: make([]float64, 0),
		activationThreshold: activationThreshold,
		nucleus: nucleus,
	}
	for i := 0; i < INPUT_COUNT; i++  {
		varianceResult := rand.Float64() * WEIGHT_VARIANCE
		if rand.Float64() > 0.5 {
			varianceResult *= -1
		}
		neuron.weights = append(neuron.weights, AVG_NEURON_WEIGHTS + varianceResult)
	}
	return neuron
}

func (n *Neuron) process(inputs []float64) float64 {
	var activationValue float64 = 0
	for index, input := range inputs {
		activationValue += input * (n.weights[index % len(n.weights)])
	}
	if activationValue >= n.activationThreshold {
		return 1
	} else {
		return 0
	}
}

func (n *Neuron) String() string {
	return fmt.Sprintf("Neuron(id:%s, weights:%v, activationThreshold:%f", n.id, n.weights, n.activationThreshold)
}
