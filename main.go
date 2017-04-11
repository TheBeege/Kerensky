package main

import (
	"github.com/TheBeege/Kerensky/structures"
	"fmt"
)

const NUCLEUS_COUNT int = 1000

func main() {
	fmt.Println("========== GENERATING NUCLEI ==========")
	nuclei := make([]*structures.Nucleus, 0)
	for i := 0; i < NUCLEUS_COUNT; i++ {
		nuclei = append(nuclei, structures.NewNucleus())
	}

	for _, nucleus := range nuclei {
		fmt.Println(nucleus)
	}
}
