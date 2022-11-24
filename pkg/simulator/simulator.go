package simulator

import (
	"io"

	"github.com/souvikhaldar/alien-invasion/pkg/writer"
)

type Simulator interface {
	Simulate()
	SaveState(writer.Writer, io.Writer) error
}
