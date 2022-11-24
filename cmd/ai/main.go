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
	// the number of aliens to play in the simulation
	numAliens := flag.Int("N", 0, "Number of aliens")
	if numAliens == nil {
		log.Panicln("Number of aliens not provided")
	}

	// path to the configuration file
	// by default it is located in root dir
	configPath := flag.String("conf", "config.json", "path to the configuration file")
	flag.Parse()

	// load the config from given path
	conf := config.LoadConfig(*configPath)

	// set tge no of aliens if provided on command line
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
	prse := parser.NewMap(conf.NoOfCities, conf.PossibleRelations)

	// create the simulator
	sim := simulator.NewSimulation(
		prse,
		fileReader,
		conf.NoOfAliens,
		writer.NewState(fileReader, prse, fileWriter),
	)
	// run the simulation
	sim.Simulate()
}
