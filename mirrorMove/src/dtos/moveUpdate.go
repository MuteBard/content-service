package mirrorMove

type MoveUpdate struct {
	Id          uint
	Name        string
	CreatedAt   string
	UpdatedAt   string
	IsHidden    bool
	Description string
	Seconds     float64
	ActionLoops []MovesActionCreate
}
