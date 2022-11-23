package simulator

import (
	"fmt"
	"io"
	"log"

	"github.com/souvikhaldar/alien-invasion/pkg/parser"
	"github.com/souvikhaldar/alien-invasion/pkg/writer"
	"github.com/souvikhaldar/go-ds/graph"
)

type Alien struct {
	name        int
	currentCity string
}

type Simulation struct {
	cityMap     graph.Graph
	aliveAliens []Alien
	cities      []string
	writer      writer.Writer
}

func NewSimulation(
	p parser.Parser,
	i io.Reader,
	noOfAliens int,
	wt writer.Writer,

) *Simulation {
	s := new(Simulation)
	s.writer = wt
	var err error
	s.cityMap, err = p.Parse(i)
	if err != nil {
		log.Fatal("Could not parse the map")
	}

	s.cities = p.GetCities()
	s.aliveAliens = make([]Alien, noOfAliens)
	for i := 0; i < noOfAliens; i++ {
		s.aliveAliens[i].name = i
		s.aliveAliens[i].currentCity = s.GetRandomCity()
	}

	return s
}

func (s *Simulation) moveOneStep() {
	for _, a := range s.aliveAliens {
		a.currentCity = s.GetRandomNextCity(a.currentCity)
	}
}

func (s *Simulation) kill() {
	collitionMap := make(map[string][]int)
	for _, a := range s.aliveAliens {
		_, ok := collitionMap[a.currentCity]
		if ok {
			collitionMap[a.currentCity] = append(collitionMap[a.currentCity], a.name)
		} else {
			collitionMap[a.currentCity] = []int{a.name}
		}
	}
	for city, aliens := range collitionMap {
		if city == "" {
			// The alien is trapped
			continue
		}
		if len(aliens) > 1 {
			killmsg := fmt.Sprintf("%s has been destroyed by ", city)
			for pos, a := range aliens {
				if pos == len(aliens)-1 {
					killmsg += fmt.Sprintf("and alien %d.", a)
				} else if pos == len(aliens)-2 {
					killmsg += fmt.Sprintf("alien %d ", a)

				} else {
					killmsg += fmt.Sprintf("alien %d, ", a)
				}
				s.removeAlien(a)
			}
			log.Println(killmsg)
			// Remove the city from the map
			s.cityMap.DeleteNode(city)
			// Remove the aliens
		}
	}
}

func (s *Simulation) Simulate() {
	// let the aliens wander 10,000 times
	for i := 0; i < 10000; i++ {
		// move the aliens in any random direction one step
		s.moveOneStep()
		// check if there is collition, kill the alien and destroy the map
		s.kill()
	}
	// save the state of the map once the great wander is over
	s.saveState()
}

func (s *Simulation) saveState() {
	if err := s.writer.Write(); err != nil {
		log.Fatal("Unable to save the state of map: ", err)
	}
}
