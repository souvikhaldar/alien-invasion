package parser

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	fileReader, err := os.Open("../../static/input_map.txt")
	if err != nil {
		t.Fatal(err)
	}
	m := NewMap()
	err = m.Parse(fileReader)
	if err != nil {
		t.Fatal(err)
	}
}
