package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"log"
	"flag"
)

type Config struct {
	NucleusCount int
	InputCount int
	ActivationResponse float64
	Bias float64
	AvgNucleusNeuronCount int
	NeuronCountVariance int
}


// Reads info from config file
func ReadConfig() *Config {
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
	data := &Config{}
	if _, err := toml.DecodeFile(filePath, data); err != nil {
		log.Fatal(err)
	}
	return data
}