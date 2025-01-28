package mirrorMove
//This is nearly syntactically Identical to Action. Just mirroring MoveUpdate
type ActionUpdate struct {
    Id              uint 
    Name            string
    CreatedAt       string 
    UpdatedAt       string
    IsHidden        bool 
    Description     string
    Seconds         float64
    Token           string
}