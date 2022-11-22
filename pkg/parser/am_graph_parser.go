package parser

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/souvikhaldar/alien-invasion/pkg/utils/config"
	"github.com/souvikhaldar/go-ds/graph"
)

var (
	ErrCorruptLine = errors.New("Could not read line from file")
)

type Map struct {
	graph     graph.Graph
	relations map[string]int
}

func NewMap() *Map {
	var counter int
	config := config.LoadConfig("../../config.json")
	cityMap := new(Map)
	cityMap.relations = make(map[string]int)
	for _, rel := range config.PossibleRelations {
		cityMap.relations[rel] = counter
		counter++
	}
	cityMap.graph = graph.NewAMGraph(
		config.NoOfCities,
		config.NoOfCities,
	)
	return cityMap
}

// Parse parses the data read and create a graph
func (m *Map) Parse(r io.Reader) error {
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
		// all subsequent elements are relations of this city
		for i := 1; i < len(fields); i++ {
			relationTo := strings.Split(fields[i], "=")
			if len(relationTo) == 0 {
				return ErrCorruptLine
			}
			relation := relationTo[0]
			to := relationTo[1]
			m.graph.AddNode(city, to, relation)
		}
	}
	return nil
}
