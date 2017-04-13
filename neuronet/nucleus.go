package neuronet

import (
	"math/rand"
	"fmt"

	"github.com/nu7hatch/gouuid"
	"github.com/TheBeege/Kerensky/config"
)

type Nucleus struct {
	id *uuid.UUID
	neurons           []*Neuron
	downstreamNucleus []*Nucleus
}

func NewNucleus(configData *config.Config) *Nucleus {
	fmt.Println("Generating new nucleus")
	varianceResult := rand.Intn(configData.NeuronCountVariance*2) - configData.NeuronCountVariance
	neurons := make([]*Neuron, 0)
	id, _ := uuid.NewV4()
	nucleus := &Nucleus{
		id: id,
		neurons: neurons,
	}
	for i := 0; i < configData.AvgNucleusNeuronCount + varianceResult; i++ {
		nucleus.neurons = append(nucleus.neurons, newNeuron(nucleus, configData))
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