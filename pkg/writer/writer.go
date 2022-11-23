package writer

import (
	"io"

	"github.com/souvikhaldar/go-ds/graph"
)

type Writer interface {
	Write(graph.Graph, io.Writer) error
}
