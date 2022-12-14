package simulator

import (
	"fmt"
	"io"
	"log"

	"github.com/souvikhaldar/alien-invasion/pkg/parser"
	"github.com/souvikhaldar/go-ds/graph"
)

type Alien struct {
	name        int
	currentCity string
}

type Simulation struct {
	cityMap        graph.Graph
	aliveAliens    []Alien
	killMsgPrinted map[string]struct{}
}

func NewSimulation(
	p parser.Parser,
	i io.Reader,
	noOfAliens int,
) *Simulation {
	s := new(Simulation)
	var err error
	s.cityMap, err = p.Parse(i)
	s.killMsgPrinted = make(map[string]struct{})
	if err != nil {
		log.Fatal("Could not parse the map")
	}

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
	// check which aliens are going to fight with each other
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
			continue
		}

		if _, ok := s.killMsgPrinted[city]; len(aliens) > 1 && !ok {
			// mark this city as printed
			s.killMsgPrinted[city] = struct{}{}
			// since more than 1 alien are present in one city
			// hence there is going to be a fight and the aliens
			// fighting will die along with the city
			killmsg := fmt.Sprintf("%s has been destroyed by ", city)
			for pos, a := range aliens {
				if pos == len(aliens)-1 {
					killmsg += fmt.Sprintf("and alien %d.", a)
				} else if pos == len(aliens)-2 {
					killmsg += fmt.Sprintf("alien %d ", a)

				} else {
					killmsg += fmt.Sprintf("alien %d, ", a)
				}
				// kill the alien
				s.removeAlien(a)
			}
			log.Println(killmsg)
			// Remove the city from the map
			s.cityMap.DeleteNode(city)
		}
	}
}

func (s *Simulation) Simulate() graph.Graph {
	// let the aliens wander 10,000 times
	for i := 0; i < 10000; i++ {
		// move the aliens in any random direction one step
		s.moveOneStep()
		// check if there is collition, kill the alien and destroy the map
		s.kill()
	}
	return s.cityMap
}
