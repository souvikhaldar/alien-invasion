package parser

import (
	"os"
	"testing"

	"github.com/souvikhaldar/alien-invasion/pkg/utils/config"
)

func TestParse(t *testing.T) {
	config := config.LoadConfig("testfiles/config.json")
	fileReader, err := os.Open(config.InputFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer fileReader.Close()

	m := NewMap(config.NoOfCities, config.PossibleRelations)
	_, err = m.Parse(fileReader)
	if err != nil {
		t.Fatal(err)
	}
}
