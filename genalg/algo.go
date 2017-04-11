package genalg

type Algo struct {
	population []*Genome
	chromosomeLength uint
	totalFitness float64
	bestFitness float64
	avgFitness float64
	worstFitness float64
	bestGenome *Genome
	mutationRate float64
	crossoverRate float64
	generationCount uint64
}

func NewAlgo(populationSize int, mutationRate float64, crossoverRate float64, numWeights int) *Algo {

	weights := make([]float64, 0)
	for i := 0; i < numWeights; i++ {
		weights = append(weights, 0)
	}

	population := make([]*Genome, 0)
	for i := 0; i < populationSize; i++ {
		population = append(population, NewGenome(weights, 0))
	}

	algo := &Algo{
		population:population,
		mutationRate:mutationRate,
		crossoverRate:crossoverRate,

	}
	return algo
}

func (a *Algo) Epoch() []*Genome {
	return nil
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

func (a *Algo) crossover() {

}

func (a *Algo) mutate() {

}

func (a *Algo) rouletteGetChromosome() *Genome {
	return nil
}

func (a *Algo) grabNBest(numBest uint, numCopies uint) []*Genome {
	return nil
}

func (a *Algo) calculateStats() {

}

func (a *Algo) reset() {

}

