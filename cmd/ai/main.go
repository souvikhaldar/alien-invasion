package main

import (
	"flag"
	"log"
	"os"

	"github.com/souvikhaldar/alien-invasion/pkg/parser"
	"github.com/souvikhaldar/alien-invasion/pkg/simulator"
	"github.com/souvikhaldar/alien-invasion/pkg/utils/config"
	"github.com/souvikhaldar/alien-invasion/pkg/writer"
)

func main() {
	numAliens := flag.Int("N", 0, "Number of aliens")
	flag.Parse()
	if numAliens == nil {
		log.Panicln("Number of aliens not provided")
	}
	log.Printf("Creating %d aliens...\n", *numAliens)
	conf := config.LoadConfig("../../config.json")
	conf.SetNoOfAliens(*numAliens)

	// create the output file
	fileWriter, err := os.Create(conf.OutputFilePath)
	if err != nil {
		log.Fatal("Could not create the output file: ", err)
	}
	defer fileWriter.Close()

	// Read the input map
	fileReader, err := os.Open(conf.InputFilePath)
	if err != nil {
		log.Fatal("Could not read the input file: ", err)
	}
	defer fileReader.Close()

	// create the parser that would parse the map
	par := parser.NewMap(conf.NoOfCities, conf.PossibleRelations)

	// create the simulator
	sim := simulator.NewSimulation(
		par,
		fileReader,
		conf.NoOfAliens,
		fileWriter,
		writer.NewState(fileReader, par),
	)
	// run the simulation
	sim.Simulate()
}
