package writer

import (
	"bytes"
	"io"
	"log"

	"github.com/souvikhaldar/go-ds/graph"
)

type State struct {
	cityMap graph.Graph
	writer  io.Writer
}

func NewState(
	c graph.Graph,
	w io.Writer,
) *State {
	return &State{
		cityMap: c,
		writer:  w,
	}
}

func (s *State) Write() error {
	var buf bytes.Buffer
	for _, c := range s.cityMap.GetAllNodes() {
		buf.WriteString(c)
		neigh := s.cityMap.GetNeighboursOf(c)
		if len(neigh) == 0 {
			buf.WriteString("\n")
		} else {
			buf.WriteString(" ")

		}
		for pos, n := range neigh {
			if pos == len(neigh)-1 {
				buf.WriteString(s.cityMap.GetRelationBetween(c, n) + "=" + n + "\n")
				break
			}
			buf.WriteString(s.cityMap.GetRelationBetween(c, n) + "=" + n + " ")
		}
		if _, err := s.writer.Write(buf.Bytes()); err != nil {
			log.Println("Unable to write: ", err)
			return err

		}
		// reset before checking for next city
		buf.Reset()
	}
	return nil
}
