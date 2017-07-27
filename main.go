package main

import (
	"github.com/TheBeege/Kerensky/config"
	"github.com/TheBeege/Kerensky/genalg"
	"github.com/TheBeege/Kerensky/neuronet"
	"log"
)

func main() {
	log.Println("========== READING CONFIG ==========")
	config.ReadConfig()
	log.Println("========== GENERATING NUCLEI ==========")
	nuclei := make([]*neuronet.Nucleus, 0)
	for i := 0; i < config.Config.NucleusCount; i++ {
		nuclei = append(nuclei, neuronet.NewNucleus())
	}

	algo := genalg.NewAlgo(
		config.Config.NucleusCount,
		config.Config.MutationRate,
		config.Config.CrossoverRate,
		10) // numWeights

	population := make([]*genalg.Genome, 0)

	algo.Epoch(population)
}
