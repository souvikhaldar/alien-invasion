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
func (c *Config) SetInputFilePath(path string) {
	c.InputFilePath = path
}
func (c *Config) SetNoOfAliens(n int) {
	c.NoOfAliens = n
}
func (c *Config) SetNoOfCities(n int) {
	c.NoOfCities = n
}
func (c *Config) SetPossibleRelation(r []string) {
	c.PossibleRelations = r
}
func (c *Config) SetOutputFilePath(p string) {
	c.OutputFilePath = p
}
