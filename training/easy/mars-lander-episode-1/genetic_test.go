package main

import (
	"math"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandPower(t *testing.T) {
	rnd = rand.New(rand.NewSource(0))

	values := make([]float64, 5)
	power := 0.0
	for i := 0; i < 100; i++ {
		next := randPower(power)

		assert.True(t, math.Abs(next-power) <= 1)
		assert.True(t, next >= 0 && next <= powerMax)

		power = next
		values[int(power)]++
	}

	for i := 0; i < len(values); i++ {
		assert.True(t, values[i] > 0)
	}
}

func TestChromosome_Simulation(t *testing.T) {
	rnd = rand.New(rand.NewSource(0))

	game := testGame()
	chr := RandChromosome()
	result := chr.Simulation(game)

	assert.NotEqual(t, game.Lander, result.Lander)
	assert.True(t, result.IsLanderInBound())
	assert.True(t, result.IsLanderBetweenFlat())
	assert.True(t, result.IsLanderOnFlat())
	assert.True(t, result.IsLanderAngleOk())
	assert.False(t, result.IsLanderSpeedOk())
	assert.False(t, result.IsLanderLandedOk())

	assert.EqualValues(t, 360, chr.Fitness)
}

func TestRandChromosome(t *testing.T) {
	chr := RandChromosome()

	assert.Len(t, chr.Genes, chromosomeLength)
	var angle, power float64
	for _, gene := range chr.Genes {
		switch command := gene.(type) {
		case MoveCommand:
			assert.True(t, math.Abs(command.Angle) <= angleMax)
			assert.True(t, math.Abs(command.Angle-angle) <= angleStep)

			assert.True(t, command.Power >= 0 && command.Power <= powerMax)
			assert.True(t, math.Abs(command.Power-power) <= 1)

			angle = command.Angle
			power = command.Power
		}
	}
}

func TestRandPopulation(t *testing.T) {
	pop := RandPopulation()

	assert.Len(t, pop.Chromosomes, populationSize)
}

func TestEvolution_Init(t *testing.T) {
	e := NewEvolution(Game{})

	assert.Len(t, e.Population, populationSize)
}

func TestEvolution_Selection(t *testing.T) {
	game := testGame()
	e := NewEvolution(game)
	selected := e.Selection()

	assert.Len(t, selected, populationSize/2)
}

func TestEvolution_Crossover(t *testing.T) {
	rnd = rand.New(rand.NewSource(0))

	parent1 := RandChromosome()
	parent2 := RandChromosome()

	offspring := parent1.Copy()
	offspring.Crossover(parent2)

	var equal1, equal2 int
	for i := 1; i < chromosomeLength-1; i++ {
		if parent1.Genes[i] == offspring.Genes[i] {
			equal1++
			continue
		}

		if parent2.Genes[i] == offspring.Genes[i] {
			equal2++
		}

		prev := offspring.Genes[i-1].(MoveCommand)
		next := offspring.Genes[i+1].(MoveCommand)
		current := offspring.Genes[i].(MoveCommand)

		assert.True(t, math.Abs(current.Angle) <= angleMax)
		assert.True(t, math.Abs(current.Angle-prev.Angle) <= angleStep)
		assert.True(t, math.Abs(current.Angle-next.Angle) <= angleStep)
		assert.True(t, current.Power >= 0 && current.Power <= powerMax)
		assert.True(t, math.Abs(current.Power-prev.Power) <= 1)
		assert.True(t, math.Abs(current.Power-next.Power) <= 1)
	}

	assert.Len(t, offspring.Genes, chromosomeLength)
	assert.True(t, equal1 < chromosomeLength)
	assert.True(t, equal2 < chromosomeLength)
}

func TestEvolution_Mutation(t *testing.T) {
	rnd = rand.New(rand.NewSource(0))

	parent := RandChromosome()
	offspring := parent.Copy()
	offspring.Mutation()

	var equal int
	for i := 1; i < chromosomeLength-1; i++ {
		if parent.Genes[i] == offspring.Genes[i] {
			equal++
			continue
		}

		prev := offspring.Genes[i-1].(MoveCommand)
		next := offspring.Genes[i+1].(MoveCommand)
		current := offspring.Genes[i].(MoveCommand)

		assert.True(t, math.Abs(current.Angle) <= angleMax)
		assert.True(t, math.Abs(current.Angle-prev.Angle) <= angleStep)
		assert.True(t, math.Abs(current.Angle-next.Angle) <= angleStep)
		assert.True(t, current.Power >= 0 && current.Power <= powerMax)
		assert.True(t, math.Abs(current.Power-prev.Power) <= 1)
		assert.True(t, math.Abs(current.Power-next.Power) <= 1)
	}

	assert.Len(t, offspring.Genes, chromosomeLength)
	assert.True(t, equal < chromosomeLength)
}

func TestEvolution_Evolve(t *testing.T) {
	rnd = rand.New(rand.NewSource(0))

	game := testGame()
	e := NewEvolution(game)
	e.Evolve(1)

	assert.Len(t, e.Population.Chromosomes, populationSize)
}

func TestEvolution_EvolveProgress(t *testing.T) {
	rnd = rand.New(rand.NewSource(0))

	game := testGame()
	e := NewEvolution(game)
	e.Evolve(100)

	p := e.Population
	sort.Slice(p.Chromosomes, func(i, j int) bool {
		return p.Chromosomes[i].Fitness > p.Chromosomes[j].Fitness
	})

	best := p.Chromosomes[0]
	assert.InDelta(t, 1820, best.Fitness, 1)
}
