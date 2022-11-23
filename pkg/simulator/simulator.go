package simulator

type Simulator interface {
	MoveOneStep()
	Simulate()
	Kill()
	SaveState()
}
