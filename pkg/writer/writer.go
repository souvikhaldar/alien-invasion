package writer

import "io"

type Writer interface {
	Write(io.Writer) error
}
