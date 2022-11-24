package writer

import "io"

type Writer interface {
	// Write parses the state of the map and writes it to the destination provided
	Write(io.Writer) error
}
