package writer

import "io"

type Writer interface {
	Write(w io.Writer) error
}
