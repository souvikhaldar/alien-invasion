package parser

import (
	"io"
)

type Parser interface {
	Parse(io.Reader) error
}
