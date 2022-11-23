package writer

import (
	"io"

	"github.com/souvikhaldar/go-ds/graph"
)

type Write struct{}

func (w *Write) Write(_ graph.Graph, _ io.Writer) error {
	panic("not implemented") // TODO: Implement
}
