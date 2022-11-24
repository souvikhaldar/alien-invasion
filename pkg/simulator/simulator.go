package simulator

import (
	"io"

	"github.com/souvikhaldar/alien-invasion/pkg/writer"
)

type Simulator interface {
	// Simulate simulates the invasion scenaio
	Simulate()
	// SaveState saves the final state to the destination provided
	SaveState(writer.Writer, io.Writer) error
}
