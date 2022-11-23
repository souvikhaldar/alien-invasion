package parser

import (
	"io"

	"github.com/souvikhaldar/go-ds/graph"
)

type Parser interface {
	Parse(io.Reader) (graph.Graph, error)
	GetCities() []string
}
