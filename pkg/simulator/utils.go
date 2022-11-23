package simulator

import (
	"math/rand"
	"time"
)

func (s *Simulation) removeAlien(name int) {
	aliens := make([]int, 0)
	for _, a := range s.aliveAliens {
		if a.name != name {
			aliens = append(aliens, a)
		}
	}
	s.aliveAliens = aliens
}

func (s *Simulation) GetRandomCity() string {
	rand.Seed(time.Now().UnixNano())
	nonce := rand.Intn(100) % len(s.cities)

	return s.cities[nonce]
}

func (s *Simulation) GetRandomNextCity(currentCity string) string {
	rand.Seed(time.Now().UnixNano())
	neighbours := s.graph.GetNeighboursOf(currentCity)
	nonce := rand.Intn(100) % len(neighbours)
	return neighbours[nonce]
}
