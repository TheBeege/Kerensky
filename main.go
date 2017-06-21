package main

import (
	"github.com/TheBeege/Kerensky/config"
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

	for _, nucleus := range nuclei {
		log.Println(nucleus)
	}
}
