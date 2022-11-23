package simulator

import (
	"os"
	"testing"

	"github.com/souvikhaldar/alien-invasion/pkg/parser"
	"github.com/souvikhaldar/alien-invasion/pkg/utils/config"
	"github.com/souvikhaldar/alien-invasion/pkg/writer"
)

func TestSimulate(t *testing.T) {
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

	sw := writer.NewState(fileReader, m)

	s := NewSimulation(
		m,
		fileReader,
		config.NoOfAliens,
		fileWriter,
		sw,
	)
	s.Simulate()
}
