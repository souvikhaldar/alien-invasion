package simulator

type Simulator interface {
	moveOneStep()
	Simulate()
	kill()
	saveState()
}
