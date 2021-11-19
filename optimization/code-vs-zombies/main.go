package main

import (
	"fmt"
	"math"
)

const (
	ashSpeed    float64 = 1000
	ashShoot    float64 = 2000
	zombieSpeed float64 = 400
)

type (
	point struct {
		x, y float64
	}
	hero struct {
		crd point
	}
	human struct {
		id, uid int
		crd     point
	}
	zombie struct {
		id, uid       int
		crd, next     point
		humanDistance []float64
		heroDistance  float64
	}
	game struct {
		ash         hero
		humanCount  int
		humans      []human
		zombieCount int
		zombies     []zombie
	}
)

func (a *hero) safeTurnsToHuman(h human) float64 {
	d := distance(a.crd, h.crd)
	if d <= ashShoot {
		return 0
	}
	return math.Ceil((d - ashShoot) / ashSpeed)
}

func (z *zombie) turnsToHuman(h human) float64 {
	return math.Ceil(distance(z.crd, h.crd) / zombieSpeed)
}

func (z *zombie) closestHuman() (float64, int) {
	var cd = math.Inf(1)
	var cid = -1
	for i, hd := range z.humanDistance {
		if hd < cd {
			cd = hd
			cid = i
		}
	}
	return cd, cid
}

func (g *game) scan() {
	fmt.Scan(&g.ash.crd.x, &g.ash.crd.y)

	fmt.Scan(&g.humanCount)
	g.humans = make([]human, g.humanCount)
	for i := 0; i < g.humanCount; i++ {
		h := &g.humans[i]
		h.id = i
		fmt.Scan(&h.uid, &h.crd.x, &h.crd.y)
	}

	fmt.Scan(&g.zombieCount)
	g.zombies = make([]zombie, g.zombieCount)
	for i := 0; i < g.zombieCount; i++ {
		z := &g.zombies[i]
		z.id = i
		fmt.Scan(&z.uid, &z.crd.x, &z.crd.y, &z.next.x, &z.next.y)

		// distances
		z.humanDistance = make([]float64, g.humanCount)
		for i, h := range g.humans {
			z.humanDistance[i] = distance(h.crd, z.crd)
		}
		z.heroDistance = distance(g.ash.crd, z.crd)
	}
}

func (g *game) turn() {
	// run to last zombie
	if g.zombieCount == 1 {
		fmt.Println(int(g.zombies[0].next.x), int(g.zombies[0].next.y))
		return
	}

	// run to last human
	if g.humanCount == 1 {
		fmt.Println(int(g.humans[0].crd.x), int(g.humans[0].crd.y))
		return
	}

	// find nearest zombie
	zombieDistance := math.Inf(1)
	zid := -1
	for i, z := range g.zombies {
		closestDistance, cid := z.closestHuman()
		h := g.humans[cid]
		if closestDistance < zombieDistance {
			// skip long run
			if g.ash.safeTurnsToHuman(h) > z.turnsToHuman(h) {
				continue
			}
			zombieDistance = closestDistance
			zid = i
		}
	}
	if zid == -1 {
		fmt.Println(int(g.zombies[0].next.x), int(g.zombies[0].next.y))
		return
	}

	fmt.Println(int(g.zombies[zid].next.x), int(g.zombies[zid].next.y))
}

func distance(a point, b point) float64 {
	return math.Floor(math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2)))
}

/**
 * Save humans, destroy zombies!
 **/

func main() {
	game := game{}

	for {
		game.scan()
		game.turn()
	}
}
