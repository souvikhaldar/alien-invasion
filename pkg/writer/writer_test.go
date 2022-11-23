package writer

import (
	"os"
	"testing"

	"github.com/souvikhaldar/alien-invasion/pkg/parser"
	"github.com/souvikhaldar/alien-invasion/pkg/utils/config"
)

func TestWrite(t *testing.T) {
	fileReader, err := os.Open("input_map.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fileReader.Close()

	config := config.LoadConfig("../../config.json")
	m := parser.NewMap(config.NoOfCities, config.PossibleRelations)

	fileWriter, err := os.Create("output_map.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fileWriter.Close()

	s := NewState(fileReader, m)
	s.Write(fileWriter)
}
