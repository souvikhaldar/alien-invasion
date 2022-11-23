package writer

import (
	"bytes"
	"io"
	"log"

	"github.com/souvikhaldar/alien-invasion/pkg/parser"
	"github.com/souvikhaldar/go-ds/graph"
)

type State struct {
	cityMap graph.Graph
	cities  []string
}

func NewState(i io.Reader, p parser.Parser) *State {
	c, err := p.Parse(i)
	if err != nil {
		log.Println("Unable to create the map: ", err)
		return nil
	}
	return &State{
		cityMap: c,
		cities:  c.GetAllNodes(),
	}
}

func (s *State) Write(w io.Writer) {
	var buf bytes.Buffer
	for _, c := range s.cities {
		log.Println("Checking city: ", c)
		buf.WriteString(c)
		buf.WriteString(" ")
		neigh := s.cityMap.GetNeighboursOf(c)
		for pos, n := range neigh {
			log.Println("Neighbour: ", n)
			if pos == len(neigh)-1 {
				buf.WriteString(s.cityMap.GetRelationBetween(c, n) + "=" + n + "\n")
				break
			}
			buf.WriteString(s.cityMap.GetRelationBetween(c, n) + "=" + n + " ")
		}
		log.Println("Writing to file: ", buf.String())
		if _, err := w.Write(buf.Bytes()); err != nil {
			log.Println("Unable to write: ", err)
		}
		// reset before checking for next city
		buf.Reset()
	}

}
