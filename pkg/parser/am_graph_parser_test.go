package parser

import (
	"os"
	"testing"

	"github.com/souvikhaldar/alien-invasion/pkg/utils/config"
)

func TestParse(t *testing.T) {
	fileReader, err := os.Open("../../static/input_map.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fileReader.Close()
	config := config.LoadConfig("../../config.json")
	m := NewMap(config.NoOfCities, config.PossibleRelations)
	_, err = m.Parse(fileReader)
	if err != nil {
		t.Fatal(err)
	}
}
