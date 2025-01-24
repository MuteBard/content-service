package mirrorMove

type MoveUpdate struct {
	Id              uint
    Name            string
    CreatedAt       string
    IsHidden        bool
    Description     string
    Seconds         float64
	ActionLoops   	[]MoveActionCreate
}

