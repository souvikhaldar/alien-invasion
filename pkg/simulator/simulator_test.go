package simulator

import (
	"os"
	"testing"

	"github.com/souvikhaldar/alien-invasion/pkg/parser"
	"github.com/souvikhaldar/alien-invasion/pkg/utils/config"
	"github.com/souvikhaldar/alien-invasion/pkg/writer"
)

func TestSimulate(t *testing.T) {
	config := config.LoadConfig("testfiles/config.json")
	fileReader, err := os.Open(config.InputFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer fileReader.Close()

	m := parser.NewMap(config.NoOfCities, config.PossibleRelations)

	fileWriter, err := os.Create(config.OutputFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer fileWriter.Close()

	sw := writer.NewState(fileReader, m, fileWriter)

	s := NewSimulation(
		m,
		fileReader,
		config.NoOfAliens,
		sw,
	)
	s.Simulate()
}
