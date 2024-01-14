package main

import (
	"math"
	"math/rand"
	"sort"
)

const (
	chromosomeLength = 100
	populationSize   = 100

	mutationRate = 0.2
)

func randPower(current float64) float64 {
	r := rnd.Float64()
	var p float64
	if r < 0.15 {
		p = -1
	} else if r < 0.3 {
		p = 1
	} else {
		p = 0
	}

	return math.Min(math.Max(current+p, 0), powerMax)
}

type Chromosome struct {
	Genes Commands

	Fitness float64
}

func (c *Chromosome) Rand() {
	c.Genes = make(Commands, 0, chromosomeLength)

	var angle, power float64
	for i := 0; i < chromosomeLength; i++ {
		power = randPower(power)
		c.Genes = append(c.Genes, MoveCommand{
			angle,
			power,
		})
	}
}

func (c *Chromosome) Simulation(game Game) Game {
	var score float64

	r := Simulate(game, c.Genes)

	// landing
	if r.IsLanderLandedOk() {
		score += 1000
	}

	// landing speed
	vSpeed := math.Abs(r.Lander.Speed.Y)
	diff := ExpectedExpDiff(vSpeedMax, vSpeed, 2)
	score += diff * 500

	// fuel
	score += r.Lander.Fuel

	c.Fitness = score

	return r
}

func (c *Chromosome) Copy() *Chromosome {
	r := &Chromosome{
		Genes: make(Commands, len(c.Genes)),
	}
	copy(r.Genes, c.Genes)
	r.Fitness = 0
	return r
}

func canCrossover(gene1, gene2 Command) bool {
	switch g1 := gene1.(type) {
	case MoveCommand:
		switch g2 := gene2.(type) {
		case MoveCommand:
			if math.Abs(g1.Angle-g2.Angle) > angleStep {
				return false
			}
			if math.Abs(g1.Power-g2.Power) > powerStep {
				return false
			}
		}
	}
	return true
}

func (c *Chromosome) Crossover(mother *Chromosome) {
	for i := 1; i < len(c.Genes)-1; i++ {
		if rnd.Float64() > 0.4 {
			continue
		}

		prev := c.Genes[i-1].(MoveCommand)
		next := c.Genes[i+1].(MoveCommand)

		current := mother.Genes[i].(MoveCommand)
		current.Power = randPower(prev.Power)

		if !canCrossover(next, current) {
			continue
		}

		c.Genes[i] = current
	}
}

func (c *Chromosome) Mutation() {
	for i := 1; i < len(c.Genes)-1; i++ {
		if rnd.Float64() > mutationRate {
			continue
		}

		prev := c.Genes[i-1].(MoveCommand)
		next := c.Genes[i+1].(MoveCommand)

		current := c.Genes[i].(MoveCommand)
		current.Power = randPower(prev.Power)

		if !canCrossover(next, current) {
			continue
		}

		c.Genes[i] = current
	}
}

func RandChromosome() *Chromosome {
	r := &Chromosome{}
	r.Rand()
	return r
}

type Population struct {
	Chromosomes []*Chromosome
}

func (p *Population) Rand() {
	p.Chromosomes = make([]*Chromosome, 0, populationSize)
	for i := 0; i < populationSize; i++ {
		p.Chromosomes = append(p.Chromosomes, RandChromosome())
	}
}

func RandPopulation() *Population {
	r := &Population{}
	r.Rand()
	return r
}

type Evolution struct {
	Game Game

	Population *Population
}

func (e *Evolution) Selection() []*Chromosome {
	for _, dna := range e.Population.Chromosomes {
		dna.Simulation(e.Game)
	}

	selected := make([]*Chromosome, 0, populationSize)
	for _, chromosome := range e.Population.Chromosomes {
		if chromosome.Fitness >= 0 {
			selected = append(selected, chromosome)
		}
	}
	sort.Slice(selected, func(i, j int) bool {
		return selected[i].Fitness > selected[j].Fitness
	})

	return selected[:populationSize/2]
}

func (e *Evolution) Evolve(generations int) {
	for i := 0; i < generations; i++ {
		selection := e.Selection()

		// elitism
		elite := selection[:populationSize/10]

		population := &Population{}
		population.Chromosomes = make([]*Chromosome, 0, populationSize)
		// crossover
		for _, parent := range selection {
			offspring := parent.Copy()
			mother := selection[rand.Intn(len(selection))]

			offspring.Crossover(mother)
			offspring.Mutation()

			population.Chromosomes = append(population.Chromosomes, offspring)
		}

		population.Chromosomes = append(population.Chromosomes, elite...)
		// fill up
		if len(population.Chromosomes) < populationSize {
			for ii := len(population.Chromosomes); ii < populationSize; ii++ {
				population.Chromosomes = append(population.Chromosomes, RandChromosome())
			}
		}

		e.Population = population
	}
}

func NewEvolution(game Game) *Evolution {
	r := &Evolution{
		Game:       game,
		Population: RandPopulation(),
	}
	return r
}
