package genalg

type Population []*Genome

func (p Population) Len() int {
	return len(p)
}

func (p Population) Less(i, j int) bool {
	return p[i].fitness < p[j].fitness
}

func (p Population) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
