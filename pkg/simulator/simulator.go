package simulator

import "github.com/souvikhaldar/go-ds/graph"

type Simulator interface {
	// Simulate simulates the invasion scenaio
	// and returns the final state of our planet
	// once the simulation is over
	Simulate() graph.Graph
}
