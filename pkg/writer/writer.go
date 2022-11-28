package writer

type Writer interface {
	// Write parses the state of the map and writes it to the destination provided
	Write() error
}
