package mirrorMove

type MoveCreate struct {
	Name        string
	Description string
	Seconds     float64
	ActionLoops []MovesActionCreate
}
