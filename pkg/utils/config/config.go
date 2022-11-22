package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	InputFilePath     string   `json:"input_file_path"`
	NoOfAliens        int      `json:"no_of_aliens"`
	NoOfCities        int      `json:"no_of_cities"`
	PossibleRelations []string `json:"possible_relations"`
	OutputFilePath    string   `json:"output_file_path"`
}

func LoadConfig(filePath string) *Config {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}
	c := new(Config)
	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		log.Fatal("Error decoding the config: ", err)
	}
	return c
}
