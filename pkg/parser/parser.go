package parser

import (
	"io"

	"github.com/souvikhaldar/go-ds/graph"
)

type Parser interface {
	// Parse parses the input map provided by reader object which can be
	// anything, from a file reader to network socket.
	// Then it creates the map and stores it as a graph
	Parse(io.Reader) (graph.Graph, error)
}
