package parser

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/souvikhaldar/go-ds/graph"
)

var (
	ErrCorruptLine = errors.New("Could not read line from file")
)

type Map struct {
	graph     graph.Graph
	relations map[string]int
	cities    []string
}

func NewMap(noOfCities int, relations []string) *Map {
	var counter int
	cityMap := new(Map)
	cityMap.relations = make(map[string]int)
	for _, rel := range relations {
		cityMap.relations[rel] = counter
		counter++
	}
	cityMap.graph = graph.NewAMGraph(
		noOfCities,
		noOfCities,
	)
	return cityMap
}

// Parse parses the data read and create a graph
func (m *Map) Parse(r io.Reader) (graph.Graph, error) {
	// scan the input line by line
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			log.Println(ErrCorruptLine)
			continue
		}
		// first element is the city
		city := fields[0]
		m.cities = append(m.cities, city)
		// all subsequent elements are relations of this city
		for i := 1; i < len(fields); i++ {
			relationTo := strings.Split(fields[i], "=")
			if len(relationTo) == 0 {
				return nil, ErrCorruptLine
			}
			relation := relationTo[0]
			to := relationTo[1]
			m.graph.AddNode(city, to, relation)
			if m.graph.GetRelationBetween(to, city) == "" {
				// add the opposite relation
				if opRel := getOppositeRelation(relation); opRel != "" {
					m.graph.AddNode(to, city, opRel)
				}
			}
		}
	}
	return m.graph, nil
}

func (m *Map) GetCities() []string {
	return m.cities
}
