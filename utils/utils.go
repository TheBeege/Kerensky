package utils

import (
	"math"
	"math/rand"
)

func Sigmoid(input float64, activationResponse float64) float64 {
	return 1 / (1 + math.Exp(-input/activationResponse))
}

func RandFloat() float64 {
	//returns a random float in the range -1 < n < 1
	return rand.Float64() - rand.Float64()
}
