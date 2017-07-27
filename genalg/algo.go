package genalg

import (
	"fmt"
	"github.com/TheBeege/Kerensky/config"
	"github.com/TheBeege/Kerensky/utils"
	"math/rand"
	"reflect"
	"sort"
)

type Algo struct {
	population       Population
	chromosomeLength int
	totalFitness     float64
	bestFitness      float64
	avgFitness       float64
	worstFitness     float64
	bestGenome       *Genome
	mutationRate     float64
	crossoverRate    float64
	generationCount  uint64
	fittestGenome    *Genome
}

func NewAlgo(populationSize int, mutationRate float64, crossoverRate float64, numWeights int) *Algo {

	weights := make([]float64, 0)
	for i := 0; i < numWeights; i++ {
		weights = append(weights, 0)
	}

	population := Population(make([]*Genome, 0))
	for i := 0; i < populationSize; i++ {
		population = append(population, NewGenome(weights, 0))
	}

	algo := &Algo{
		population:    population,
		mutationRate:  mutationRate,
		crossoverRate: crossoverRate,
	}
	return algo
}

func (a *Algo) Epoch(originalPopulation []*Genome) []*Genome {
	population := Population(originalPopulation)
	a.reset()

	// Maybe should implement Len, Less, and Swap in Genome
	sort.Sort(population)

	a.calculateStats()

	newPopulation := make([]*Genome, 0)
	if config.Config.NumCopiesElite*config.Config.NumElite%2 == 0 {
		a.grabNBest(config.Config.NumElite, config.Config.NumCopiesElite)
	}
	for len(newPopulation) < len(a.population) {
		parent1 := a.rouletteGetChromosome()
		parent2 := a.rouletteGetChromosome()
		fmt.Printf("parent1: %s", parent1.String())

		var childWeights1, childWeights2 []float64
		a.crossover(parent1.weights, parent2.weights, childWeights1, childWeights2)

		a.mutate(childWeights1, childWeights2)

		newPopulation = append(newPopulation, NewGenome(childWeights1, 0))
		newPopulation = append(newPopulation, NewGenome(childWeights2, 0))
	}
	return newPopulation
}

func (a *Algo) GetChromosomes() []*Genome {
	return a.population
}

func (a *Algo) GetAverageFitness() float64 {
	return a.totalFitness / float64(len(a.population))
}

func (a *Algo) GetBestFitness() float64 {
	return a.bestFitness
}

func (a *Algo) crossover(parentWeights1 []float64, parentWeights2 []float64, offspringWeights1 []float64, offspringWeights2 []float64) {
	if utils.RandFloat() > a.crossoverRate || reflect.DeepEqual(parentWeights1, parentWeights2) {
		offspringWeights1 = parentWeights1
		offspringWeights2 = parentWeights2
		return
	}

	// In the tutorial, they do length -1 and ++i... I'm trying switching them
	crossoverPoint := rand.Intn(a.chromosomeLength)
	for i := 0; i < crossoverPoint; i++ {
		offspringWeights1 = append(offspringWeights1, parentWeights1[i])
		offspringWeights2 = append(offspringWeights2, parentWeights2[i])
	}
	for i := crossoverPoint; i < len(parentWeights1); i++ {
		offspringWeights1 = append(offspringWeights1, parentWeights2[i])
		offspringWeights2 = append(offspringWeights2, parentWeights1[i])
	}
}

func (a *Algo) mutate(childWeights1 []float64, childWeights2 []float64) {

}

func (a *Algo) rouletteGetChromosome() *Genome {
	//generate a random number between 0 & total fitness count
	randNum := utils.RandFloat() * a.totalFitness

	//this will be set to the chosen chromosome
	var theChosenOne *Genome

	//go through the chromosones adding up the fitness so far
	var fitnessSoFar float64 = 0

	for i := 0; i < len(a.population); i++ {
		fitnessSoFar += a.population[i].fitness

		//if the fitness so far > random number return the chromo at
		//this point
		if fitnessSoFar >= randNum {
			theChosenOne = a.population[i]
			break
		}
	}

	return theChosenOne
}

func (a *Algo) grabNBest(numBest int, numCopies int) {
	fmt.Printf("numBest: %d, len(a.population): %d, index: %d\n", numBest, len(a.population), (len(a.population)-1)-numBest)
	for ; numBest > 0; numBest-- {
		best := a.population[(len(a.population)-1)-numBest]
		for i := 0; i < numCopies; i++ {
			a.population = append(a.population, best)
		}
	}
}

func (a *Algo) calculateStats() {
	a.totalFitness = 0

	var highestSoFar float64 = 0
	var lowestSoFar float64 = 9999999

	for i := 0; i < len(a.population); i++ {
		//update fittest if necessary
		if a.population[i].fitness > highestSoFar {
			highestSoFar = a.population[i].fitness

			a.fittestGenome = a.population[i]

			a.bestFitness = highestSoFar
		}

		//update worst if necessary
		if a.population[i].fitness < lowestSoFar {
			lowestSoFar = a.population[i].fitness
			a.worstFitness = lowestSoFar
		}

		a.totalFitness += a.population[i].fitness

	} //next chromo

	a.avgFitness = a.totalFitness / float64(len(a.population))
}

func (a *Algo) reset() {
	a.totalFitness = 0
	a.bestFitness = 0
	a.worstFitness = 9999999
	a.avgFitness = 0
}
