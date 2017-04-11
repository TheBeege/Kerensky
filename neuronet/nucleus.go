package neuronet

import (
	"math/rand"
	"fmt"

	"github.com/nu7hatch/gouuid"
)

const AVG_NUCLEUS_NEURON_COUNT int = 15
const NEURON_COUNT_VARIANCE int = 3

type Nucleus struct {
	id *uuid.UUID
	neurons           []*Neuron
	downstreamNucleus []*Nucleus
}

func NewNucleus() *Nucleus {
	fmt.Println("Generating new nucleus")
	varianceResult := rand.Intn(NEURON_COUNT_VARIANCE*2) - NEURON_COUNT_VARIANCE
	neurons := make([]*Neuron, 0)
	id, _ := uuid.NewV4()
	nucleus := &Nucleus{
		id: id,
		neurons: neurons,
	}
	for i := 0; i < AVG_NUCLEUS_NEURON_COUNT + varianceResult; i++ {
		nucleus.neurons = append(nucleus.neurons, newNeuron(nucleus))
	}
	return nucleus
}

func (c *Nucleus) process(inputs []float64) {
	outputs := make([]float64, 0)
	for _, neuron := range c.neurons {
		outputs = append(outputs, neuron.process(inputs))
	}
	for _, nucleus := range c.downstreamNucleus {
		nucleus.process(outputs)
	}
}

func (c *Nucleus) String() string {
	return fmt.Sprintf("Nucleus(id:%s, neurons:%v)", c.id, c.neurons)
}