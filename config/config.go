package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

// https://blog.gopheracademy.com/advent-2014/reading-config-files-the-go-way/
var Config struct {
	NucleusCount          int
	InputCount            int
	ActivationResponse    float64
	Bias                  float64
	AvgNucleusNeuronCount int
	NeuronCountVariance   int
	CrossoverRate         float64
	MutationRate          float64
	NumCopiesElite        int
	NumElite              int
}

// Reads info from config file
func ReadConfig() {
	var filePath = *flag.String("configFile", "config.toml", "The path to the data file to be used")
	log.Println("Reading data file...")
	cwd, _ := os.Getwd()
	log.Println("cwd:", cwd)
	filePath = cwd + "/" + filePath
	log.Println("filePath:", filePath)
	_, err := os.Stat(filePath)
	if err != nil {
		log.Fatal("Config file is missing: ", filePath)
	}

	log.Println("Decoding data file...")
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		log.Fatal(err)
	}
}
