package neuronet

import (
	"fmt"
	"math/rand"

	"github.com/TheBeege/Kerensky/config"
	"github.com/nu7hatch/gouuid"
)

type Nucleus struct {
	id                *uuid.UUID
	neurons           []*Neuron
	downstreamNucleus []*Nucleus
}

func NewNucleus() *Nucleus {
	fmt.Println("Generating new nucleus")
	varianceResult := rand.Intn(config.Config.NeuronCountVariance*2) - config.Config.NeuronCountVariance
	neurons := make([]*Neuron, 0)
	id, _ := uuid.NewV4()
	nucleus := &Nucleus{
		id:      id,
		neurons: neurons,
	}
	for i := 0; i < config.Config.AvgNucleusNeuronCount+varianceResult; i++ {
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
