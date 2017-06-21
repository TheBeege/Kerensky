package genalg

type Genome struct {
	weights []float64
	fitness float64
}

func NewGenome(weights []float64, fitness float64) *Genome {
	return &Genome{
		weights: weights,
		fitness: fitness,
	}
}

func (g *Genome) compare(other *Genome) int8 {
	if g.fitness > other.fitness {
		return 1
	} else if g.fitness < other.fitness {
		return -1
	} else {
		return 0
	}
}
