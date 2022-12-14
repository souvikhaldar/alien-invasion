package writer

import (
	"os"
	"testing"

	"github.com/souvikhaldar/alien-invasion/pkg/parser"
	"github.com/souvikhaldar/alien-invasion/pkg/utils/config"
)

func TestWrite(t *testing.T) {
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
	g, err := m.Parse(fileReader)
	if err != nil {
		t.Fatal(err)
	}

	s := NewState(g, fileWriter)
	err = s.Write()
	if err != nil {
		t.Fatal(err)
	}
}
