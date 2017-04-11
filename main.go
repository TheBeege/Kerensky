package main

import (
	"github.com/TheBeege/Kerensky/neuronet"
	"fmt"
)

const NUCLEUS_COUNT int = 1000

func main() {
	fmt.Println("========== GENERATING NUCLEI ==========")
	nuclei := make([]*neuronet.Nucleus, 0)
	for i := 0; i < NUCLEUS_COUNT; i++ {
		nuclei = append(nuclei, neuronet.NewNucleus())
	}

	for _, nucleus := range nuclei {
		fmt.Println(nucleus)
	}
}
