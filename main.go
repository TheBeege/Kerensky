package main

import (
	"github.com/TheBeege/Kerensky/neuronet"
	"github.com/TheBeege/Kerensky/config"
	"log"
)

func main() {
	log.Println("========== READING CONFIG ==========")
	configData := config.ReadConfig()
	log.Println("========== GENERATING NUCLEI ==========")
	nuclei := make([]*neuronet.Nucleus, 0)
	for i := 0; i < configData.NucleusCount; i++ {
		nuclei = append(nuclei, neuronet.NewNucleus(configData))
	}

	for _, nucleus := range nuclei {
		log.Println(nucleus)
	}
}
